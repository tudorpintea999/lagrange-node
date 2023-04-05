// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: network/v1/network.proto

package types

import (
	types "github.com/Lagrange-Labs/Lagrange-Node/sequencer/types"
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

// JoinNetworkRequest is the request to join the network
type JoinNetworkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublicKey    string `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	StakeAddress string `protobuf:"bytes,2,opt,name=stake_address,json=stakeAddress,proto3" json:"stake_address,omitempty"`
	Signature    string `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *JoinNetworkRequest) Reset() {
	*x = JoinNetworkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_v1_network_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinNetworkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinNetworkRequest) ProtoMessage() {}

func (x *JoinNetworkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_network_v1_network_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinNetworkRequest.ProtoReflect.Descriptor instead.
func (*JoinNetworkRequest) Descriptor() ([]byte, []int) {
	return file_network_v1_network_proto_rawDescGZIP(), []int{0}
}

func (x *JoinNetworkRequest) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *JoinNetworkRequest) GetStakeAddress() string {
	if x != nil {
		return x.StakeAddress
	}
	return ""
}

func (x *JoinNetworkRequest) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

// JoinNetworkResponse is the response for joining the network
type JoinNetworkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result  bool   `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *JoinNetworkResponse) Reset() {
	*x = JoinNetworkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_v1_network_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinNetworkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinNetworkResponse) ProtoMessage() {}

func (x *JoinNetworkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_network_v1_network_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinNetworkResponse.ProtoReflect.Descriptor instead.
func (*JoinNetworkResponse) Descriptor() ([]byte, []int) {
	return file_network_v1_network_proto_rawDescGZIP(), []int{1}
}

func (x *JoinNetworkResponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

func (x *JoinNetworkResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// GetBlockRequest is the request to get the given block
type GetBlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockNumber uint64 `protobuf:"varint,1,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
}

func (x *GetBlockRequest) Reset() {
	*x = GetBlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_v1_network_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockRequest) ProtoMessage() {}

func (x *GetBlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_network_v1_network_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlockRequest.ProtoReflect.Descriptor instead.
func (*GetBlockRequest) Descriptor() ([]byte, []int) {
	return file_network_v1_network_proto_rawDescGZIP(), []int{2}
}

func (x *GetBlockRequest) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

// GetBlockResponse is the response for getting the given block
type GetBlockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Block *types.Block `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
}

func (x *GetBlockResponse) Reset() {
	*x = GetBlockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_v1_network_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockResponse) ProtoMessage() {}

func (x *GetBlockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_network_v1_network_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlockResponse.ProtoReflect.Descriptor instead.
func (*GetBlockResponse) Descriptor() ([]byte, []int) {
	return file_network_v1_network_proto_rawDescGZIP(), []int{3}
}

func (x *GetBlockResponse) GetBlock() *types.Block {
	if x != nil {
		return x.Block
	}
	return nil
}

// CommitBlockRequest is the request to commit the signature
type CommitBlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockNumber uint64 `protobuf:"varint,1,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	Signature   string `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *CommitBlockRequest) Reset() {
	*x = CommitBlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_v1_network_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitBlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitBlockRequest) ProtoMessage() {}

func (x *CommitBlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_network_v1_network_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitBlockRequest.ProtoReflect.Descriptor instead.
func (*CommitBlockRequest) Descriptor() ([]byte, []int) {
	return file_network_v1_network_proto_rawDescGZIP(), []int{4}
}

func (x *CommitBlockRequest) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *CommitBlockRequest) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

