// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.0
// source: playerRecordNL.proto

package namenode

import (
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

type PlayerReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player string `protobuf:"bytes,1,opt,name=Player,proto3" json:"Player,omitempty"`
}

func (x *PlayerReq) Reset() {
	*x = PlayerReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playerRecordNL_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerReq) ProtoMessage() {}

func (x *PlayerReq) ProtoReflect() protoreflect.Message {
	mi := &file_playerRecordNL_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerReq.ProtoReflect.Descriptor instead.
func (*PlayerReq) Descriptor() ([]byte, []int) {
	return file_playerRecordNL_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerReq) GetPlayer() string {
	if x != nil {
		return x.Player
	}
	return ""
}

type PlayerResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Records string `protobuf:"bytes,1,opt,name=Records,proto3" json:"Records,omitempty"`
}

func (x *PlayerResp) Reset() {
	*x = PlayerResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playerRecordNL_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerResp) ProtoMessage() {}

func (x *PlayerResp) ProtoReflect() protoreflect.Message {
	mi := &file_playerRecordNL_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerResp.ProtoReflect.Descriptor instead.
func (*PlayerResp) Descriptor() ([]byte, []int) {
	return file_playerRecordNL_proto_rawDescGZIP(), []int{1}
}

func (x *PlayerResp) GetRecords() string {
	if x != nil {
		return x.Records
	}
	return ""
}

var File_playerRecordNL_proto protoreflect.FileDescriptor

var file_playerRecordNL_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4e, 0x4c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x4e, 0x4c, 0x22, 0x23, 0x0a, 0x09, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x22, 0x26, 0x0a, 0x0a, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x73, 0x32, 0x5c, 0x0a, 0x13, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0c, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x19, 0x2e, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4e, 0x4c, 0x2e, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x4e, 0x4c, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x42, 0x0e, 0x5a, 0x0c, 0x6c, 0x61, 0x62, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x6f, 0x64,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_playerRecordNL_proto_rawDescOnce sync.Once
	file_playerRecordNL_proto_rawDescData = file_playerRecordNL_proto_rawDesc
)

func file_playerRecordNL_proto_rawDescGZIP() []byte {
	file_playerRecordNL_proto_rawDescOnce.Do(func() {
		file_playerRecordNL_proto_rawDescData = protoimpl.X.CompressGZIP(file_playerRecordNL_proto_rawDescData)
	})
	return file_playerRecordNL_proto_rawDescData
}

var file_playerRecordNL_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_playerRecordNL_proto_goTypes = []interface{}{
	(*PlayerReq)(nil),  // 0: playerRecordNL.PlayerReq
	(*PlayerResp)(nil), // 1: playerRecordNL.PlayerResp
}
var file_playerRecordNL_proto_depIdxs = []int32{
	0, // 0: playerRecordNL.PlayerRecordService.PlayerRecord:input_type -> playerRecordNL.PlayerReq
	1, // 1: playerRecordNL.PlayerRecordService.PlayerRecord:output_type -> playerRecordNL.PlayerResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_playerRecordNL_proto_init() }
func file_playerRecordNL_proto_init() {
	if File_playerRecordNL_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_playerRecordNL_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerReq); i {
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
		file_playerRecordNL_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerResp); i {
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
			RawDescriptor: file_playerRecordNL_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_playerRecordNL_proto_goTypes,
		DependencyIndexes: file_playerRecordNL_proto_depIdxs,
		MessageInfos:      file_playerRecordNL_proto_msgTypes,
	}.Build()
	File_playerRecordNL_proto = out.File
	file_playerRecordNL_proto_rawDesc = nil
	file_playerRecordNL_proto_goTypes = nil
	file_playerRecordNL_proto_depIdxs = nil
}
