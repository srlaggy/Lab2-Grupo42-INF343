// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.0
// source: LJ.proto

package jugador

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
		mi := &file_LJ_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameReq) ProtoMessage() {}

func (x *GameReq) ProtoReflect() protoreflect.Message {
	mi := &file_LJ_proto_msgTypes[0]
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
	return file_LJ_proto_rawDescGZIP(), []int{0}
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

	GameMsg    string `protobuf:"bytes,1,opt,name=gameMsg,proto3" json:"gameMsg,omitempty"`
	NroJugador int64  `protobuf:"varint,2,opt,name=nroJugador,proto3" json:"nroJugador,omitempty"`
}

func (x *GameResp) Reset() {
	*x = GameResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LJ_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameResp) ProtoMessage() {}

func (x *GameResp) ProtoReflect() protoreflect.Message {
	mi := &file_LJ_proto_msgTypes[1]
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
	return file_LJ_proto_rawDescGZIP(), []int{1}
}

func (x *GameResp) GetGameMsg() string {
	if x != nil {
		return x.GameMsg
	}
	return ""
}

func (x *GameResp) GetNroJugador() int64 {
	if x != nil {
		return x.NroJugador
	}
	return 0
}

type NumPasosReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayMsg    int64 `protobuf:"varint,1,opt,name=PlayMsg,proto3" json:"PlayMsg,omitempty"`
	NroJugador int64 `protobuf:"varint,2,opt,name=nroJugador,proto3" json:"nroJugador,omitempty"`
	Total      int64 `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	Ronda      int64 `protobuf:"varint,4,opt,name=ronda,proto3" json:"ronda,omitempty"`
}

func (x *NumPasosReq) Reset() {
	*x = NumPasosReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LJ_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NumPasosReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumPasosReq) ProtoMessage() {}

func (x *NumPasosReq) ProtoReflect() protoreflect.Message {
	mi := &file_LJ_proto_msgTypes[2]
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
	return file_LJ_proto_rawDescGZIP(), []int{2}
}

func (x *NumPasosReq) GetPlayMsg() int64 {
	if x != nil {
		return x.PlayMsg
	}
	return 0
}

func (x *NumPasosReq) GetNroJugador() int64 {
	if x != nil {
		return x.NroJugador
	}
	return 0
}

func (x *NumPasosReq) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *NumPasosReq) GetRonda() int64 {
	if x != nil {
		return x.Ronda
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
		mi := &file_LJ_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NumPasosResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumPasosResp) ProtoMessage() {}

func (x *NumPasosResp) ProtoReflect() protoreflect.Message {
	mi := &file_LJ_proto_msgTypes[3]
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
	return file_LJ_proto_rawDescGZIP(), []int{3}
}

func (x *NumPasosResp) GetStateMsg() int64 {
	if x != nil {
		return x.StateMsg
	}
	return 0
}

type E2ConnReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NroJugador int64 `protobuf:"varint,1,opt,name=nroJugador,proto3" json:"nroJugador,omitempty"`
}

func (x *E2ConnReq) Reset() {
	*x = E2ConnReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LJ_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *E2ConnReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*E2ConnReq) ProtoMessage() {}

func (x *E2ConnReq) ProtoReflect() protoreflect.Message {
	mi := &file_LJ_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use E2ConnReq.ProtoReflect.Descriptor instead.
func (*E2ConnReq) Descriptor() ([]byte, []int) {
	return file_LJ_proto_rawDescGZIP(), []int{4}
}

func (x *E2ConnReq) GetNroJugador() int64 {
	if x != nil {
		return x.NroJugador
	}
	return 0
}

type E2ConnResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NroGroup int64 `protobuf:"varint,1,opt,name=nroGroup,proto3" json:"nroGroup,omitempty"`
}

func (x *E2ConnResp) Reset() {
	*x = E2ConnResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LJ_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *E2ConnResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*E2ConnResp) ProtoMessage() {}

func (x *E2ConnResp) ProtoReflect() protoreflect.Message {
	mi := &file_LJ_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use E2ConnResp.ProtoReflect.Descriptor instead.
func (*E2ConnResp) Descriptor() ([]byte, []int) {
	return file_LJ_proto_rawDescGZIP(), []int{5}
}

func (x *E2ConnResp) GetNroGroup() int64 {
	if x != nil {
		return x.NroGroup
	}
	return 0
}

type Etapa2Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NroJugador int64 `protobuf:"varint,1,opt,name=nroJugador,proto3" json:"nroJugador,omitempty"`
	Numero     int64 `protobuf:"varint,2,opt,name=numero,proto3" json:"numero,omitempty"`
}

func (x *Etapa2Req) Reset() {
	*x = Etapa2Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LJ_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Etapa2Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Etapa2Req) ProtoMessage() {}

func (x *Etapa2Req) ProtoReflect() protoreflect.Message {
	mi := &file_LJ_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Etapa2Req.ProtoReflect.Descriptor instead.
func (*Etapa2Req) Descriptor() ([]byte, []int) {
	return file_LJ_proto_rawDescGZIP(), []int{6}
}

func (x *Etapa2Req) GetNroJugador() int64 {
	if x != nil {
		return x.NroJugador
	}
	return 0
}

func (x *Etapa2Req) GetNumero() int64 {
	if x != nil {
		return x.Numero
	}
	return 0
}

type Etapa2Resp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StateMsg int64 `protobuf:"varint,1,opt,name=stateMsg,proto3" json:"stateMsg,omitempty"`
}

func (x *Etapa2Resp) Reset() {
	*x = Etapa2Resp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LJ_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Etapa2Resp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Etapa2Resp) ProtoMessage() {}

func (x *Etapa2Resp) ProtoReflect() protoreflect.Message {
	mi := &file_LJ_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Etapa2Resp.ProtoReflect.Descriptor instead.
func (*Etapa2Resp) Descriptor() ([]byte, []int) {
	return file_LJ_proto_rawDescGZIP(), []int{7}
}

func (x *Etapa2Resp) GetStateMsg() int64 {
	if x != nil {
		return x.StateMsg
	}
	return 0
}

type E2FinReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NroJugador int64 `protobuf:"varint,1,opt,name=nroJugador,proto3" json:"nroJugador,omitempty"`
	NroGroup   int64 `protobuf:"varint,2,opt,name=nroGroup,proto3" json:"nroGroup,omitempty"`
}

func (x *E2FinReq) Reset() {
	*x = E2FinReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LJ_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *E2FinReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*E2FinReq) ProtoMessage() {}

func (x *E2FinReq) ProtoReflect() protoreflect.Message {
	mi := &file_LJ_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use E2FinReq.ProtoReflect.Descriptor instead.
func (*E2FinReq) Descriptor() ([]byte, []int) {
	return file_LJ_proto_rawDescGZIP(), []int{8}
}

func (x *E2FinReq) GetNroJugador() int64 {
	if x != nil {
		return x.NroJugador
	}
	return 0
}

func (x *E2FinReq) GetNroGroup() int64 {
	if x != nil {
		return x.NroGroup
	}
	return 0
}

type E2FinResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dead int64 `protobuf:"varint,1,opt,name=dead,proto3" json:"dead,omitempty"`
}

func (x *E2FinResp) Reset() {
	*x = E2FinResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LJ_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *E2FinResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*E2FinResp) ProtoMessage() {}

func (x *E2FinResp) ProtoReflect() protoreflect.Message {
	mi := &file_LJ_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use E2FinResp.ProtoReflect.Descriptor instead.
func (*E2FinResp) Descriptor() ([]byte, []int) {
	return file_LJ_proto_rawDescGZIP(), []int{9}
}

func (x *E2FinResp) GetDead() int64 {
	if x != nil {
		return x.Dead
	}
	return 0
}

var File_LJ_proto protoreflect.FileDescriptor

var file_LJ_proto_rawDesc = []byte{
	0x0a, 0x08, 0x4c, 0x4a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x4c, 0x4a, 0x22, 0x25,
	0x0a, 0x07, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x4d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x4d, 0x73, 0x67, 0x22, 0x44, 0x0a, 0x08, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x61, 0x6d, 0x65, 0x4d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x67, 0x61, 0x6d, 0x65, 0x4d, 0x73, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x6e,
	0x72, 0x6f, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x6e, 0x72, 0x6f, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x6f, 0x72, 0x22, 0x73, 0x0a, 0x0b, 0x4e,
	0x75, 0x6d, 0x50, 0x61, 0x73, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x6c,
	0x61, 0x79, 0x4d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x50, 0x6c, 0x61,
	0x79, 0x4d, 0x73, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x6e, 0x72, 0x6f, 0x4a, 0x75, 0x67, 0x61, 0x64,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6e, 0x72, 0x6f, 0x4a, 0x75, 0x67,
	0x61, 0x64, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f,
	0x6e, 0x64, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x72, 0x6f, 0x6e, 0x64, 0x61,
	0x22, 0x2a, 0x0a, 0x0c, 0x4e, 0x75, 0x6d, 0x50, 0x61, 0x73, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x74, 0x65, 0x4d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x73, 0x74, 0x61, 0x74, 0x65, 0x4d, 0x73, 0x67, 0x22, 0x2b, 0x0a, 0x09,
	0x45, 0x32, 0x43, 0x6f, 0x6e, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1e, 0x0a, 0x0a, 0x6e, 0x72, 0x6f,
	0x4a, 0x75, 0x67, 0x61, 0x64, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6e,
	0x72, 0x6f, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x6f, 0x72, 0x22, 0x28, 0x0a, 0x0a, 0x45, 0x32, 0x43,
	0x6f, 0x6e, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x72, 0x6f, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6e, 0x72, 0x6f, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x22, 0x43, 0x0a, 0x09, 0x45, 0x74, 0x61, 0x70, 0x61, 0x32, 0x52, 0x65, 0x71,
	0x12, 0x1e, 0x0a, 0x0a, 0x6e, 0x72, 0x6f, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6e, 0x72, 0x6f, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x6f, 0x72,
	0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x6f, 0x22, 0x28, 0x0a, 0x0a, 0x45, 0x74, 0x61, 0x70,
	0x61, 0x32, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x74, 0x65, 0x4d,
	0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x74, 0x61, 0x74, 0x65, 0x4d,
	0x73, 0x67, 0x22, 0x46, 0x0a, 0x08, 0x45, 0x32, 0x46, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1e,
	0x0a, 0x0a, 0x6e, 0x72, 0x6f, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x6e, 0x72, 0x6f, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x6f, 0x72, 0x12, 0x1a,
	0x0a, 0x08, 0x6e, 0x72, 0x6f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x6e, 0x72, 0x6f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x1f, 0x0a, 0x09, 0x45, 0x32,
	0x46, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x61, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x64, 0x65, 0x61, 0x64, 0x32, 0xee, 0x01, 0x0a, 0x13,
	0x4c, 0x69, 0x64, 0x65, 0x72, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x47, 0x61,
	0x6d, 0x65, 0x12, 0x0b, 0x2e, 0x4c, 0x4a, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x1a,
	0x0c, 0x2e, 0x4c, 0x4a, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2d, 0x0a,
	0x08, 0x45, 0x74, 0x61, 0x70, 0x61, 0x55, 0x6e, 0x6f, 0x12, 0x0f, 0x2e, 0x4c, 0x4a, 0x2e, 0x4e,
	0x75, 0x6d, 0x50, 0x61, 0x73, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x4c, 0x4a, 0x2e,
	0x4e, 0x75, 0x6d, 0x50, 0x61, 0x73, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2b, 0x0a, 0x0a,
	0x45, 0x74, 0x61, 0x70, 0x61, 0x32, 0x43, 0x6f, 0x6e, 0x6e, 0x12, 0x0d, 0x2e, 0x4c, 0x4a, 0x2e,
	0x45, 0x32, 0x43, 0x6f, 0x6e, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x4c, 0x4a, 0x2e, 0x45,
	0x32, 0x43, 0x6f, 0x6e, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x27, 0x0a, 0x06, 0x45, 0x74, 0x61,
	0x70, 0x61, 0x32, 0x12, 0x0d, 0x2e, 0x4c, 0x4a, 0x2e, 0x45, 0x74, 0x61, 0x70, 0x61, 0x32, 0x52,
	0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x4c, 0x4a, 0x2e, 0x45, 0x74, 0x61, 0x70, 0x61, 0x32, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x28, 0x0a, 0x09, 0x45, 0x74, 0x61, 0x70, 0x61, 0x32, 0x46, 0x69, 0x6e, 0x12,
	0x0c, 0x2e, 0x4c, 0x4a, 0x2e, 0x45, 0x32, 0x46, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e,
	0x4c, 0x4a, 0x2e, 0x45, 0x32, 0x46, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x42, 0x0d, 0x5a, 0x0b,
	0x6c, 0x61, 0x62, 0x2f, 0x6a, 0x75, 0x67, 0x61, 0x64, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_LJ_proto_rawDescOnce sync.Once
	file_LJ_proto_rawDescData = file_LJ_proto_rawDesc
)

func file_LJ_proto_rawDescGZIP() []byte {
	file_LJ_proto_rawDescOnce.Do(func() {
		file_LJ_proto_rawDescData = protoimpl.X.CompressGZIP(file_LJ_proto_rawDescData)
	})
	return file_LJ_proto_rawDescData
}

var file_LJ_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_LJ_proto_goTypes = []interface{}{
	(*GameReq)(nil),      // 0: LJ.GameReq
	(*GameResp)(nil),     // 1: LJ.GameResp
	(*NumPasosReq)(nil),  // 2: LJ.NumPasosReq
	(*NumPasosResp)(nil), // 3: LJ.NumPasosResp
	(*E2ConnReq)(nil),    // 4: LJ.E2ConnReq
	(*E2ConnResp)(nil),   // 5: LJ.E2ConnResp
	(*Etapa2Req)(nil),    // 6: LJ.Etapa2Req
	(*Etapa2Resp)(nil),   // 7: LJ.Etapa2Resp
	(*E2FinReq)(nil),     // 8: LJ.E2FinReq
	(*E2FinResp)(nil),    // 9: LJ.E2FinResp
}
var file_LJ_proto_depIdxs = []int32{
	0, // 0: LJ.LiderJugadorService.RequestGame:input_type -> LJ.GameReq
	2, // 1: LJ.LiderJugadorService.EtapaUno:input_type -> LJ.NumPasosReq
	4, // 2: LJ.LiderJugadorService.Etapa2Conn:input_type -> LJ.E2ConnReq
	6, // 3: LJ.LiderJugadorService.Etapa2:input_type -> LJ.Etapa2Req
	8, // 4: LJ.LiderJugadorService.Etapa2Fin:input_type -> LJ.E2FinReq
	1, // 5: LJ.LiderJugadorService.RequestGame:output_type -> LJ.GameResp
	3, // 6: LJ.LiderJugadorService.EtapaUno:output_type -> LJ.NumPasosResp
	5, // 7: LJ.LiderJugadorService.Etapa2Conn:output_type -> LJ.E2ConnResp
	7, // 8: LJ.LiderJugadorService.Etapa2:output_type -> LJ.Etapa2Resp
	9, // 9: LJ.LiderJugadorService.Etapa2Fin:output_type -> LJ.E2FinResp
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_LJ_proto_init() }
func file_LJ_proto_init() {
	if File_LJ_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_LJ_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_LJ_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_LJ_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_LJ_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_LJ_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*E2ConnReq); i {
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
		file_LJ_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*E2ConnResp); i {
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
		file_LJ_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Etapa2Req); i {
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
		file_LJ_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Etapa2Resp); i {
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
		file_LJ_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*E2FinReq); i {
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
		file_LJ_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*E2FinResp); i {
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
			RawDescriptor: file_LJ_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_LJ_proto_goTypes,
		DependencyIndexes: file_LJ_proto_depIdxs,
		MessageInfos:      file_LJ_proto_msgTypes,
	}.Build()
	File_LJ_proto = out.File
	file_LJ_proto_rawDesc = nil
	file_LJ_proto_goTypes = nil
	file_LJ_proto_depIdxs = nil
}
