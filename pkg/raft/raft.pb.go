// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: raft/raft.proto

package raft

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AppendEntryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term         uint64   `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	LeaderId     uint32   `protobuf:"varint,2,opt,name=leader_id,json=leaderId,proto3" json:"leader_id,omitempty"`
	PrevLogTerm  uint64   `protobuf:"varint,3,opt,name=prev_log_term,json=prevLogTerm,proto3" json:"prev_log_term,omitempty"`
	PrevLogIndex uint64   `protobuf:"varint,4,opt,name=prev_log_index,json=prevLogIndex,proto3" json:"prev_log_index,omitempty"`
	Entries      [][]byte `protobuf:"bytes,5,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *AppendEntryRequest) Reset() {
	*x = AppendEntryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raft_raft_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppendEntryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppendEntryRequest) ProtoMessage() {}

func (x *AppendEntryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_raft_raft_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppendEntryRequest.ProtoReflect.Descriptor instead.
func (*AppendEntryRequest) Descriptor() ([]byte, []int) {
	return file_raft_raft_proto_rawDescGZIP(), []int{0}
}

func (x *AppendEntryRequest) GetTerm() uint64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *AppendEntryRequest) GetLeaderId() uint32 {
	if x != nil {
		return x.LeaderId
	}
	return 0
}

func (x *AppendEntryRequest) GetPrevLogTerm() uint64 {
	if x != nil {
		return x.PrevLogTerm
	}
	return 0
}

func (x *AppendEntryRequest) GetPrevLogIndex() uint64 {
	if x != nil {
		return x.PrevLogIndex
	}
	return 0
}

func (x *AppendEntryRequest) GetEntries() [][]byte {
	if x != nil {
		return x.Entries
	}
	return nil
}

type AppendEntryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term    uint64 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	Success bool   `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *AppendEntryResponse) Reset() {
	*x = AppendEntryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raft_raft_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppendEntryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppendEntryResponse) ProtoMessage() {}

func (x *AppendEntryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_raft_raft_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppendEntryResponse.ProtoReflect.Descriptor instead.
func (*AppendEntryResponse) Descriptor() ([]byte, []int) {
	return file_raft_raft_proto_rawDescGZIP(), []int{1}
}

func (x *AppendEntryResponse) GetTerm() uint64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *AppendEntryResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type VoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term        uint64 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	CandidateId uint32 `protobuf:"varint,2,opt,name=candidate_id,json=candidateId,proto3" json:"candidate_id,omitempty"`
	LogTerm     uint64 `protobuf:"varint,3,opt,name=log_term,json=logTerm,proto3" json:"log_term,omitempty"`
	LogIndex    uint64 `protobuf:"varint,4,opt,name=log_index,json=logIndex,proto3" json:"log_index,omitempty"`
}

func (x *VoteRequest) Reset() {
	*x = VoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raft_raft_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteRequest) ProtoMessage() {}

func (x *VoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_raft_raft_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteRequest.ProtoReflect.Descriptor instead.
func (*VoteRequest) Descriptor() ([]byte, []int) {
	return file_raft_raft_proto_rawDescGZIP(), []int{2}
}

func (x *VoteRequest) GetTerm() uint64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *VoteRequest) GetCandidateId() uint32 {
	if x != nil {
		return x.CandidateId
	}
	return 0
}

func (x *VoteRequest) GetLogTerm() uint64 {
	if x != nil {
		return x.LogTerm
	}
	return 0
}

func (x *VoteRequest) GetLogIndex() uint64 {
	if x != nil {
		return x.LogIndex
	}
	return 0
}

type VoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term        uint64 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	VoteGranted bool   `protobuf:"varint,2,opt,name=vote_granted,json=voteGranted,proto3" json:"vote_granted,omitempty"`
}

func (x *VoteResponse) Reset() {
	*x = VoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raft_raft_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteResponse) ProtoMessage() {}

func (x *VoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_raft_raft_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteResponse.ProtoReflect.Descriptor instead.
func (*VoteResponse) Descriptor() ([]byte, []int) {
	return file_raft_raft_proto_rawDescGZIP(), []int{3}
}

func (x *VoteResponse) GetTerm() uint64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *VoteResponse) GetVoteGranted() bool {
	if x != nil {
		return x.VoteGranted
	}
	return false
}

var File_raft_raft_proto protoreflect.FileDescriptor

var file_raft_raft_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x72, 0x61, 0x66, 0x74, 0x2f, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x70, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x01,
	0x0a, 0x12, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6c, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x76, 0x5f, 0x6c, 0x6f,
	0x67, 0x5f, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x70, 0x72,
	0x65, 0x76, 0x4c, 0x6f, 0x67, 0x54, 0x65, 0x72, 0x6d, 0x12, 0x24, 0x0a, 0x0e, 0x70, 0x72, 0x65,
	0x76, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0c, 0x70, 0x72, 0x65, 0x76, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x18, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0c,
	0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22, 0x43, 0x0a, 0x13, 0x41, 0x70, 0x70,
	0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04,
	0x74, 0x65, 0x72, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x7c,
	0x0a, 0x0b, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x65, 0x72,
	0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x5f, 0x74, 0x65, 0x72, 0x6d,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x54, 0x65, 0x72, 0x6d, 0x12,
	0x1b, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x45, 0x0a, 0x0c,
	0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d,
	0x12, 0x21, 0x0a, 0x0c, 0x76, 0x6f, 0x74, 0x65, 0x5f, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x76, 0x6f, 0x74, 0x65, 0x47, 0x72, 0x61, 0x6e,
	0x74, 0x65, 0x64, 0x32, 0xf3, 0x01, 0x0a, 0x04, 0x52, 0x61, 0x66, 0x74, 0x12, 0x85, 0x01, 0x0a,
	0x0b, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1b, 0x2e, 0x72,
	0x61, 0x66, 0x74, 0x5f, 0x70, 0x62, 0x2e, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x72, 0x61, 0x66, 0x74,
	0x5f, 0x70, 0x62, 0x2e, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3b, 0x92, 0x41, 0x1b, 0x1a, 0x19, 0x41, 0x70,
	0x70, 0x65, 0x6e, 0x64, 0x20, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x20, 0x74, 0x6f, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x6c, 0x6f, 0x67, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a,
	0x22, 0x12, 0x2f, 0x72, 0x61, 0x66, 0x74, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x2d, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x63, 0x0a, 0x04, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x72,
	0x61, 0x66, 0x74, 0x5f, 0x70, 0x62, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x70, 0x62, 0x2e, 0x56, 0x6f, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2e, 0x92, 0x41, 0x16, 0x1a, 0x14,
	0x56, 0x6f, 0x74, 0x65, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x61, 0x20, 0x63, 0x61, 0x6e, 0x64, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x3a, 0x01, 0x2a, 0x22, 0x0a, 0x2f,
	0x72, 0x61, 0x66, 0x74, 0x2f, 0x76, 0x6f, 0x74, 0x65, 0x42, 0x76, 0x0a, 0x0b, 0x63, 0x6f, 0x6d,
	0x2e, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x70, 0x62, 0x42, 0x09, 0x52, 0x61, 0x66, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x65, 0x73, 0x63, 0x61, 0x6c, 0x6f, 0x70, 0x61, 0x2f, 0x72, 0x61, 0x66, 0x74, 0x2d,
	0x6b, 0x76, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x61, 0x66, 0x74, 0xa2, 0x02, 0x03, 0x52, 0x58,
	0x58, 0xaa, 0x02, 0x06, 0x52, 0x61, 0x66, 0x74, 0x50, 0x62, 0xca, 0x02, 0x06, 0x52, 0x61, 0x66,
	0x74, 0x50, 0x62, 0xe2, 0x02, 0x12, 0x52, 0x61, 0x66, 0x74, 0x50, 0x62, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x06, 0x52, 0x61, 0x66, 0x74, 0x50,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_raft_raft_proto_rawDescOnce sync.Once
	file_raft_raft_proto_rawDescData = file_raft_raft_proto_rawDesc
)

func file_raft_raft_proto_rawDescGZIP() []byte {
	file_raft_raft_proto_rawDescOnce.Do(func() {
		file_raft_raft_proto_rawDescData = protoimpl.X.CompressGZIP(file_raft_raft_proto_rawDescData)
	})
	return file_raft_raft_proto_rawDescData
}

var file_raft_raft_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_raft_raft_proto_goTypes = []interface{}{
	(*AppendEntryRequest)(nil),  // 0: raft_pb.AppendEntryRequest
	(*AppendEntryResponse)(nil), // 1: raft_pb.AppendEntryResponse
	(*VoteRequest)(nil),         // 2: raft_pb.VoteRequest
	(*VoteResponse)(nil),        // 3: raft_pb.VoteResponse
}
var file_raft_raft_proto_depIdxs = []int32{
	0, // 0: raft_pb.Raft.AppendEntry:input_type -> raft_pb.AppendEntryRequest
	2, // 1: raft_pb.Raft.Vote:input_type -> raft_pb.VoteRequest
	1, // 2: raft_pb.Raft.AppendEntry:output_type -> raft_pb.AppendEntryResponse
	3, // 3: raft_pb.Raft.Vote:output_type -> raft_pb.VoteResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_raft_raft_proto_init() }
func file_raft_raft_proto_init() {
	if File_raft_raft_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_raft_raft_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppendEntryRequest); i {
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
		file_raft_raft_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppendEntryResponse); i {
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
		file_raft_raft_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoteRequest); i {
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
		file_raft_raft_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoteResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_raft_raft_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_raft_raft_proto_goTypes,
		DependencyIndexes: file_raft_raft_proto_depIdxs,
		MessageInfos:      file_raft_raft_proto_msgTypes,
	}.Build()
	File_raft_raft_proto = out.File
	file_raft_raft_proto_rawDesc = nil
	file_raft_raft_proto_goTypes = nil
	file_raft_raft_proto_depIdxs = nil
}
