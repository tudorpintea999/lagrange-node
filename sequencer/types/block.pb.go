// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: sequencer/v1/block.proto

package types

import (
	types "github.com/Lagrange-Labs/Lagrange-Node/synchronizer/types"
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

// BlockHeader is the block header structure
type BlockHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockNumber       uint64 `protobuf:"varint,1,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	ParentHash        string `protobuf:"bytes,2,opt,name=parent_hash,json=parentHash,proto3" json:"parent_hash,omitempty"`
	BlockHash         string `protobuf:"bytes,3,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
	ProposerPubKey    string `protobuf:"bytes,4,opt,name=proposer_pub_key,json=proposerPubKey,proto3" json:"proposer_pub_key,omitempty"`
	ProposerSignature string `protobuf:"bytes,5,opt,name=proposer_signature,json=proposerSignature,proto3" json:"proposer_signature,omitempty"`
}

func (x *BlockHeader) Reset() {
	*x = BlockHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sequencer_v1_block_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockHeader) ProtoMessage() {}

func (x *BlockHeader) ProtoReflect() protoreflect.Message {
	mi := &file_sequencer_v1_block_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockHeader.ProtoReflect.Descriptor instead.
func (*BlockHeader) Descriptor() ([]byte, []int) {
	return file_sequencer_v1_block_proto_rawDescGZIP(), []int{0}
}

func (x *BlockHeader) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *BlockHeader) GetParentHash() string {
	if x != nil {
		return x.ParentHash
	}
	return ""
}

func (x *BlockHeader) GetBlockHash() string {
	if x != nil {
		return x.BlockHash
	}
	return ""
}

func (x *BlockHeader) GetProposerPubKey() string {
	if x != nil {
		return x.ProposerPubKey
	}
	return ""
}

func (x *BlockHeader) GetProposerSignature() string {
	if x != nil {
		return x.ProposerSignature
	}
	return ""
}

// Block is the block body structure for zk proofs of the block delta
type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header    *BlockHeader      `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Delta     *types.BlockDelta `protobuf:"bytes,2,opt,name=delta,proto3" json:"delta,omitempty"`
	Proof     string            `protobuf:"bytes,3,opt,name=proof,proto3" json:"proof,omitempty"`
	Signature string            `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sequencer_v1_block_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_sequencer_v1_block_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_sequencer_v1_block_proto_rawDescGZIP(), []int{1}
}

func (x *Block) GetHeader() *BlockHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Block) GetDelta() *types.BlockDelta {
	if x != nil {
		return x.Delta
	}
	return nil
}

func (x *Block) GetProof() string {
	if x != nil {
		return x.Proof
	}
	return ""
}

func (x *Block) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

var File_sequencer_v1_block_proto protoreflect.FileDescriptor

var file_sequencer_v1_block_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72,
	0x6f, 0x6e, 0x69, 0x7a, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc9, 0x01, 0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x70,
	0x6f, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x75, 0x62, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x50, 0x75, 0x62, 0x4b,
	0x65, 0x79, 0x12, 0x2d, 0x0a, 0x12, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x5f, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x22, 0xa1, 0x01, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x31, 0x0a, 0x06, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65,
	0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x31,
	0x0a, 0x05, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x52, 0x05, 0x64, 0x65, 0x6c, 0x74,
	0x61, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x61, 0x67, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x2d, 0x4c, 0x61, 0x62,
	0x73, 0x2f, 0x4c, 0x61, 0x67, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x2d, 0x4e, 0x6f, 0x64, 0x65, 0x2f,
	0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x72, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sequencer_v1_block_proto_rawDescOnce sync.Once
	file_sequencer_v1_block_proto_rawDescData = file_sequencer_v1_block_proto_rawDesc
)

func file_sequencer_v1_block_proto_rawDescGZIP() []byte {
	file_sequencer_v1_block_proto_rawDescOnce.Do(func() {
		file_sequencer_v1_block_proto_rawDescData = protoimpl.X.CompressGZIP(file_sequencer_v1_block_proto_rawDescData)
	})
	return file_sequencer_v1_block_proto_rawDescData
}

var file_sequencer_v1_block_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sequencer_v1_block_proto_goTypes = []interface{}{
	(*BlockHeader)(nil),      // 0: sequencer.v1.BlockHeader
	(*Block)(nil),            // 1: sequencer.v1.Block
	(*types.BlockDelta)(nil), // 2: synchronizer.v1.BlockDelta
}
var file_sequencer_v1_block_proto_depIdxs = []int32{
	0, // 0: sequencer.v1.Block.header:type_name -> sequencer.v1.BlockHeader
	2, // 1: sequencer.v1.Block.delta:type_name -> synchronizer.v1.BlockDelta
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_sequencer_v1_block_proto_init() }
func file_sequencer_v1_block_proto_init() {
	if File_sequencer_v1_block_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sequencer_v1_block_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockHeader); i {
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
		file_sequencer_v1_block_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
			RawDescriptor: file_sequencer_v1_block_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sequencer_v1_block_proto_goTypes,
		DependencyIndexes: file_sequencer_v1_block_proto_depIdxs,
		MessageInfos:      file_sequencer_v1_block_proto_msgTypes,
	}.Build()
	File_sequencer_v1_block_proto = out.File
	file_sequencer_v1_block_proto_rawDesc = nil
	file_sequencer_v1_block_proto_goTypes = nil
	file_sequencer_v1_block_proto_depIdxs = nil
}