// CommitBlockResponse is the response for uploading the signature
type CommitBlockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result  bool   `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CommitBlockResponse) Reset() {
	*x = CommitBlockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_v1_network_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitBlockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitBlockResponse) ProtoMessage() {}

func (x *CommitBlockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_network_v1_network_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitBlockResponse.ProtoReflect.Descriptor instead.
func (*CommitBlockResponse) Descriptor() ([]byte, []int) {
	return file_network_v1_network_proto_rawDescGZIP(), []int{5}
}

func (x *CommitBlockResponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

func (x *CommitBlockResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_network_v1_network_proto protoreflect.FileDescriptor

var file_network_v1_network_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x1a, 0x18, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x76, 0x0a, 0x12, 0x4a, 0x6f, 0x69, 0x6e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x74,
	0x61, 0x6b, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x47, 0x0a, 0x13, 0x4a, 0x6f, 0x69, 0x6e,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x34, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x3d, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x71,
	0x75, 0x65, 0x6e, 0x63, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52,
	0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x55, 0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x47, 0x0a,
	0x13, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xf7, 0x01, 0x0a, 0x0e, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0b, 0x4a, 0x6f, 0x69,
	0x6e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x1e, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x08, 0x47, 0x65, 0x74,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1b, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4e, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12,
	0x1e, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1f, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c,
	0x61, 0x67, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x2d, 0x4c, 0x61, 0x62, 0x73, 0x2f, 0x4c, 0x61, 0x67,
	0x72, 0x61, 0x6e, 0x67, 0x65, 0x2d, 0x4e, 0x6f, 0x64, 0x65, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_network_v1_network_proto_rawDescOnce sync.Once
	file_network_v1_network_proto_rawDescData = file_network_v1_network_proto_rawDesc
)

func file_network_v1_network_proto_rawDescGZIP() []byte {
	file_network_v1_network_proto_rawDescOnce.Do(func() {
		file_network_v1_network_proto_rawDescData = protoimpl.X.CompressGZIP(file_network_v1_network_proto_rawDescData)
	})
	return file_network_v1_network_proto_rawDescData
}

var file_network_v1_network_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_network_v1_network_proto_goTypes = []interface{}{
	(*JoinNetworkRequest)(nil),  // 0: network.v1.JoinNetworkRequest
	(*JoinNetworkResponse)(nil), // 1: network.v1.JoinNetworkResponse
	(*GetBlockRequest)(nil),     // 2: network.v1.GetBlockRequest
	(*GetBlockResponse)(nil),    // 3: network.v1.GetBlockResponse
	(*CommitBlockRequest)(nil),  // 4: network.v1.CommitBlockRequest
	(*CommitBlockResponse)(nil), // 5: network.v1.CommitBlockResponse
	(*types.Block)(nil),         // 6: sequencer.v1.Block
}
var file_network_v1_network_proto_depIdxs = []int32{
	6, // 0: network.v1.GetBlockResponse.block:type_name -> sequencer.v1.Block
	0, // 1: network.v1.NetworkService.JoinNetwork:input_type -> network.v1.JoinNetworkRequest
	2, // 2: network.v1.NetworkService.GetBlock:input_type -> network.v1.GetBlockRequest
	4, // 3: network.v1.NetworkService.CommitBlock:input_type -> network.v1.CommitBlockRequest
	1, // 4: network.v1.NetworkService.JoinNetwork:output_type -> network.v1.JoinNetworkResponse
	3, // 5: network.v1.NetworkService.GetBlock:output_type -> network.v1.GetBlockResponse
	5, // 6: network.v1.NetworkService.CommitBlock:output_type -> network.v1.CommitBlockResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_network_v1_network_proto_init() }
func file_network_v1_network_proto_init() {
	if File_network_v1_network_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_network_v1_network_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinNetworkRequest); i {
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
		file_network_v1_network_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinNetworkResponse); i {
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
		file_network_v1_network_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlockRequest); i {
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
		file_network_v1_network_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlockResponse); i {
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
		file_network_v1_network_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitBlockRequest); i {
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
		file_network_v1_network_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitBlockResponse); i {
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
			RawDescriptor: file_network_v1_network_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_network_v1_network_proto_goTypes,
		DependencyIndexes: file_network_v1_network_proto_depIdxs,
		MessageInfos:      file_network_v1_network_proto_msgTypes,
	}.Build()
	File_network_v1_network_proto = out.File
	file_network_v1_network_proto_rawDesc = nil
	file_network_v1_network_proto_goTypes = nil
	file_network_v1_network_proto_depIdxs = nil
}
