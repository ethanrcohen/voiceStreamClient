// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: api/voice.proto

package voicestream

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StreamingConfig_CallLeg int32

const (
	StreamingConfig_AGENT      StreamingConfig_CallLeg = 0
	StreamingConfig_CUSTOMER   StreamingConfig_CallLeg = 1
	StreamingConfig_SUPERVISOR StreamingConfig_CallLeg = 2
)

// Enum value maps for StreamingConfig_CallLeg.
var (
	StreamingConfig_CallLeg_name = map[int32]string{
		0: "AGENT",
		1: "CUSTOMER",
		2: "SUPERVISOR",
	}
	StreamingConfig_CallLeg_value = map[string]int32{
		"AGENT":      0,
		"CUSTOMER":   1,
		"SUPERVISOR": 2,
	}
)

func (x StreamingConfig_CallLeg) Enum() *StreamingConfig_CallLeg {
	p := new(StreamingConfig_CallLeg)
	*p = x
	return p
}

func (x StreamingConfig_CallLeg) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StreamingConfig_CallLeg) Descriptor() protoreflect.EnumDescriptor {
	return file_api_voice_proto_enumTypes[0].Descriptor()
}

func (StreamingConfig_CallLeg) Type() protoreflect.EnumType {
	return &file_api_voice_proto_enumTypes[0]
}

func (x StreamingConfig_CallLeg) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StreamingConfig_CallLeg.Descriptor instead.
func (StreamingConfig_CallLeg) EnumDescriptor() ([]byte, []int) {
	return file_api_voice_proto_rawDescGZIP(), []int{1, 0}
}

type VoiceConfig_AudioEncoding int32

const (
	// Not specified.
	VoiceConfig_ENCODING_UNSPECIFIED VoiceConfig_AudioEncoding = 0
	// Uncompressed 16-bit signed little-endian samples (Linear PCM).
	VoiceConfig_LINEAR16 VoiceConfig_AudioEncoding = 1
	// 8-bit samples that compand 14-bit audio samples using G.711 PCMU/mu-law.
	VoiceConfig_MULAW VoiceConfig_AudioEncoding = 2
	// Adaptive Multi-Rate Narrowband codec. `sample_rate_hertz` must be 8000.
	VoiceConfig_AMR VoiceConfig_AudioEncoding = 3
	// Adaptive Multi-Rate Wideband codec. `sample_rate_hertz` must be 16000.
	VoiceConfig_AMR_WB VoiceConfig_AudioEncoding = 4
)

// Enum value maps for VoiceConfig_AudioEncoding.
var (
	VoiceConfig_AudioEncoding_name = map[int32]string{
		0: "ENCODING_UNSPECIFIED",
		1: "LINEAR16",
		2: "MULAW",
		3: "AMR",
		4: "AMR_WB",
	}
	VoiceConfig_AudioEncoding_value = map[string]int32{
		"ENCODING_UNSPECIFIED": 0,
		"LINEAR16":             1,
		"MULAW":                2,
		"AMR":                  3,
		"AMR_WB":               4,
	}
)

func (x VoiceConfig_AudioEncoding) Enum() *VoiceConfig_AudioEncoding {
	p := new(VoiceConfig_AudioEncoding)
	*p = x
	return p
}

func (x VoiceConfig_AudioEncoding) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VoiceConfig_AudioEncoding) Descriptor() protoreflect.EnumDescriptor {
	return file_api_voice_proto_enumTypes[1].Descriptor()
}

func (VoiceConfig_AudioEncoding) Type() protoreflect.EnumType {
	return &file_api_voice_proto_enumTypes[1]
}

func (x VoiceConfig_AudioEncoding) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VoiceConfig_AudioEncoding.Descriptor instead.
func (VoiceConfig_AudioEncoding) EnumDescriptor() ([]byte, []int) {
	return file_api_voice_proto_rawDescGZIP(), []int{2, 0}
}

type StreamingStatus_StatusCode int32

