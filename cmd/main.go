package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	// Import the generated protobuf code
	voicestream "github.com/ethanrcohen/voiceStreamClient/api"
	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// loadAndProcessAudioFile loads a WAV file and returns a slice of byte slices, each representing an audio channel.
func loadAndProcessAudioFile(filePath string) ([][]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("failed to get file info: %v", err)
	}
	// Subtracting the typical WAV header size to approximate the size of audio data.
	// This is a simplification and may vary based on the WAV file specifics.
	totalBytes := fileInfo.Size() - 44

	decoder := wav.NewDecoder(file)
	decoder.ReadInfo()

	if decoder.WavAudioFormat != 1 { // 1 means Linear PCM
		return nil, fmt.Errorf("unsupported audio format: must be Linear PCM")
	}
	if decoder.SampleRate != 8000 {
		return nil, fmt.Errorf("unsupported sample rate: must be 8000 Hz")
	}
	// Read the full format but not the data
	format := decoder.Format()

	// Assuming 16 bits per sample, calculate total samples
	// For stereo (2 channels), this will later need to be split into two mono channels
	bytesPerSample := 2 // 16 bits per sample
	samples := int(totalBytes) / (bytesPerSample * int(format.NumChannels))

	buf := &audio.IntBuffer{
		Data: make([]int, samples*int(format.NumChannels)),
		Format: &audio.Format{
			SampleRate:  int(format.SampleRate),
			NumChannels: int(format.NumChannels),
		},
	}
	if _, err := decoder.PCMBuffer(buf); err != nil && err != io.EOF {
		return nil, fmt.Errorf("error reading audio data: %v", err)
	}

	// Split the buffer into separate channels
	channelsData := make([][]byte, decoder.NumChans)
	for i := range channelsData {
		channelsData[i] = make([]byte, 0, len(buf.Data)/int(decoder.NumChans*2)) // *2 because LINEAR16 = 2 bytes per sample
	}

	for i := 0; i < len(buf.Data); i += int(decoder.NumChans) {
		for channel := 0; channel < int(decoder.NumChans); channel++ {
			// Convert each sample to bytes and append to the respective channel slice
			sample := buf.Data[i+channel]
			channelsData[channel] = append(channelsData[channel], byte(sample&0xFF), byte((sample>>8)&0xFF))
		}
	}

	return channelsData, nil
}

func streamAudioToServer(serverAddr string, audioBytes []byte) error {
	fmt.Println("connecting to server")
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("did not connect: %v", err)
		return err
	}
	fmt.Println("connected to server")
	defer conn.Close()

	// load from env
	five9TrustToken := os.Getenv("FIVE9_TRUST_TOKEN")
	md := metadata.New(map[string]string{
		"x-five9-trust-token": five9TrustToken,
	})
	// Create a new context with the above metadata
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	client := voicestream.NewVoiceClient(conn)

	stream, err := client.StreamingVoice(ctx)
	if err != nil {
		log.Fatalf("could not open stream: %v", err)
	}

	fmt.Println("streaming audio to server")

	// Assume streamingConfig is a pb.StreamingConfig object you've filled out
	err = stream.Send(&voicestream.StreamingVoiceRequest{
		StreamingRequest: &voicestream.StreamingVoiceRequest_StreamingConfig{
			StreamingConfig: &voicestream.StreamingConfig{
				VoiceConfig: &voicestream.VoiceConfig{
					Encoding:        voicestream.VoiceConfig_LINEAR16,
					SampleRateHertz: 8000,
				},
				VccCallId: "testCallId",
				DomainId:  "testDomainId",
				AgentId:   "testAgentId",
				// Hard coded dev campaign
				CampaignId: "1137786",
			},
		},
	})
	if err != nil {
		log.Fatalf("Failed to send streaming config: %v", err)
		return err
	}

	// Wait for the server to respond that we can start streaming
	if _, err := stream.Recv(); err != nil {
		log.Fatalf("Failed to receive server start signal: %v", err)
		return err
	}

	// Split the audioBytes into chunks and stream
	chunkSize := 1600 // For example, this might represent 20ms of audio at 8000Hz, 16-bit mono
	for i := 0; i < len(audioBytes); i += chunkSize {
		end := i + chunkSize
		if end > len(audioBytes) {
			end = len(audioBytes)
		}
		err = stream.Send(&voicestream.StreamingVoiceRequest{
			StreamingRequest: &voicestream.StreamingVoiceRequest_AudioContent{
				AudioContent: audioBytes[i:end],
			},
			SendTime: timestamppb.Now(),
		})
		if err != nil {
			log.Fatalf("Failed to send audio chunk: %v", err)
			return err
		}
	}

	fmt.Println("finished streaming audio to server")

	// Close the stream once all the audio has been sent
	if err := stream.CloseSend(); err != nil {
		log.Fatalf("Could not close stream: %v", err)
		return err
	}

	return nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if os.Getenv("FIVE9_TRUST_TOKEN") == "" {
		log.Fatal("FIVE9_TRUST_TOKEN is not set")
	}

	serverAddr := flag.String("server", "", "The server address in the format of host:port")
	audioFilePath := flag.String("audiofile", "", "Path to the audio file to stream. Must be a linear16 encoded WAV file with 8000 HZ sample rate")
	flag.Parse()

	if *serverAddr == "" {
		fmt.Println("Please provide the server address")
		return
	}
	if *audioFilePath == "" {
		fmt.Println("Please provide the path to the audio file to stream")
		return
	}

	audioChannels, err := loadAndProcessAudioFile(*audioFilePath)
	if err != nil {
		log.Fatalf("Failed to load and process audio file: %v", err)
	} else {
		log.Printf("Loaded %d audio channels", len(audioChannels))
	}

	var wg sync.WaitGroup

	for i, audioData := range audioChannels {
		wg.Add(1)
		go func(channelIndex int, data []byte) {
			fmt.Printf("Streaming audio for channel %d\n", channelIndex)
			defer wg.Done()
			if err := streamAudioToServer(*serverAddr, data); err != nil {
				fmt.Printf("Error streaming audio for channel %d: %v\n", channelIndex, err)
			}
			fmt.Printf("Finished streaming audio for channel %d\n", channelIndex)
		}(i, audioData)
	}

	wg.Wait()

	log.Println("All channels have been streamed successfully.")
}
