// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.0
// source: requestGameLJ.proto

package lider

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

type GameReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntryMsg string `protobuf:"bytes,1,opt,name=entryMsg,proto3" json:"entryMsg,omitempty"`
}

func (x *GameReq) Reset() {
	*x = GameReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_requestGameLJ_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameReq) ProtoMessage() {}

func (x *GameReq) ProtoReflect() protoreflect.Message {
	mi := &file_requestGameLJ_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameReq.ProtoReflect.Descriptor instead.
func (*GameReq) Descriptor() ([]byte, []int) {
	return file_requestGameLJ_proto_rawDescGZIP(), []int{0}
}

func (x *GameReq) GetEntryMsg() string {
	if x != nil {
		return x.EntryMsg
	}
	return ""
}

type GameResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameMsg string `protobuf:"bytes,1,opt,name=gameMsg,proto3" json:"gameMsg,omitempty"`
}

func (x *GameResp) Reset() {
	*x = GameResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_requestGameLJ_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameResp) ProtoMessage() {}

func (x *GameResp) ProtoReflect() protoreflect.Message {
	mi := &file_requestGameLJ_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameResp.ProtoReflect.Descriptor instead.
func (*GameResp) Descriptor() ([]byte, []int) {
	return file_requestGameLJ_proto_rawDescGZIP(), []int{1}
}

func (x *GameResp) GetGameMsg() string {
	if x != nil {
		return x.GameMsg
	}
	return ""
}

var File_requestGameLJ_proto protoreflect.FileDescriptor

var file_requestGameLJ_proto_rawDesc = []byte{
	0x0a, 0x13, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x4c, 0x4a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x47, 0x61,
	0x6d, 0x65, 0x4c, 0x4a, 0x22, 0x25, 0x0a, 0x07, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x12,
	0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x4d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x4d, 0x73, 0x67, 0x22, 0x24, 0x0a, 0x08, 0x47,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x61, 0x6d, 0x65, 0x4d,
	0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x61, 0x6d, 0x65, 0x4d, 0x73,
	0x67, 0x32, 0x54, 0x0a, 0x12, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x47, 0x61, 0x6d, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x47, 0x61, 0x6d, 0x65, 0x4c, 0x4a, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x17,
	0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x4c, 0x4a, 0x2e, 0x47,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x42, 0x0b, 0x5a, 0x09, 0x6c, 0x61, 0x62, 0x2f, 0x6c,
	0x69, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_requestGameLJ_proto_rawDescOnce sync.Once
	file_requestGameLJ_proto_rawDescData = file_requestGameLJ_proto_rawDesc
)

func file_requestGameLJ_proto_rawDescGZIP() []byte {
	file_requestGameLJ_proto_rawDescOnce.Do(func() {
		file_requestGameLJ_proto_rawDescData = protoimpl.X.CompressGZIP(file_requestGameLJ_proto_rawDescData)
	})
	return file_requestGameLJ_proto_rawDescData
}

var file_requestGameLJ_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_requestGameLJ_proto_goTypes = []interface{}{
	(*GameReq)(nil),  // 0: requestGameLJ.GameReq
	(*GameResp)(nil), // 1: requestGameLJ.GameResp
}
var file_requestGameLJ_proto_depIdxs = []int32{
	0, // 0: requestGameLJ.RequestGameService.RequestGame:input_type -> requestGameLJ.GameReq
	1, // 1: requestGameLJ.RequestGameService.RequestGame:output_type -> requestGameLJ.GameResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_requestGameLJ_proto_init() }
func file_requestGameLJ_proto_init() {
	if File_requestGameLJ_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_requestGameLJ_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameReq); i {
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
		file_requestGameLJ_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameResp); i {
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
			RawDescriptor: file_requestGameLJ_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_requestGameLJ_proto_goTypes,
		DependencyIndexes: file_requestGameLJ_proto_depIdxs,
		MessageInfos:      file_requestGameLJ_proto_msgTypes,
	}.Build()
	File_requestGameLJ_proto = out.File
	file_requestGameLJ_proto_rawDesc = nil
	file_requestGameLJ_proto_goTypes = nil
	file_requestGameLJ_proto_depIdxs = nil
}