const (
	StreamingStatus_SUCCESS StreamingStatus_StatusCode = 0
	// Client status codes
	StreamingStatus_CLT_CALL_ENDED        StreamingStatus_StatusCode = 1 // Call ended. Close the gRPC channel.
	StreamingStatus_CLT_CALL_HOLD         StreamingStatus_StatusCode = 2
	StreamingStatus_CLT_CALL_RESUME       StreamingStatus_StatusCode = 3
	StreamingStatus_CLT_DISCONNECT        StreamingStatus_StatusCode = 4 // Client closing gRPC channel.
	StreamingStatus_CLT_ERROR_NO_RESOURCE StreamingStatus_StatusCode = 100
	StreamingStatus_CLT_ERROR_TIMEOUT     StreamingStatus_StatusCode = 101
	StreamingStatus_CLT_ERROR_GENERIC     StreamingStatus_StatusCode = 102
	// Server status codes
	StreamingStatus_SRV_REQ_START_STREAMING StreamingStatus_StatusCode = 1001 // Start sending audio
	StreamingStatus_SRV_REQ_PAUSE_STREAMING StreamingStatus_StatusCode = 1002 // Stop sending audio
	StreamingStatus_SRV_REQ_DISCONNECT      StreamingStatus_StatusCode = 1011 // Close the existing gRPC channel.
	StreamingStatus_SRV_REQ_RECONNECT       StreamingStatus_StatusCode = 1012 // Close the existing channel and then start all over again on a new channel.
	StreamingStatus_SRV_ERROR_NO_RESOURCE   StreamingStatus_StatusCode = 1100
	StreamingStatus_SRV_ERROR_TIMEOUT       StreamingStatus_StatusCode = 1101
	StreamingStatus_SRV_ERROR_GENERIC       StreamingStatus_StatusCode = 1102
)

// Enum value maps for StreamingStatus_StatusCode.
var (
	StreamingStatus_StatusCode_name = map[int32]string{
		0:    "SUCCESS",
		1:    "CLT_CALL_ENDED",
		2:    "CLT_CALL_HOLD",
		3:    "CLT_CALL_RESUME",
		4:    "CLT_DISCONNECT",
		100:  "CLT_ERROR_NO_RESOURCE",
		101:  "CLT_ERROR_TIMEOUT",
		102:  "CLT_ERROR_GENERIC",
		1001: "SRV_REQ_START_STREAMING",
		1002: "SRV_REQ_PAUSE_STREAMING",
		1011: "SRV_REQ_DISCONNECT",
		1012: "SRV_REQ_RECONNECT",
		1100: "SRV_ERROR_NO_RESOURCE",
		1101: "SRV_ERROR_TIMEOUT",
		1102: "SRV_ERROR_GENERIC",
	}
	StreamingStatus_StatusCode_value = map[string]int32{
		"SUCCESS":                 0,
		"CLT_CALL_ENDED":          1,
		"CLT_CALL_HOLD":           2,
		"CLT_CALL_RESUME":         3,
		"CLT_DISCONNECT":          4,
		"CLT_ERROR_NO_RESOURCE":   100,
		"CLT_ERROR_TIMEOUT":       101,
		"CLT_ERROR_GENERIC":       102,
		"SRV_REQ_START_STREAMING": 1001,
		"SRV_REQ_PAUSE_STREAMING": 1002,
		"SRV_REQ_DISCONNECT":      1011,
		"SRV_REQ_RECONNECT":       1012,
		"SRV_ERROR_NO_RESOURCE":   1100,
		"SRV_ERROR_TIMEOUT":       1101,
		"SRV_ERROR_GENERIC":       1102,
	}
)

func (x StreamingStatus_StatusCode) Enum() *StreamingStatus_StatusCode {
	p := new(StreamingStatus_StatusCode)
	*p = x
	return p
}

func (x StreamingStatus_StatusCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StreamingStatus_StatusCode) Descriptor() protoreflect.EnumDescriptor {
	return file_api_voice_proto_enumTypes[2].Descriptor()
}

