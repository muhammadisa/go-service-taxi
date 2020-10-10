// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: car.proto

package car_grpc

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

type Car struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID            string               `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	UserID        string               `protobuf:"bytes,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
	BrandID       string               `protobuf:"bytes,3,opt,name=BrandID,proto3" json:"BrandID,omitempty"`
	CarName       string               `protobuf:"bytes,4,opt,name=CarName,proto3" json:"CarName,omitempty"`
	Condition     string               `protobuf:"bytes,5,opt,name=Condition,proto3" json:"Condition,omitempty"`
	Description   string               `protobuf:"bytes,6,opt,name=Description,proto3" json:"Description,omitempty"`
	Specification string               `protobuf:"bytes,7,opt,name=Specification,proto3" json:"Specification,omitempty"`
	Unit          uint64               `protobuf:"varint,8,opt,name=Unit,proto3" json:"Unit,omitempty"`
	Price         uint64               `protobuf:"varint,9,opt,name=Price,proto3" json:"Price,omitempty"`
	UpdatedAt     *timestamp.Timestamp `protobuf:"bytes,10,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	CreatedAt     *timestamp.Timestamp `protobuf:"bytes,11,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
}

func (x *Car) Reset() {
	*x = Car{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Car) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Car) ProtoMessage() {}

func (x *Car) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Car.ProtoReflect.Descriptor instead.
func (*Car) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{0}
}

func (x *Car) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Car) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *Car) GetBrandID() string {
	if x != nil {
		return x.BrandID
	}
	return ""
}

func (x *Car) GetCarName() string {
	if x != nil {
		return x.CarName
	}
	return ""
}

func (x *Car) GetCondition() string {
	if x != nil {
		return x.Condition
	}
	return ""
}

func (x *Car) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Car) GetSpecification() string {
	if x != nil {
		return x.Specification
	}
	return ""
}

func (x *Car) GetUnit() uint64 {
	if x != nil {
		return x.Unit
	}
	return 0
}

func (x *Car) GetPrice() uint64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Car) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Car) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type CreateCar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Car    *Car   `protobuf:"bytes,1,opt,name=car,proto3" json:"car,omitempty"`
	UserID string `protobuf:"bytes,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *CreateCar) Reset() {
	*x = CreateCar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCar) ProtoMessage() {}

func (x *CreateCar) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCar.ProtoReflect.Descriptor instead.
func (*CreateCar) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCar) GetCar() *Car {
	if x != nil {
		return x.Car
	}
	return nil
}

func (x *CreateCar) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type RetrieveCars struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num uint64 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *RetrieveCars) Reset() {
	*x = RetrieveCars{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveCars) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveCars) ProtoMessage() {}

func (x *RetrieveCars) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveCars.ProtoReflect.Descriptor instead.
func (*RetrieveCars) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{2}
}

func (x *RetrieveCars) GetNum() uint64 {
	if x != nil {
		return x.Num
	}
	return 0
}

type RetrieveByID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *RetrieveByID) Reset() {
	*x = RetrieveByID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveByID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveByID) ProtoMessage() {}

func (x *RetrieveByID) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveByID.ProtoReflect.Descriptor instead.
func (*RetrieveByID) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{3}
}

func (x *RetrieveByID) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type RetrieveByUserID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *RetrieveByUserID) Reset() {
	*x = RetrieveByUserID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveByUserID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveByUserID) ProtoMessage() {}

func (x *RetrieveByUserID) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveByUserID.ProtoReflect.Descriptor instead.
func (*RetrieveByUserID) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{4}
}

func (x *RetrieveByUserID) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type DeleteByID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID     string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	UserID string `protobuf:"bytes,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *DeleteByID) Reset() {
	*x = DeleteByID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteByID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteByID) ProtoMessage() {}

func (x *DeleteByID) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteByID.ProtoReflect.Descriptor instead.
func (*DeleteByID) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteByID) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *DeleteByID) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok uint64 `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[6]
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
	return file_car_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteResponse) GetOk() uint64 {
	if x != nil {
		return x.Ok
	}
	return 0
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Car    *Car   `protobuf:"bytes,1,opt,name=car,proto3" json:"car,omitempty"`
	ID     string `protobuf:"bytes,2,opt,name=ID,proto3" json:"ID,omitempty"`
	UserID string `protobuf:"bytes,3,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateRequest) GetCar() *Car {
	if x != nil {
		return x.Car
	}
	return nil
}

