// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: foobar.proto

package foobar_grpc

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Foobar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID            string               `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	FoobarContent string               `protobuf:"bytes,2,opt,name=FoobarContent,proto3" json:"FoobarContent,omitempty"`
	UpdatedAt     *timestamp.Timestamp `protobuf:"bytes,3,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	CreatedAt     *timestamp.Timestamp `protobuf:"bytes,4,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
}

func (x *Foobar) Reset() {
	*x = Foobar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foobar_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Foobar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Foobar) ProtoMessage() {}

func (x *Foobar) ProtoReflect() protoreflect.Message {
	mi := &file_foobar_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Foobar.ProtoReflect.Descriptor instead.
func (*Foobar) Descriptor() ([]byte, []int) {
	return file_foobar_proto_rawDescGZIP(), []int{0}
}

func (x *Foobar) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Foobar) GetFoobarContent() string {
	if x != nil {
		return x.FoobarContent
	}
	return ""
}

func (x *Foobar) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Foobar) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type ListFoobar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Foobars []*Foobar `protobuf:"bytes,1,rep,name=Foobars,proto3" json:"Foobars,omitempty"`
}

func (x *ListFoobar) Reset() {
	*x = ListFoobar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foobar_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFoobar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFoobar) ProtoMessage() {}

func (x *ListFoobar) ProtoReflect() protoreflect.Message {
	mi := &file_foobar_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFoobar.ProtoReflect.Descriptor instead.
func (*ListFoobar) Descriptor() ([]byte, []int) {
	return file_foobar_proto_rawDescGZIP(), []int{1}
}

func (x *ListFoobar) GetFoobars() []*Foobar {
	if x != nil {
		return x.Foobars
	}
	return nil
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Code   uint64 `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foobar_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_foobar_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_foobar_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *DeleteResponse) GetCode() uint64 {
	if x != nil {
		return x.Code
	}
	return 0
}

type FetchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num uint64 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *FetchRequest) Reset() {
	*x = FetchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foobar_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchRequest) ProtoMessage() {}

func (x *FetchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_foobar_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchRequest.ProtoReflect.Descriptor instead.
func (*FetchRequest) Descriptor() ([]byte, []int) {
	return file_foobar_proto_rawDescGZIP(), []int{3}
}

func (x *FetchRequest) GetNum() uint64 {
	if x != nil {
		return x.Num
	}
	return 0
}

type SingleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *SingleRequest) Reset() {
	*x = SingleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foobar_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SingleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SingleRequest) ProtoMessage() {}

func (x *SingleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_foobar_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SingleRequest.ProtoReflect.Descriptor instead.
func (*SingleRequest) Descriptor() ([]byte, []int) {
	return file_foobar_proto_rawDescGZIP(), []int{4}
}

func (x *SingleRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type ErrorMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ErrorMessage) Reset() {
	*x = ErrorMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foobar_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorMessage) ProtoMessage() {}

func (x *ErrorMessage) ProtoReflect() protoreflect.Message {
	mi := &file_foobar_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorMessage.ProtoReflect.Descriptor instead.
func (*ErrorMessage) Descriptor() ([]byte, []int) {
	return file_foobar_proto_rawDescGZIP(), []int{5}
}

func (x *ErrorMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_foobar_proto protoreflect.FileDescriptor

var file_foobar_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x66, 0x6f, 0x6f, 0x62, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x66, 0x6f, 0x6f, 0x62, 0x61, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x01, 0x0a,
	0x06, 0x46, 0x6f, 0x6f, 0x62, 0x61, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x46, 0x6f, 0x6f, 0x62, 0x61,
	0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x46, 0x6f, 0x6f, 0x62, 0x61, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x38, 0x0a,
	0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0x3b, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6f, 0x62, 0x61, 0x72, 0x12,
	0x2d, 0x0a, 0x07, 0x46, 0x6f, 0x6f, 0x62, 0x61, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x66, 0x6f, 0x6f, 0x62, 0x61, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x46,
	0x6f, 0x6f, 0x62, 0x61, 0x72, 0x52, 0x07, 0x46, 0x6f, 0x6f, 0x62, 0x61, 0x72, 0x73, 0x22, 0x3c,
	0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x20, 0x0a, 0x0c,
	0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x1f,
	0x0a, 0x0d, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22,
	0x28, 0x0a, 0x0c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xc2, 0x02, 0x0a, 0x0d, 0x46, 0x6f,
	0x6f, 0x62, 0x61, 0x72, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x46, 0x6f, 0x6f, 0x62, 0x61, 0x72, 0x12, 0x1a, 0x2e, 0x66, 0x6f, 0x6f, 0x62, 0x61,
	0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x66, 0x6f, 0x6f, 0x62, 0x61, 0x72, 0x5f, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x46, 0x6f, 0x6f, 0x62, 0x61, 0x72, 0x12, 0x43, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6f, 0x62, 0x61, 0x72, 0x12, 0x19, 0x2e, 0x66, 0x6f, 0x6f,
	0x62, 0x61, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x66, 0x6f, 0x6f, 0x62, 0x61, 0x72, 0x5f, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6f, 0x62, 0x61, 0x72, 0x12, 0x38,
	0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x62, 0x61, 0x72, 0x12, 0x13,
	0x2e, 0x66, 0x6f, 0x6f, 0x62, 0x61, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x6f, 0x6f,
	0x62, 0x61, 0x72, 0x1a, 0x13, 0x2e, 0x66, 0x6f, 0x6f, 0x62, 0x61, 0x72, 0x5f, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x46, 0x6f, 0x6f, 0x62, 0x61, 0x72, 0x12, 0x41, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x1a, 0x2e, 0x66, 0x6f, 0x6f, 0x62, 0x61, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x66, 0x6f, 0x6f, 0x62, 0x61, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x05, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x12, 0x13, 0x2e, 0x66, 0x6f, 0x6f, 0x62, 0x61, 0x72, 0x5f, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x46, 0x6f, 0x6f, 0x62, 0x61, 0x72, 0x1a, 0x13, 0x2e, 0x66, 0x6f, 0x6f, 0x62,
	0x61, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x6f, 0x6f, 0x62, 0x61, 0x72, 0x42, 0x19,
	0x5a, 0x17, 0x66, 0x6f, 0x6f, 0x62, 0x61, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x3b, 0x66, 0x6f,
	0x6f, 0x62, 0x61, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_foobar_proto_rawDescOnce sync.Once
	file_foobar_proto_rawDescData = file_foobar_proto_rawDesc
)

func file_foobar_proto_rawDescGZIP() []byte {
	file_foobar_proto_rawDescOnce.Do(func() {
		file_foobar_proto_rawDescData = protoimpl.X.CompressGZIP(file_foobar_proto_rawDescData)
	})
	return file_foobar_proto_rawDescData
}

var file_foobar_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_foobar_proto_goTypes = []interface{}{
	(*Foobar)(nil),              // 0: foobar_grpc.Foobar
	(*ListFoobar)(nil),          // 1: foobar_grpc.ListFoobar
	(*DeleteResponse)(nil),      // 2: foobar_grpc.DeleteResponse
	(*FetchRequest)(nil),        // 3: foobar_grpc.FetchRequest
	(*SingleRequest)(nil),       // 4: foobar_grpc.SingleRequest
	(*ErrorMessage)(nil),        // 5: foobar_grpc.ErrorMessage
	(*timestamp.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_foobar_proto_depIdxs = []int32{
	6, // 0: foobar_grpc.Foobar.UpdatedAt:type_name -> google.protobuf.Timestamp
	6, // 1: foobar_grpc.Foobar.CreatedAt:type_name -> google.protobuf.Timestamp
	0, // 2: foobar_grpc.ListFoobar.Foobars:type_name -> foobar_grpc.Foobar
	4, // 3: foobar_grpc.FoobarHandler.GetFoobar:input_type -> foobar_grpc.SingleRequest
	3, // 4: foobar_grpc.FoobarHandler.GetListFoobar:input_type -> foobar_grpc.FetchRequest
	0, // 5: foobar_grpc.FoobarHandler.UpdateFoobar:input_type -> foobar_grpc.Foobar
	4, // 6: foobar_grpc.FoobarHandler.Delete:input_type -> foobar_grpc.SingleRequest
	0, // 7: foobar_grpc.FoobarHandler.Store:input_type -> foobar_grpc.Foobar
	0, // 8: foobar_grpc.FoobarHandler.GetFoobar:output_type -> foobar_grpc.Foobar
	1, // 9: foobar_grpc.FoobarHandler.GetListFoobar:output_type -> foobar_grpc.ListFoobar
	0, // 10: foobar_grpc.FoobarHandler.UpdateFoobar:output_type -> foobar_grpc.Foobar
	2, // 11: foobar_grpc.FoobarHandler.Delete:output_type -> foobar_grpc.DeleteResponse
	0, // 12: foobar_grpc.FoobarHandler.Store:output_type -> foobar_grpc.Foobar
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_foobar_proto_init() }
func file_foobar_proto_init() {
	if File_foobar_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_foobar_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Foobar); i {
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
		file_foobar_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFoobar); i {
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
		file_foobar_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
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
		file_foobar_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchRequest); i {
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
		file_foobar_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SingleRequest); i {
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
		file_foobar_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorMessage); i {
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
			RawDescriptor: file_foobar_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_foobar_proto_goTypes,
		DependencyIndexes: file_foobar_proto_depIdxs,
		MessageInfos:      file_foobar_proto_msgTypes,
	}.Build()
	File_foobar_proto = out.File
	file_foobar_proto_rawDesc = nil
	file_foobar_proto_goTypes = nil
	file_foobar_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FoobarHandlerClient is the client API for FoobarHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FoobarHandlerClient interface {
	GetFoobar(ctx context.Context, in *SingleRequest, opts ...grpc.CallOption) (*Foobar, error)
	GetListFoobar(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*ListFoobar, error)
	UpdateFoobar(ctx context.Context, in *Foobar, opts ...grpc.CallOption) (*Foobar, error)
	Delete(ctx context.Context, in *SingleRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	Store(ctx context.Context, in *Foobar, opts ...grpc.CallOption) (*Foobar, error)
}

type foobarHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewFoobarHandlerClient(cc grpc.ClientConnInterface) FoobarHandlerClient {
	return &foobarHandlerClient{cc}
}

func (c *foobarHandlerClient) GetFoobar(ctx context.Context, in *SingleRequest, opts ...grpc.CallOption) (*Foobar, error) {
	out := new(Foobar)
	err := c.cc.Invoke(ctx, "/foobar_grpc.FoobarHandler/GetFoobar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foobarHandlerClient) GetListFoobar(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*ListFoobar, error) {
	out := new(ListFoobar)
	err := c.cc.Invoke(ctx, "/foobar_grpc.FoobarHandler/GetListFoobar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foobarHandlerClient) UpdateFoobar(ctx context.Context, in *Foobar, opts ...grpc.CallOption) (*Foobar, error) {
	out := new(Foobar)
	err := c.cc.Invoke(ctx, "/foobar_grpc.FoobarHandler/UpdateFoobar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foobarHandlerClient) Delete(ctx context.Context, in *SingleRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/foobar_grpc.FoobarHandler/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foobarHandlerClient) Store(ctx context.Context, in *Foobar, opts ...grpc.CallOption) (*Foobar, error) {
	out := new(Foobar)
	err := c.cc.Invoke(ctx, "/foobar_grpc.FoobarHandler/Store", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FoobarHandlerServer is the server API for FoobarHandler service.
type FoobarHandlerServer interface {
	GetFoobar(context.Context, *SingleRequest) (*Foobar, error)
	GetListFoobar(context.Context, *FetchRequest) (*ListFoobar, error)
	UpdateFoobar(context.Context, *Foobar) (*Foobar, error)
	Delete(context.Context, *SingleRequest) (*DeleteResponse, error)
	Store(context.Context, *Foobar) (*Foobar, error)
}

// UnimplementedFoobarHandlerServer can be embedded to have forward compatible implementations.
type UnimplementedFoobarHandlerServer struct {
}

func (*UnimplementedFoobarHandlerServer) GetFoobar(context.Context, *SingleRequest) (*Foobar, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFoobar not implemented")
}
func (*UnimplementedFoobarHandlerServer) GetListFoobar(context.Context, *FetchRequest) (*ListFoobar, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListFoobar not implemented")
}
func (*UnimplementedFoobarHandlerServer) UpdateFoobar(context.Context, *Foobar) (*Foobar, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFoobar not implemented")
}
func (*UnimplementedFoobarHandlerServer) Delete(context.Context, *SingleRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedFoobarHandlerServer) Store(context.Context, *Foobar) (*Foobar, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Store not implemented")
}

func RegisterFoobarHandlerServer(s *grpc.Server, srv FoobarHandlerServer) {
	s.RegisterService(&_FoobarHandler_serviceDesc, srv)
}

func _FoobarHandler_GetFoobar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoobarHandlerServer).GetFoobar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/foobar_grpc.FoobarHandler/GetFoobar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoobarHandlerServer).GetFoobar(ctx, req.(*SingleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoobarHandler_GetListFoobar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoobarHandlerServer).GetListFoobar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/foobar_grpc.FoobarHandler/GetListFoobar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoobarHandlerServer).GetListFoobar(ctx, req.(*FetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoobarHandler_UpdateFoobar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Foobar)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoobarHandlerServer).UpdateFoobar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/foobar_grpc.FoobarHandler/UpdateFoobar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoobarHandlerServer).UpdateFoobar(ctx, req.(*Foobar))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoobarHandler_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoobarHandlerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/foobar_grpc.FoobarHandler/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoobarHandlerServer).Delete(ctx, req.(*SingleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoobarHandler_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Foobar)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoobarHandlerServer).Store(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/foobar_grpc.FoobarHandler/Store",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoobarHandlerServer).Store(ctx, req.(*Foobar))
	}
	return interceptor(ctx, in, info, handler)
}

var _FoobarHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "foobar_grpc.FoobarHandler",
	HandlerType: (*FoobarHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFoobar",
			Handler:    _FoobarHandler_GetFoobar_Handler,
		},
		{
			MethodName: "GetListFoobar",
			Handler:    _FoobarHandler_GetListFoobar_Handler,
		},
		{
			MethodName: "UpdateFoobar",
			Handler:    _FoobarHandler_UpdateFoobar_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _FoobarHandler_Delete_Handler,
		},
		{
			MethodName: "Store",
			Handler:    _FoobarHandler_Store_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "foobar.proto",
}