func (StreamingStatus_StatusCode) Type() protoreflect.EnumType {
	return &file_api_voice_proto_enumTypes[2]
}

func (x StreamingStatus_StatusCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StreamingStatus_StatusCode.Descriptor instead.
func (StreamingStatus_StatusCode) EnumDescriptor() ([]byte, []int) {
	return file_api_voice_proto_rawDescGZIP(), []int{3, 0}
}

// The message sent by the client for the 'StreamingVoice' method.
// Multiple 'StreamingVoiceRequest' messages are sent repeatedly as defined below.
// The first message must be 'streaming_config' containing control data specific to the call being streamed.
// The subsequent messages must be 'audio_content' with audio payload.
// After sending the 'streaming_config' message, the client must wait for a response from the server
// with status code SRV_START_STREAMING before sending audio payloads.
// Optionally status messages 'streaming_status' can be sent any time to provide
// additional information e.g events, notifications about the call.
type StreamingVoiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The streaming request, which is either a streaming config, audio content or client status.
	//
	// Types that are assignable to StreamingRequest:
	//	*StreamingVoiceRequest_StreamingConfig
	//	*StreamingVoiceRequest_AudioContent
	//	*StreamingVoiceRequest_StreamingStatus
	StreamingRequest isStreamingVoiceRequest_StreamingRequest `protobuf_oneof:"streaming_request"`
	// The time this message was created.
	// This must be set for messages with audio data. Optional for other type of messages.
	SendTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=send_time,json=sendTime,proto3" json:"send_time,omitempty"`
}

func (x *StreamingVoiceRequest) Reset() {
	*x = StreamingVoiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_voice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamingVoiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamingVoiceRequest) ProtoMessage() {}

func (x *StreamingVoiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_voice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamingVoiceRequest.ProtoReflect.Descriptor instead.
func (*StreamingVoiceRequest) Descriptor() ([]byte, []int) {
	return file_api_voice_proto_rawDescGZIP(), []int{0}
}

func (m *StreamingVoiceRequest) GetStreamingRequest() isStreamingVoiceRequest_StreamingRequest {
	if m != nil {
		return m.StreamingRequest
	}
	return nil
}

func (x *StreamingVoiceRequest) GetStreamingConfig() *StreamingConfig {
	if x, ok := x.GetStreamingRequest().(*StreamingVoiceRequest_StreamingConfig); ok {
		return x.StreamingConfig
	}
	return nil
}

func (x *StreamingVoiceRequest) GetAudioContent() []byte {
	if x, ok := x.GetStreamingRequest().(*StreamingVoiceRequest_AudioContent); ok {
		return x.AudioContent
	}
	return nil
}

func (x *StreamingVoiceRequest) GetStreamingStatus() *StreamingStatus {
	if x, ok := x.GetStreamingRequest().(*StreamingVoiceRequest_StreamingStatus); ok {
		return x.StreamingStatus
	}
	return nil
}

func (x *StreamingVoiceRequest) GetSendTime() *timestamppb.Timestamp {
	if x != nil {
		return x.SendTime
	}
	return nil
}

type isStreamingVoiceRequest_StreamingRequest interface {
	isStreamingVoiceRequest_StreamingRequest()
}

type StreamingVoiceRequest_StreamingConfig struct {
	// Provides information about the call and the participant to the receiver.
	// The first 'StreamingVoiceRequest' message must contain a 'StreamingConfig' message.
	StreamingConfig *StreamingConfig `protobuf:"bytes,1,opt,name=streaming_config,json=streamingConfig,proto3,oneof"`
}

type StreamingVoiceRequest_AudioContent struct {
	// As the user speaks the speech is sent in chunks of audio data
	// sequentially in a stream of 'StreamingVoiceRequest' messages.
	// The audio bytes must have been encoded as specified in 'StreamingConfig'.
	AudioContent []byte `protobuf:"bytes,2,opt,name=audio_content,json=audioContent,proto3,oneof"`
}

type StreamingVoiceRequest_StreamingStatus struct {
	// Provides additional information related to the call or stream e.g. events
	// like CALL_ENDED, HOLD, RESUME etc. Errors or statistics etc.
	StreamingStatus *StreamingStatus `protobuf:"bytes,3,opt,name=streaming_status,json=streamingStatus,proto3,oneof"`
}

func (*StreamingVoiceRequest_StreamingConfig) isStreamingVoiceRequest_StreamingRequest() {}

func (*StreamingVoiceRequest_AudioContent) isStreamingVoiceRequest_StreamingRequest() {}

func (*StreamingVoiceRequest_StreamingStatus) isStreamingVoiceRequest_StreamingRequest() {}

// Provides information to the receiver that specifies how to process the request.
type StreamingConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VoiceConfig *VoiceConfig `protobuf:"bytes,1,opt,name=voice_config,json=voiceConfig,proto3" json:"voice_config,omitempty"`
	// CallID to identify a call within a domain in Five9.
	VccCallId  string `protobuf:"bytes,2,opt,name=vcc_call_id,json=vccCallId,proto3" json:"vcc_call_id,omitempty"`
	DomainId   string `protobuf:"bytes,3,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	CampaignId string `protobuf:"bytes,4,opt,name=campaign_id,json=campaignId,proto3" json:"campaign_id,omitempty"`
	AgentId    string `protobuf:"bytes,5,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	// Identifies the role of the participant
	CallLeg    StreamingConfig_CallLeg `protobuf:"varint,6,opt,name=call_leg,json=callLeg,proto3,enum=five9.voicestream.StreamingConfig_CallLeg" json:"call_leg,omitempty"`
	TrustToken string                  `protobuf:"bytes,7,opt,name=trust_token,json=trustToken,proto3" json:"trust_token,omitempty"`
	// same call can be streamed for multiple subscribers/filters
	SubscriptionId string `protobuf:"bytes,8,opt,name=subscription_id,json=subscriptionId,proto3" json:"subscription_id,omitempty"`
	// Skill Id associated with this audio stream.
	SkillId string `protobuf:"bytes,9,opt,name=skill_id,json=skillId,proto3" json:"skill_id,omitempty"`
}