func (x *UpdateRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *UpdateRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type ListCars struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cars []*Car `protobuf:"bytes,1,rep,name=cars,proto3" json:"cars,omitempty"`
}

func (x *ListCars) Reset() {
	*x = ListCars{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCars) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCars) ProtoMessage() {}

func (x *ListCars) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCars.ProtoReflect.Descriptor instead.
func (*ListCars) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{8}
}

func (x *ListCars) GetCars() []*Car {
	if x != nil {
		return x.Cars
	}
	return nil
}

var File_car_proto protoreflect.FileDescriptor

var file_car_proto_rawDesc = []byte{
	0x0a, 0x09, 0x63, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x61, 0x72,
	0x5f, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe5, 0x02, 0x0a, 0x03, 0x43, 0x61, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x16,
	0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x49,
	0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x49, 0x44,
	0x12, 0x18, 0x0a, 0x07, 0x43, 0x61, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x43, 0x61, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x53, 0x70,
	0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04,
	0x55, 0x6e, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x44,
	0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x12, 0x1f, 0x0a, 0x03, 0x63,
	0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x61, 0x72, 0x5f, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x43, 0x61, 0x72, 0x52, 0x03, 0x63, 0x61, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x22, 0x20, 0x0a, 0x0c, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65,
	0x43, 0x61, 0x72, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x1e, 0x0a, 0x0c, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65,
	0x76, 0x65, 0x42, 0x79, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x22, 0x0a, 0x10, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65,
	0x76, 0x65, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x34, 0x0a, 0x0a, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x22, 0x20, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x6f, 0x6b, 0x22, 0x58, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x03, 0x63, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x63, 0x61, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x61, 0x72, 0x52,
	0x03, 0x63, 0x61, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x2d, 0x0a, 0x08,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72, 0x73, 0x12, 0x21, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x61, 0x72, 0x5f, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x43, 0x61, 0x72, 0x52, 0x04, 0x63, 0x61, 0x72, 0x73, 0x32, 0xd4, 0x02, 0x0a, 0x0a,
	0x43, 0x61, 0x72, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x43, 0x61, 0x72, 0x73, 0x12, 0x16, 0x2e, 0x63, 0x61, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x43, 0x61, 0x72, 0x73, 0x1a, 0x12, 0x2e,
	0x63, 0x61, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72,
	0x73, 0x12, 0x33, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x42, 0x79, 0x49, 0x44, 0x12,
	0x16, 0x2e, 0x63, 0x61, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69,
	0x65, 0x76, 0x65, 0x42, 0x79, 0x49, 0x44, 0x1a, 0x0d, 0x2e, 0x63, 0x61, 0x72, 0x5f, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x43, 0x61, 0x72, 0x12, 0x41, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72,
	0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x2e, 0x63, 0x61, 0x72, 0x5f,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x42, 0x79, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x1a, 0x12, 0x2e, 0x63, 0x61, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72, 0x73, 0x12, 0x2b, 0x0a, 0x05, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x12, 0x13, 0x2e, 0x63, 0x61, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x1a, 0x0d, 0x2e, 0x63, 0x61, 0x72, 0x5f, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x43, 0x61, 0x72, 0x12, 0x30, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x17, 0x2e, 0x63, 0x61, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x63, 0x61, 0x72, 0x5f,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x61, 0x72, 0x12, 0x38, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x14, 0x2e, 0x63, 0x61, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x44, 0x1a, 0x18, 0x2e, 0x63, 0x61, 0x72, 0x5f, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x13, 0x5a, 0x11, 0x63, 0x61, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x3b, 0x63,
	0x61, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_car_proto_rawDescOnce sync.Once
	file_car_proto_rawDescData = file_car_proto_rawDesc
)

func file_car_proto_rawDescGZIP() []byte {
	file_car_proto_rawDescOnce.Do(func() {
		file_car_proto_rawDescData = protoimpl.X.CompressGZIP(file_car_proto_rawDescData)
	})
	return file_car_proto_rawDescData
}

