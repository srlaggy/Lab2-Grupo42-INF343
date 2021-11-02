// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.0
// source: sendPlaysLJ.proto

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

type NumPasosReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayMsg int64 `protobuf:"varint,1,opt,name=PlayMsg,proto3" json:"PlayMsg,omitempty"`
}

func (x *NumPasosReq) Reset() {
	*x = NumPasosReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sendPlaysLJ_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NumPasosReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumPasosReq) ProtoMessage() {}

func (x *NumPasosReq) ProtoReflect() protoreflect.Message {
	mi := &file_sendPlaysLJ_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NumPasosReq.ProtoReflect.Descriptor instead.
func (*NumPasosReq) Descriptor() ([]byte, []int) {
	return file_sendPlaysLJ_proto_rawDescGZIP(), []int{0}
}

func (x *NumPasosReq) GetPlayMsg() int64 {
	if x != nil {
		return x.PlayMsg
	}
	return 0
}

type NumPasosResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StateMsg int64 `protobuf:"varint,1,opt,name=stateMsg,proto3" json:"stateMsg,omitempty"`
}

func (x *NumPasosResp) Reset() {
	*x = NumPasosResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sendPlaysLJ_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NumPasosResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumPasosResp) ProtoMessage() {}

func (x *NumPasosResp) ProtoReflect() protoreflect.Message {
	mi := &file_sendPlaysLJ_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NumPasosResp.ProtoReflect.Descriptor instead.
func (*NumPasosResp) Descriptor() ([]byte, []int) {
	return file_sendPlaysLJ_proto_rawDescGZIP(), []int{1}
}

func (x *NumPasosResp) GetStateMsg() int64 {
	if x != nil {
		return x.StateMsg
	}
	return 0
}

var File_sendPlaysLJ_proto protoreflect.FileDescriptor

var file_sendPlaysLJ_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x65, 0x6e, 0x64, 0x50, 0x6c, 0x61, 0x79, 0x73, 0x4c, 0x4a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x65, 0x6e, 0x64, 0x50, 0x6c, 0x61, 0x79, 0x73, 0x4c, 0x4a,
	0x22, 0x27, 0x0a, 0x0b, 0x4e, 0x75, 0x6d, 0x50, 0x61, 0x73, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x12,
	0x18, 0x0a, 0x07, 0x50, 0x6c, 0x61, 0x79, 0x4d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x50, 0x6c, 0x61, 0x79, 0x4d, 0x73, 0x67, 0x22, 0x2a, 0x0a, 0x0c, 0x4e, 0x75, 0x6d,
	0x50, 0x61, 0x73, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x4d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x4d, 0x73, 0x67, 0x32, 0x54, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x50, 0x61, 0x73,
	0x6f, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x09, 0x53, 0x65, 0x6e,
	0x64, 0x50, 0x61, 0x73, 0x6f, 0x73, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x6e, 0x64, 0x50, 0x6c, 0x61,
	0x79, 0x73, 0x4c, 0x4a, 0x2e, 0x4e, 0x75, 0x6d, 0x50, 0x61, 0x73, 0x6f, 0x73, 0x52, 0x65, 0x71,
	0x1a, 0x19, 0x2e, 0x73, 0x65, 0x6e, 0x64, 0x50, 0x6c, 0x61, 0x79, 0x73, 0x4c, 0x4a, 0x2e, 0x4e,
	0x75, 0x6d, 0x50, 0x61, 0x73, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x42, 0x0b, 0x5a, 0x09, 0x6c,
	0x61, 0x62, 0x2f, 0x6c, 0x69, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sendPlaysLJ_proto_rawDescOnce sync.Once
	file_sendPlaysLJ_proto_rawDescData = file_sendPlaysLJ_proto_rawDesc
)

func file_sendPlaysLJ_proto_rawDescGZIP() []byte {
	file_sendPlaysLJ_proto_rawDescOnce.Do(func() {
		file_sendPlaysLJ_proto_rawDescData = protoimpl.X.CompressGZIP(file_sendPlaysLJ_proto_rawDescData)
	})
	return file_sendPlaysLJ_proto_rawDescData
}

var file_sendPlaysLJ_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sendPlaysLJ_proto_goTypes = []interface{}{
	(*NumPasosReq)(nil),  // 0: sendPlaysLJ.NumPasosReq
	(*NumPasosResp)(nil), // 1: sendPlaysLJ.NumPasosResp
}
var file_sendPlaysLJ_proto_depIdxs = []int32{
	0, // 0: sendPlaysLJ.SendPasosService.SendPasos:input_type -> sendPlaysLJ.NumPasosReq
	1, // 1: sendPlaysLJ.SendPasosService.SendPasos:output_type -> sendPlaysLJ.NumPasosResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sendPlaysLJ_proto_init() }
func file_sendPlaysLJ_proto_init() {
	if File_sendPlaysLJ_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sendPlaysLJ_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NumPasosReq); i {
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
		file_sendPlaysLJ_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NumPasosResp); i {
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
			RawDescriptor: file_sendPlaysLJ_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sendPlaysLJ_proto_goTypes,
		DependencyIndexes: file_sendPlaysLJ_proto_depIdxs,
		MessageInfos:      file_sendPlaysLJ_proto_msgTypes,
	}.Build()
	File_sendPlaysLJ_proto = out.File
	file_sendPlaysLJ_proto_rawDesc = nil
	file_sendPlaysLJ_proto_goTypes = nil
	file_sendPlaysLJ_proto_depIdxs = nil
}