func (x *StreamingConfig) Reset() {
	*x = StreamingConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_voice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamingConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamingConfig) ProtoMessage() {}

func (x *StreamingConfig) ProtoReflect() protoreflect.Message {
	mi := &file_api_voice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamingConfig.ProtoReflect.Descriptor instead.
func (*StreamingConfig) Descriptor() ([]byte, []int) {
	return file_api_voice_proto_rawDescGZIP(), []int{1}
}

func (x *StreamingConfig) GetVoiceConfig() *VoiceConfig {
	if x != nil {
		return x.VoiceConfig
	}
	return nil
}

func (x *StreamingConfig) GetVccCallId() string {
	if x != nil {
		return x.VccCallId
	}
	return ""
}

func (x *StreamingConfig) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *StreamingConfig) GetCampaignId() string {
	if x != nil {
		return x.CampaignId
	}
	return ""
}

func (x *StreamingConfig) GetAgentId() string {
	if x != nil {
		return x.AgentId
	}
	return ""
}

func (x *StreamingConfig) GetCallLeg() StreamingConfig_CallLeg {
	if x != nil {
		return x.CallLeg
	}
	return StreamingConfig_AGENT
}

func (x *StreamingConfig) GetTrustToken() string {
	if x != nil {
		return x.TrustToken
	}
	return ""
}

func (x *StreamingConfig) GetSubscriptionId() string {
	if x != nil {
		return x.SubscriptionId
	}
	return ""
}

func (x *StreamingConfig) GetSkillId() string {
	if x != nil {
		return x.SkillId
	}
	return ""
}