var file_car_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_car_proto_goTypes = []interface{}{
	(*Car)(nil),                 // 0: car_grpc.Car
	(*CreateCar)(nil),           // 1: car_grpc.CreateCar
	(*RetrieveCars)(nil),        // 2: car_grpc.RetrieveCars
	(*RetrieveByID)(nil),        // 3: car_grpc.RetrieveByID
	(*RetrieveByUserID)(nil),    // 4: car_grpc.RetrieveByUserID
	(*DeleteByID)(nil),          // 5: car_grpc.DeleteByID
	(*DeleteResponse)(nil),      // 6: car_grpc.DeleteResponse
	(*UpdateRequest)(nil),       // 7: car_grpc.UpdateRequest
	(*ListCars)(nil),            // 8: car_grpc.ListCars
	(*timestamp.Timestamp)(nil), // 9: google.protobuf.Timestamp
}
var file_car_proto_depIdxs = []int32{
	9,  // 0: car_grpc.Car.UpdatedAt:type_name -> google.protobuf.Timestamp
	9,  // 1: car_grpc.Car.CreatedAt:type_name -> google.protobuf.Timestamp
	0,  // 2: car_grpc.CreateCar.car:type_name -> car_grpc.Car
	0,  // 3: car_grpc.UpdateRequest.car:type_name -> car_grpc.Car
	0,  // 4: car_grpc.ListCars.cars:type_name -> car_grpc.Car
	2,  // 5: car_grpc.CarHandler.GetCars:input_type -> car_grpc.RetrieveCars
	3,  // 6: car_grpc.CarHandler.GetCarByID:input_type -> car_grpc.RetrieveByID
	4,  // 7: car_grpc.CarHandler.GetCarsByUserID:input_type -> car_grpc.RetrieveByUserID
	1,  // 8: car_grpc.CarHandler.Store:input_type -> car_grpc.CreateCar
	7,  // 9: car_grpc.CarHandler.Update:input_type -> car_grpc.UpdateRequest
	5,  // 10: car_grpc.CarHandler.Delete:input_type -> car_grpc.DeleteByID
	8,  // 11: car_grpc.CarHandler.GetCars:output_type -> car_grpc.ListCars
	0,  // 12: car_grpc.CarHandler.GetCarByID:output_type -> car_grpc.Car
	8,  // 13: car_grpc.CarHandler.GetCarsByUserID:output_type -> car_grpc.ListCars
	0,  // 14: car_grpc.CarHandler.Store:output_type -> car_grpc.Car
	0,  // 15: car_grpc.CarHandler.Update:output_type -> car_grpc.Car
	6,  // 16: car_grpc.CarHandler.Delete:output_type -> car_grpc.DeleteResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_car_proto_init() }
func file_car_proto_init() {
	if File_car_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_car_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Car); i {
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
		file_car_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCar); i {
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
		file_car_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveCars); i {
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
		file_car_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveByID); i {
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
		file_car_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveByUserID); i {
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
		file_car_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteByID); i {
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
		file_car_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_car_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_car_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCars); i {
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
			RawDescriptor: file_car_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_car_proto_goTypes,
		DependencyIndexes: file_car_proto_depIdxs,
		MessageInfos:      file_car_proto_msgTypes,
	}.Build()
	File_car_proto = out.File
	file_car_proto_rawDesc = nil
	file_car_proto_goTypes = nil
	file_car_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CarHandlerClient is the client API for CarHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CarHandlerClient interface {
	GetCars(ctx context.Context, in *RetrieveCars, opts ...grpc.CallOption) (*ListCars, error)
	GetCarByID(ctx context.Context, in *RetrieveByID, opts ...grpc.CallOption) (*Car, error)
	GetCarsByUserID(ctx context.Context, in *RetrieveByUserID, opts ...grpc.CallOption) (*ListCars, error)
	Store(ctx context.Context, in *CreateCar, opts ...grpc.CallOption) (*Car, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Car, error)
	Delete(ctx context.Context, in *DeleteByID, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type carHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewCarHandlerClient(cc grpc.ClientConnInterface) CarHandlerClient {
	return &carHandlerClient{cc}
}

func (c *carHandlerClient) GetCars(ctx context.Context, in *RetrieveCars, opts ...grpc.CallOption) (*ListCars, error) {
	out := new(ListCars)
	err := c.cc.Invoke(ctx, "/car_grpc.CarHandler/GetCars", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carHandlerClient) GetCarByID(ctx context.Context, in *RetrieveByID, opts ...grpc.CallOption) (*Car, error) {
	out := new(Car)
	err := c.cc.Invoke(ctx, "/car_grpc.CarHandler/GetCarByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carHandlerClient) GetCarsByUserID(ctx context.Context, in *RetrieveByUserID, opts ...grpc.CallOption) (*ListCars, error) {
	out := new(ListCars)
	err := c.cc.Invoke(ctx, "/car_grpc.CarHandler/GetCarsByUserID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carHandlerClient) Store(ctx context.Context, in *CreateCar, opts ...grpc.CallOption) (*Car, error) {
	out := new(Car)
	err := c.cc.Invoke(ctx, "/car_grpc.CarHandler/Store", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carHandlerClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Car, error) {
	out := new(Car)
	err := c.cc.Invoke(ctx, "/car_grpc.CarHandler/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carHandlerClient) Delete(ctx context.Context, in *DeleteByID, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/car_grpc.CarHandler/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CarHandlerServer is the server API for CarHandler service.
type CarHandlerServer interface {
	GetCars(context.Context, *RetrieveCars) (*ListCars, error)
	GetCarByID(context.Context, *RetrieveByID) (*Car, error)
	GetCarsByUserID(context.Context, *RetrieveByUserID) (*ListCars, error)
	Store(context.Context, *CreateCar) (*Car, error)
	Update(context.Context, *UpdateRequest) (*Car, error)
	Delete(context.Context, *DeleteByID) (*DeleteResponse, error)
}

// UnimplementedCarHandlerServer can be embedded to have forward compatible implementations.
type UnimplementedCarHandlerServer struct {
}

func (*UnimplementedCarHandlerServer) GetCars(context.Context, *RetrieveCars) (*ListCars, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCars not implemented")
}
func (*UnimplementedCarHandlerServer) GetCarByID(context.Context, *RetrieveByID) (*Car, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCarByID not implemented")
}
func (*UnimplementedCarHandlerServer) GetCarsByUserID(context.Context, *RetrieveByUserID) (*ListCars, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCarsByUserID not implemented")
}
func (*UnimplementedCarHandlerServer) Store(context.Context, *CreateCar) (*Car, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Store not implemented")
}
func (*UnimplementedCarHandlerServer) Update(context.Context, *UpdateRequest) (*Car, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedCarHandlerServer) Delete(context.Context, *DeleteByID) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterCarHandlerServer(s *grpc.Server, srv CarHandlerServer) {
	s.RegisterService(&_CarHandler_serviceDesc, srv)
}

func _CarHandler_GetCars_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveCars)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarHandlerServer).GetCars(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/car_grpc.CarHandler/GetCars",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarHandlerServer).GetCars(ctx, req.(*RetrieveCars))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarHandler_GetCarByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarHandlerServer).GetCarByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/car_grpc.CarHandler/GetCarByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarHandlerServer).GetCarByID(ctx, req.(*RetrieveByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarHandler_GetCarsByUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveByUserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarHandlerServer).GetCarsByUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/car_grpc.CarHandler/GetCarsByUserID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarHandlerServer).GetCarsByUserID(ctx, req.(*RetrieveByUserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarHandler_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCar)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarHandlerServer).Store(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/car_grpc.CarHandler/Store",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarHandlerServer).Store(ctx, req.(*CreateCar))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarHandler_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarHandlerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/car_grpc.CarHandler/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarHandlerServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarHandler_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarHandlerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/car_grpc.CarHandler/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarHandlerServer).Delete(ctx, req.(*DeleteByID))
	}
	return interceptor(ctx, in, info, handler)
}

var _CarHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "car_grpc.CarHandler",
	HandlerType: (*CarHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCars",
			Handler:    _CarHandler_GetCars_Handler,
		},
		{
			MethodName: "GetCarByID",
			Handler:    _CarHandler_GetCarByID_Handler,
		},
		{
			MethodName: "GetCarsByUserID",
			Handler:    _CarHandler_GetCarsByUserID_Handler,
		},
		{
			MethodName: "Store",
			Handler:    _CarHandler_Store_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CarHandler_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CarHandler_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "car.proto",
}