// Provides information about audio data in the request.
type VoiceConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Encoding of data sent in 'StreamingVoiceRequest' messages with audio.
	Encoding VoiceConfig_AudioEncoding `protobuf:"varint,1,opt,name=encoding,proto3,enum=five9.voicestream.VoiceConfig_AudioEncoding" json:"encoding,omitempty"`
	// Sampling rate in Hertz of the audio data sent in 'StreamingVoiceRequest' messages.
	// Currently only 8000 is supported by VoiceStream.
	SampleRateHertz int32 `protobuf:"varint,2,opt,name=sample_rate_hertz,json=sampleRateHertz,proto3" json:"sample_rate_hertz,omitempty"`
}

func (x *VoiceConfig) Reset() {
	*x = VoiceConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_voice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoiceConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoiceConfig) ProtoMessage() {}

func (x *VoiceConfig) ProtoReflect() protoreflect.Message {
	mi := &file_api_voice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoiceConfig.ProtoReflect.Descriptor instead.
func (*VoiceConfig) Descriptor() ([]byte, []int) {
	return file_api_voice_proto_rawDescGZIP(), []int{2}
}

func (x *VoiceConfig) GetEncoding() VoiceConfig_AudioEncoding {
	if x != nil {
		return x.Encoding
	}
	return VoiceConfig_ENCODING_UNSPECIFIED
}

func (x *VoiceConfig) GetSampleRateHertz() int32 {
	if x != nil {
		return x.SampleRateHertz
	}
	return 0
}

// The status message which can be used by either client or server
// in the Request or Response message respectively.
type StreamingStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The status code, which is either a client side code or server side code.
	Code StreamingStatus_StatusCode `protobuf:"varint,1,opt,name=code,proto3,enum=five9.voicestream.StreamingStatus_StatusCode" json:"code,omitempty"`
	// A description of the status
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *StreamingStatus) Reset() {
	*x = StreamingStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_voice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamingStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamingStatus) ProtoMessage() {}

func (x *StreamingStatus) ProtoReflect() protoreflect.Message {
	mi := &file_api_voice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamingStatus.ProtoReflect.Descriptor instead.
func (*StreamingStatus) Descriptor() ([]byte, []int) {
	return file_api_voice_proto_rawDescGZIP(), []int{3}
}

func (x *StreamingStatus) GetCode() StreamingStatus_StatusCode {
	if x != nil {
		return x.Code
	}
	return StreamingStatus_SUCCESS
}

func (x *StreamingStatus) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// The message sent by the server for the 'StreamingVoice' method.
// The server may send status or provide feedback about the call using this message.
// The first message must be with status code SRV_START_STREAMING, and it must be
// sent after receiving configuration 'streaming_config' in request message from the client.
type StreamingVoiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Provides notifications e.g. events, errors etc about the stream to the client.
	Status *StreamingStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// Optional: Provides additional data e.g. feedback about the stream to the client.
	Feedback *StreamingFeedback `protobuf:"bytes,2,opt,name=feedback,proto3" json:"feedback,omitempty"`
}

func (x *StreamingVoiceResponse) Reset() {
	*x = StreamingVoiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_voice_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamingVoiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamingVoiceResponse) ProtoMessage() {}

func (x *StreamingVoiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_voice_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamingVoiceResponse.ProtoReflect.Descriptor instead.
func (*StreamingVoiceResponse) Descriptor() ([]byte, []int) {
	return file_api_voice_proto_rawDescGZIP(), []int{4}
}

func (x *StreamingVoiceResponse) GetStatus() *StreamingStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *StreamingVoiceResponse) GetFeedback() *StreamingFeedback {
	if x != nil {
		return x.Feedback
	}
	return nil
}

type StreamingFeedback struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VoiceConfig *VoiceConfig `protobuf:"bytes,1,opt,name=voice_config,json=voiceConfig,proto3" json:"voice_config,omitempty"`
}

func (x *StreamingFeedback) Reset() {
	*x = StreamingFeedback{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_voice_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamingFeedback) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamingFeedback) ProtoMessage() {}

func (x *StreamingFeedback) ProtoReflect() protoreflect.Message {
	mi := &file_api_voice_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamingFeedback.ProtoReflect.Descriptor instead.
func (*StreamingFeedback) Descriptor() ([]byte, []int) {
	return file_api_voice_proto_rawDescGZIP(), []int{5}
}

func (x *StreamingFeedback) GetVoiceConfig() *VoiceConfig {
	if x != nil {
		return x.VoiceConfig
	}
	return nil
}

var File_api_voice_proto protoreflect.FileDescriptor

var file_api_voice_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x11, 0x66, 0x69, 0x76, 0x65, 0x39, 0x2e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xae, 0x02, 0x0a, 0x15, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x69, 0x6e, 0x67, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x4f, 0x0a, 0x10, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x66, 0x69, 0x76, 0x65,
	0x39, 0x2e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52,
	0x0f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x25, 0x0a, 0x0d, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x0c, 0x61, 0x75, 0x64, 0x69, 0x6f,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x4f, 0x0a, 0x10, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x66, 0x69, 0x76, 0x65, 0x39, 0x2e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x00, 0x52, 0x0f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69,
	0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x37, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x42, 0x13, 0x0a, 0x11, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x5f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xad, 0x03, 0x0a, 0x0f, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x41, 0x0a, 0x0c, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x66, 0x69, 0x76, 0x65, 0x39, 0x2e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x2e, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x0b, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1e, 0x0a,
	0x0b, 0x76, 0x63, 0x63, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x76, 0x63, 0x63, 0x43, 0x61, 0x6c, 0x6c, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61,
	0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x45, 0x0a, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x6c,
	0x65, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x66, 0x69, 0x76, 0x65, 0x39,
	0x2e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x61, 0x6c,
	0x6c, 0x4c, 0x65, 0x67, 0x52, 0x07, 0x63, 0x61, 0x6c, 0x6c, 0x4c, 0x65, 0x67, 0x12, 0x1f, 0x0a,
	0x0b, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x74, 0x72, 0x75, 0x73, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x27,
	0x0a, 0x0f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x6b, 0x69, 0x6c, 0x6c,
	0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6b, 0x69, 0x6c, 0x6c,
	0x49, 0x64, 0x22, 0x32, 0x0a, 0x07, 0x43, 0x61, 0x6c, 0x6c, 0x4c, 0x65, 0x67, 0x12, 0x09, 0x0a,
	0x05, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x55, 0x53, 0x54,
	0x4f, 0x4d, 0x45, 0x52, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x55, 0x50, 0x45, 0x52, 0x56,
	0x49, 0x53, 0x4f, 0x52, 0x10, 0x02, 0x22, 0xdc, 0x01, 0x0a, 0x0b, 0x56, 0x6f, 0x69, 0x63, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x48, 0x0a, 0x08, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69,
	0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x66, 0x69, 0x76, 0x65, 0x39,
	0x2e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x56, 0x6f, 0x69,
	0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x45, 0x6e,
	0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67,
	0x12, 0x2a, 0x0a, 0x11, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f,
	0x68, 0x65, 0x72, 0x74, 0x7a, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x73, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x52, 0x61, 0x74, 0x65, 0x48, 0x65, 0x72, 0x74, 0x7a, 0x22, 0x57, 0x0a, 0x0d,
	0x41, 0x75, 0x64, 0x69, 0x6f, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a,
	0x14, 0x45, 0x4e, 0x43, 0x4f, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x4c, 0x49, 0x4e, 0x45, 0x41,
	0x52, 0x31, 0x36, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x55, 0x4c, 0x41, 0x57, 0x10, 0x02,
	0x12, 0x07, 0x0a, 0x03, 0x41, 0x4d, 0x52, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x4d, 0x52,
	0x5f, 0x57, 0x42, 0x10, 0x04, 0x22, 0xdc, 0x03, 0x0a, 0x0f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x41, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x66, 0x69, 0x76, 0x65, 0x39, 0x2e,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xeb, 0x02, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53,
	0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x4c, 0x54, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x45,
	0x4e, 0x44, 0x45, 0x44, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x4c, 0x54, 0x5f, 0x43, 0x41,
	0x4c, 0x4c, 0x5f, 0x48, 0x4f, 0x4c, 0x44, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x4c, 0x54,
	0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4d, 0x45, 0x10, 0x03, 0x12, 0x12,
	0x0a, 0x0e, 0x43, 0x4c, 0x54, 0x5f, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54,
	0x10, 0x04, 0x12, 0x19, 0x0a, 0x15, 0x43, 0x4c, 0x54, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f,
	0x4e, 0x4f, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x10, 0x64, 0x12, 0x15, 0x0a,
	0x11, 0x43, 0x4c, 0x54, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f,
	0x55, 0x54, 0x10, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x4c, 0x54, 0x5f, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x5f, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x49, 0x43, 0x10, 0x66, 0x12, 0x1c, 0x0a, 0x17, 0x53,
	0x52, 0x56, 0x5f, 0x52, 0x45, 0x51, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x5f, 0x53, 0x54, 0x52,
	0x45, 0x41, 0x4d, 0x49, 0x4e, 0x47, 0x10, 0xe9, 0x07, 0x12, 0x1c, 0x0a, 0x17, 0x53, 0x52, 0x56,
	0x5f, 0x52, 0x45, 0x51, 0x5f, 0x50, 0x41, 0x55, 0x53, 0x45, 0x5f, 0x53, 0x54, 0x52, 0x45, 0x41,
	0x4d, 0x49, 0x4e, 0x47, 0x10, 0xea, 0x07, 0x12, 0x17, 0x0a, 0x12, 0x53, 0x52, 0x56, 0x5f, 0x52,
	0x45, 0x51, 0x5f, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x10, 0xf3, 0x07,
	0x12, 0x16, 0x0a, 0x11, 0x53, 0x52, 0x56, 0x5f, 0x52, 0x45, 0x51, 0x5f, 0x52, 0x45, 0x43, 0x4f,
	0x4e, 0x4e, 0x45, 0x43, 0x54, 0x10, 0xf4, 0x07, 0x12, 0x1a, 0x0a, 0x15, 0x53, 0x52, 0x56, 0x5f,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x4e, 0x4f, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43,
	0x45, 0x10, 0xcc, 0x08, 0x12, 0x16, 0x0a, 0x11, 0x53, 0x52, 0x56, 0x5f, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0xcd, 0x08, 0x12, 0x16, 0x0a, 0x11,
	0x53, 0x52, 0x56, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x49,
	0x43, 0x10, 0xce, 0x08, 0x22, 0x96, 0x01, 0x0a, 0x16, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69,
	0x6e, 0x67, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x66, 0x69, 0x76, 0x65, 0x39, 0x2e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x40, 0x0a, 0x08, 0x66,
	0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x66, 0x69, 0x76, 0x65, 0x39, 0x2e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x46, 0x65, 0x65, 0x64, 0x62,
	0x61, 0x63, 0x6b, 0x52, 0x08, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x22, 0x56, 0x0a,
	0x11, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61,
	0x63, 0x6b, 0x12, 0x41, 0x0a, 0x0c, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x66, 0x69, 0x76, 0x65, 0x39,
	0x2e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x56, 0x6f, 0x69,
	0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0b, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x32, 0x74, 0x0a, 0x05, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x6b,
	0x0a, 0x0e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x56, 0x6f, 0x69, 0x63, 0x65,
	0x12, 0x28, 0x2e, 0x66, 0x69, 0x76, 0x65, 0x39, 0x2e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x56, 0x6f,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x66, 0x69, 0x76,
	0x65, 0x39, 0x2e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x23, 0x5a, 0x21, 0x63,
	0x6f, 0x6d, 0x2e, 0x66, 0x69, 0x76, 0x65, 0x39, 0x2e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x2f, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_voice_proto_rawDescOnce sync.Once
	file_api_voice_proto_rawDescData = file_api_voice_proto_rawDesc
)

func file_api_voice_proto_rawDescGZIP() []byte {
	file_api_voice_proto_rawDescOnce.Do(func() {
		file_api_voice_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_voice_proto_rawDescData)
	})
	return file_api_voice_proto_rawDescData
}

var file_api_voice_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_api_voice_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_voice_proto_goTypes = []interface{}{
	(StreamingConfig_CallLeg)(0),    // 0: five9.voicestream.StreamingConfig.CallLeg
	(VoiceConfig_AudioEncoding)(0),  // 1: five9.voicestream.VoiceConfig.AudioEncoding
	(StreamingStatus_StatusCode)(0), // 2: five9.voicestream.StreamingStatus.StatusCode
	(*StreamingVoiceRequest)(nil),   // 3: five9.voicestream.StreamingVoiceRequest
	(*StreamingConfig)(nil),         // 4: five9.voicestream.StreamingConfig
	(*VoiceConfig)(nil),             // 5: five9.voicestream.VoiceConfig
	(*StreamingStatus)(nil),         // 6: five9.voicestream.StreamingStatus
	(*StreamingVoiceResponse)(nil),  // 7: five9.voicestream.StreamingVoiceResponse
	(*StreamingFeedback)(nil),       // 8: five9.voicestream.StreamingFeedback
	(*timestamppb.Timestamp)(nil),   // 9: google.protobuf.Timestamp
}
var file_api_voice_proto_depIdxs = []int32{
	4,  // 0: five9.voicestream.StreamingVoiceRequest.streaming_config:type_name -> five9.voicestream.StreamingConfig
	6,  // 1: five9.voicestream.StreamingVoiceRequest.streaming_status:type_name -> five9.voicestream.StreamingStatus
	9,  // 2: five9.voicestream.StreamingVoiceRequest.send_time:type_name -> google.protobuf.Timestamp
	5,  // 3: five9.voicestream.StreamingConfig.voice_config:type_name -> five9.voicestream.VoiceConfig
	0,  // 4: five9.voicestream.StreamingConfig.call_leg:type_name -> five9.voicestream.StreamingConfig.CallLeg
	1,  // 5: five9.voicestream.VoiceConfig.encoding:type_name -> five9.voicestream.VoiceConfig.AudioEncoding
	2,  // 6: five9.voicestream.StreamingStatus.code:type_name -> five9.voicestream.StreamingStatus.StatusCode
	6,  // 7: five9.voicestream.StreamingVoiceResponse.status:type_name -> five9.voicestream.StreamingStatus
	8,  // 8: five9.voicestream.StreamingVoiceResponse.feedback:type_name -> five9.voicestream.StreamingFeedback
	5,  // 9: five9.voicestream.StreamingFeedback.voice_config:type_name -> five9.voicestream.VoiceConfig
	3,  // 10: five9.voicestream.Voice.StreamingVoice:input_type -> five9.voicestream.StreamingVoiceRequest
	7,  // 11: five9.voicestream.Voice.StreamingVoice:output_type -> five9.voicestream.StreamingVoiceResponse
	11, // [11:12] is the sub-list for method output_type
	10, // [10:11] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_api_voice_proto_init() }
func file_api_voice_proto_init() {
	if File_api_voice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_voice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamingVoiceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_voice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamingConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_voice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoiceConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_voice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamingStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_voice_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamingVoiceResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_voice_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamingFeedback); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_api_voice_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*StreamingVoiceRequest_StreamingConfig)(nil),
		(*StreamingVoiceRequest_AudioContent)(nil),
		(*StreamingVoiceRequest_StreamingStatus)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_voice_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_voice_proto_goTypes,
		DependencyIndexes: file_api_voice_proto_depIdxs,
		EnumInfos:         file_api_voice_proto_enumTypes,
		MessageInfos:      file_api_voice_proto_msgTypes,
	}.Build()
	File_api_voice_proto = out.File
	file_api_voice_proto_rawDesc = nil
	file_api_voice_proto_goTypes = nil
	file_api_voice_proto_depIdxs = nil
}
