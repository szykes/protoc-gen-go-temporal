// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: test/simple/v1/simple.proto

// buf:lint:ignore PACKAGE_DIRECTORY_MATCH

package v1

import (
	_ "github.com/cludden/protoc-gen-go-temporal/gen/temporal/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SomeWorkflow1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestVal string `protobuf:"bytes,1,opt,name=request_val,json=requestVal,proto3" json:"request_val,omitempty"`
	Id         string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SomeWorkflow1Request) Reset() {
	*x = SomeWorkflow1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_simple_v1_simple_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SomeWorkflow1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SomeWorkflow1Request) ProtoMessage() {}

func (x *SomeWorkflow1Request) ProtoReflect() protoreflect.Message {
	mi := &file_test_simple_v1_simple_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SomeWorkflow1Request.ProtoReflect.Descriptor instead.
func (*SomeWorkflow1Request) Descriptor() ([]byte, []int) {
	return file_test_simple_v1_simple_proto_rawDescGZIP(), []int{0}
}

func (x *SomeWorkflow1Request) GetRequestVal() string {
	if x != nil {
		return x.RequestVal
	}
	return ""
}

func (x *SomeWorkflow1Request) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SomeWorkflow1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResponseVal string `protobuf:"bytes,1,opt,name=response_val,json=responseVal,proto3" json:"response_val,omitempty"`
}

func (x *SomeWorkflow1Response) Reset() {
	*x = SomeWorkflow1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_simple_v1_simple_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SomeWorkflow1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SomeWorkflow1Response) ProtoMessage() {}

func (x *SomeWorkflow1Response) ProtoReflect() protoreflect.Message {
	mi := &file_test_simple_v1_simple_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SomeWorkflow1Response.ProtoReflect.Descriptor instead.
func (*SomeWorkflow1Response) Descriptor() ([]byte, []int) {
	return file_test_simple_v1_simple_proto_rawDescGZIP(), []int{1}
}

func (x *SomeWorkflow1Response) GetResponseVal() string {
	if x != nil {
		return x.ResponseVal
	}
	return ""
}

type SomeWorkflow3Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RequestVal string `protobuf:"bytes,2,opt,name=request_val,json=requestVal,proto3" json:"request_val,omitempty"`
}

func (x *SomeWorkflow3Request) Reset() {
	*x = SomeWorkflow3Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_simple_v1_simple_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SomeWorkflow3Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SomeWorkflow3Request) ProtoMessage() {}

func (x *SomeWorkflow3Request) ProtoReflect() protoreflect.Message {
	mi := &file_test_simple_v1_simple_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SomeWorkflow3Request.ProtoReflect.Descriptor instead.
func (*SomeWorkflow3Request) Descriptor() ([]byte, []int) {
	return file_test_simple_v1_simple_proto_rawDescGZIP(), []int{2}
}

func (x *SomeWorkflow3Request) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SomeWorkflow3Request) GetRequestVal() string {
	if x != nil {
		return x.RequestVal
	}
	return ""
}

type SomeActivity2Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestVal string `protobuf:"bytes,1,opt,name=request_val,json=requestVal,proto3" json:"request_val,omitempty"`
}

func (x *SomeActivity2Request) Reset() {
	*x = SomeActivity2Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_simple_v1_simple_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SomeActivity2Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SomeActivity2Request) ProtoMessage() {}

func (x *SomeActivity2Request) ProtoReflect() protoreflect.Message {
	mi := &file_test_simple_v1_simple_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SomeActivity2Request.ProtoReflect.Descriptor instead.
func (*SomeActivity2Request) Descriptor() ([]byte, []int) {
	return file_test_simple_v1_simple_proto_rawDescGZIP(), []int{3}
}

func (x *SomeActivity2Request) GetRequestVal() string {
	if x != nil {
		return x.RequestVal
	}
	return ""
}

type SomeActivity3Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestVal string `protobuf:"bytes,1,opt,name=request_val,json=requestVal,proto3" json:"request_val,omitempty"`
}

func (x *SomeActivity3Request) Reset() {
	*x = SomeActivity3Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_simple_v1_simple_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SomeActivity3Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SomeActivity3Request) ProtoMessage() {}

func (x *SomeActivity3Request) ProtoReflect() protoreflect.Message {
	mi := &file_test_simple_v1_simple_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SomeActivity3Request.ProtoReflect.Descriptor instead.
func (*SomeActivity3Request) Descriptor() ([]byte, []int) {
	return file_test_simple_v1_simple_proto_rawDescGZIP(), []int{4}
}

func (x *SomeActivity3Request) GetRequestVal() string {
	if x != nil {
		return x.RequestVal
	}
	return ""
}

type SomeActivity3Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResponseVal string `protobuf:"bytes,1,opt,name=response_val,json=responseVal,proto3" json:"response_val,omitempty"`
}

func (x *SomeActivity3Response) Reset() {
	*x = SomeActivity3Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_simple_v1_simple_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SomeActivity3Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SomeActivity3Response) ProtoMessage() {}

func (x *SomeActivity3Response) ProtoReflect() protoreflect.Message {
	mi := &file_test_simple_v1_simple_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SomeActivity3Response.ProtoReflect.Descriptor instead.
func (*SomeActivity3Response) Descriptor() ([]byte, []int) {
	return file_test_simple_v1_simple_proto_rawDescGZIP(), []int{5}
}

func (x *SomeActivity3Response) GetResponseVal() string {
	if x != nil {
		return x.ResponseVal
	}
	return ""
}

type SomeQuery1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResponseVal string `protobuf:"bytes,1,opt,name=response_val,json=responseVal,proto3" json:"response_val,omitempty"`
}

func (x *SomeQuery1Response) Reset() {
	*x = SomeQuery1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_simple_v1_simple_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SomeQuery1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SomeQuery1Response) ProtoMessage() {}

func (x *SomeQuery1Response) ProtoReflect() protoreflect.Message {
	mi := &file_test_simple_v1_simple_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SomeQuery1Response.ProtoReflect.Descriptor instead.
func (*SomeQuery1Response) Descriptor() ([]byte, []int) {
	return file_test_simple_v1_simple_proto_rawDescGZIP(), []int{6}
}

func (x *SomeQuery1Response) GetResponseVal() string {
	if x != nil {
		return x.ResponseVal
	}
	return ""
}

type SomeQuery2Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestVal string `protobuf:"bytes,1,opt,name=request_val,json=requestVal,proto3" json:"request_val,omitempty"`
}

func (x *SomeQuery2Request) Reset() {
	*x = SomeQuery2Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_simple_v1_simple_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SomeQuery2Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SomeQuery2Request) ProtoMessage() {}

func (x *SomeQuery2Request) ProtoReflect() protoreflect.Message {
	mi := &file_test_simple_v1_simple_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SomeQuery2Request.ProtoReflect.Descriptor instead.
func (*SomeQuery2Request) Descriptor() ([]byte, []int) {
	return file_test_simple_v1_simple_proto_rawDescGZIP(), []int{7}
}

func (x *SomeQuery2Request) GetRequestVal() string {
	if x != nil {
		return x.RequestVal
	}
	return ""
}

type SomeQuery2Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResponseVal string `protobuf:"bytes,1,opt,name=response_val,json=responseVal,proto3" json:"response_val,omitempty"`
}

func (x *SomeQuery2Response) Reset() {
	*x = SomeQuery2Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_simple_v1_simple_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SomeQuery2Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SomeQuery2Response) ProtoMessage() {}

func (x *SomeQuery2Response) ProtoReflect() protoreflect.Message {
	mi := &file_test_simple_v1_simple_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SomeQuery2Response.ProtoReflect.Descriptor instead.
func (*SomeQuery2Response) Descriptor() ([]byte, []int) {
	return file_test_simple_v1_simple_proto_rawDescGZIP(), []int{8}
}

func (x *SomeQuery2Response) GetResponseVal() string {
	if x != nil {
		return x.ResponseVal
	}
	return ""
}

type SomeSignal2Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestVal string `protobuf:"bytes,1,opt,name=request_val,json=requestVal,proto3" json:"request_val,omitempty"`
}

func (x *SomeSignal2Request) Reset() {
	*x = SomeSignal2Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_simple_v1_simple_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SomeSignal2Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SomeSignal2Request) ProtoMessage() {}

func (x *SomeSignal2Request) ProtoReflect() protoreflect.Message {
	mi := &file_test_simple_v1_simple_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SomeSignal2Request.ProtoReflect.Descriptor instead.
func (*SomeSignal2Request) Descriptor() ([]byte, []int) {
	return file_test_simple_v1_simple_proto_rawDescGZIP(), []int{9}
}

func (x *SomeSignal2Request) GetRequestVal() string {
	if x != nil {
		return x.RequestVal
	}
	return ""
}

type SomeUpdate1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestVal string `protobuf:"bytes,1,opt,name=request_val,json=requestVal,proto3" json:"request_val,omitempty"`
}

func (x *SomeUpdate1Request) Reset() {
	*x = SomeUpdate1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_simple_v1_simple_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SomeUpdate1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SomeUpdate1Request) ProtoMessage() {}

func (x *SomeUpdate1Request) ProtoReflect() protoreflect.Message {
	mi := &file_test_simple_v1_simple_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SomeUpdate1Request.ProtoReflect.Descriptor instead.
func (*SomeUpdate1Request) Descriptor() ([]byte, []int) {
	return file_test_simple_v1_simple_proto_rawDescGZIP(), []int{10}
}

func (x *SomeUpdate1Request) GetRequestVal() string {
	if x != nil {
		return x.RequestVal
	}
	return ""
}

type SomeUpdate1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResponseVal string `protobuf:"bytes,1,opt,name=response_val,json=responseVal,proto3" json:"response_val,omitempty"`
}

func (x *SomeUpdate1Response) Reset() {
	*x = SomeUpdate1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_simple_v1_simple_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SomeUpdate1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SomeUpdate1Response) ProtoMessage() {}

func (x *SomeUpdate1Response) ProtoReflect() protoreflect.Message {
	mi := &file_test_simple_v1_simple_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SomeUpdate1Response.ProtoReflect.Descriptor instead.
func (*SomeUpdate1Response) Descriptor() ([]byte, []int) {
	return file_test_simple_v1_simple_proto_rawDescGZIP(), []int{11}
}

func (x *SomeUpdate1Response) GetResponseVal() string {
	if x != nil {
		return x.ResponseVal
	}
	return ""
}

var File_test_simple_v1_simple_proto protoreflect.FileDescriptor

var file_test_simple_v1_simple_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6d,
	0x79, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x74, 0x65,
	0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72,
	0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x47, 0x0a, 0x14, 0x53, 0x6f, 0x6d, 0x65,
	0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x61,
	0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x3a, 0x0a, 0x15, 0x53, 0x6f, 0x6d, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x61, 0x6c, 0x22, 0x47, 0x0a,
	0x14, 0x53, 0x6f, 0x6d, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x33, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x5f, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x22, 0x37, 0x0a, 0x14, 0x53, 0x6f, 0x6d, 0x65, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x32, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x22,
	0x37, 0x0a, 0x14, 0x53, 0x6f, 0x6d, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x33,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x22, 0x3a, 0x0a, 0x15, 0x53, 0x6f, 0x6d, 0x65,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x33, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x76, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x56, 0x61, 0x6c, 0x22, 0x37, 0x0a, 0x12, 0x53, 0x6f, 0x6d, 0x65, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x61, 0x6c, 0x22, 0x34, 0x0a,
	0x11, 0x53, 0x6f, 0x6d, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x32, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x76, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x56, 0x61, 0x6c, 0x22, 0x37, 0x0a, 0x12, 0x53, 0x6f, 0x6d, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x32, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x61, 0x6c, 0x22, 0x35, 0x0a, 0x12,
	0x53, 0x6f, 0x6d, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x32, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x76, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x56, 0x61, 0x6c, 0x22, 0x35, 0x0a, 0x12, 0x53, 0x6f, 0x6d, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x22, 0x38, 0x0a, 0x13, 0x53, 0x6f,
	0x6d, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x76, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x56, 0x61, 0x6c, 0x32, 0x8b, 0x0a, 0x0a, 0x06, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x12,
	0xca, 0x01, 0x0a, 0x0d, 0x53, 0x6f, 0x6d, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x31, 0x12, 0x26, 0x2e, 0x6d, 0x79, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x73, 0x69,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x53, 0x6f, 0x6d, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6d, 0x79, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x53, 0x6f, 0x6d,
	0x65, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x68, 0x8a, 0xc4, 0x03, 0x64, 0x0a, 0x0c, 0x0a, 0x0a, 0x53, 0x6f, 0x6d, 0x65,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x31, 0x0a, 0x0c, 0x0a, 0x0a, 0x53, 0x6f, 0x6d, 0x65, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x32, 0x12, 0x0d, 0x0a, 0x0b, 0x53, 0x6f, 0x6d, 0x65, 0x53, 0x69, 0x67, 0x6e,
	0x61, 0x6c, 0x31, 0x12, 0x0d, 0x0a, 0x0b, 0x53, 0x6f, 0x6d, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x6c, 0x32, 0x2a, 0x28, 0x73, 0x6f, 0x6d, 0x65, 0x2d, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x2d, 0x31, 0x2f, 0x24, 0x7b, 0x21, 0x20, 0x69, 0x64, 0x20, 0x7d, 0x2f, 0x24, 0x7b, 0x21,
	0x20, 0x75, 0x75, 0x69, 0x64, 0x5f, 0x76, 0x34, 0x28, 0x29, 0x20, 0x7d, 0x12, 0x65, 0x0a, 0x0d,
	0x53, 0x6f, 0x6d, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x32, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x24, 0x8a,
	0xc4, 0x03, 0x20, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x6f, 0x6d, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x6c, 0x31, 0x10, 0x01, 0x1a, 0x0d, 0x0a, 0x0b, 0x53, 0x6f, 0x6d, 0x65, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x31, 0x12, 0xad, 0x01, 0x0a, 0x0d, 0x53, 0x6f, 0x6d, 0x65, 0x57, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x33, 0x12, 0x26, 0x2e, 0x6d, 0x79, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x53, 0x6f, 0x6d, 0x65, 0x57, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x33, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x5c, 0x8a, 0xc4, 0x03, 0x58, 0x12, 0x0f, 0x0a, 0x0b, 0x53,
	0x6f, 0x6d, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x32, 0x10, 0x01, 0x22, 0x03, 0x08, 0x90,
	0x1c, 0x2a, 0x29, 0x73, 0x6f, 0x6d, 0x65, 0x2d, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x2d, 0x33, 0x2f, 0x24, 0x7b, 0x21, 0x20, 0x69, 0x64, 0x20, 0x7d, 0x2f, 0x24, 0x7b, 0x21, 0x20,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x20, 0x7d, 0x30, 0x01, 0x4a, 0x02,
	0x20, 0x02, 0x5a, 0x0f, 0x6d, 0x79, 0x2d, 0x74, 0x61, 0x73, 0x6b, 0x2d, 0x71, 0x75, 0x65, 0x75,
	0x65, 0x2d, 0x32, 0x12, 0x45, 0x0a, 0x0d, 0x53, 0x6f, 0x6d, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x31, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x04, 0x92, 0xc4, 0x03, 0x00, 0x12, 0x5f, 0x0a, 0x0d, 0x53, 0x6f,
	0x6d, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x32, 0x12, 0x26, 0x2e, 0x6d, 0x79,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x53,
	0x6f, 0x6d, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x32, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x0e, 0x92, 0xc4, 0x03,
	0x0a, 0x22, 0x02, 0x08, 0x0a, 0x32, 0x04, 0x1a, 0x02, 0x08, 0x1e, 0x12, 0x6e, 0x0a, 0x0d, 0x53,
	0x6f, 0x6d, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x33, 0x12, 0x26, 0x2e, 0x6d,
	0x79, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x53, 0x6f, 0x6d, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x33, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6d, 0x79, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x53, 0x6f, 0x6d, 0x65, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x33, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0c, 0x92,
	0xc4, 0x03, 0x08, 0x22, 0x02, 0x08, 0x0a, 0x32, 0x02, 0x20, 0x05, 0x12, 0x50, 0x0a, 0x0a, 0x53,
	0x6f, 0x6d, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x31, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x24, 0x2e, 0x6d, 0x79, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x73, 0x69,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x53, 0x6f, 0x6d, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x31, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x04, 0x9a, 0xc4, 0x03, 0x00, 0x12, 0x5d, 0x0a,
	0x0a, 0x53, 0x6f, 0x6d, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x32, 0x12, 0x23, 0x2e, 0x6d, 0x79,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x53,
	0x6f, 0x6d, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x32, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x24, 0x2e, 0x6d, 0x79, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x73, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x53, 0x6f, 0x6d, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x32, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x04, 0x9a, 0xc4, 0x03, 0x00, 0x12, 0x43, 0x0a, 0x0b,
	0x53, 0x6f, 0x6d, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x31, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x04, 0xa2, 0xc4, 0x03,
	0x00, 0x12, 0x51, 0x0a, 0x0b, 0x53, 0x6f, 0x6d, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x32,
	0x12, 0x24, 0x2e, 0x6d, 0x79, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x73, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x53, 0x6f, 0x6d, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x32, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x04,
	0xa2, 0xc4, 0x03, 0x00, 0x12, 0xa6, 0x01, 0x0a, 0x0b, 0x53, 0x6f, 0x6d, 0x65, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x31, 0x12, 0x24, 0x2e, 0x6d, 0x79, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x53, 0x6f, 0x6d, 0x65, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6d, 0x79, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x53, 0x6f,
	0x6d, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x4a, 0xaa, 0xc4, 0x03, 0x46, 0x0a, 0x40, 0x73, 0x6f, 0x6d, 0x65, 0x2d, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x24, 0x7b, 0x21, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x56, 0x61, 0x6c, 0x2e, 0x6e, 0x6f, 0x74, 0x5f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x28, 0x29, 0x2e,
	0x63, 0x61, 0x74, 0x63, 0x68, 0x28, 0x22, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x22, 0x29,
	0x2e, 0x73, 0x6c, 0x75, 0x67, 0x28, 0x29, 0x20, 0x7d, 0x10, 0x01, 0x18, 0x03, 0x1a, 0x13, 0x8a,
	0xc4, 0x03, 0x0f, 0x0a, 0x0d, 0x6d, 0x79, 0x2d, 0x74, 0x61, 0x73, 0x6b, 0x2d, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x42, 0xc2, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x79, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x42, 0x0b, 0x53, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x75, 0x64, 0x64, 0x65, 0x6e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x74, 0x65, 0x6d,
	0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x73,
	0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4d, 0x53, 0x58, 0xaa, 0x02,
	0x10, 0x4d, 0x79, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0xca, 0x02, 0x10, 0x4d, 0x79, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5c, 0x53, 0x69,
	0x6d, 0x70, 0x6c, 0x65, 0xe2, 0x02, 0x1c, 0x4d, 0x79, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x5c, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x4d, 0x79, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x3a,
	0x3a, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test_simple_v1_simple_proto_rawDescOnce sync.Once
	file_test_simple_v1_simple_proto_rawDescData = file_test_simple_v1_simple_proto_rawDesc
)

func file_test_simple_v1_simple_proto_rawDescGZIP() []byte {
	file_test_simple_v1_simple_proto_rawDescOnce.Do(func() {
		file_test_simple_v1_simple_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_simple_v1_simple_proto_rawDescData)
	})
	return file_test_simple_v1_simple_proto_rawDescData
}

var file_test_simple_v1_simple_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_test_simple_v1_simple_proto_goTypes = []interface{}{
	(*SomeWorkflow1Request)(nil),  // 0: mycompany.simple.SomeWorkflow1Request
	(*SomeWorkflow1Response)(nil), // 1: mycompany.simple.SomeWorkflow1Response
	(*SomeWorkflow3Request)(nil),  // 2: mycompany.simple.SomeWorkflow3Request
	(*SomeActivity2Request)(nil),  // 3: mycompany.simple.SomeActivity2Request
	(*SomeActivity3Request)(nil),  // 4: mycompany.simple.SomeActivity3Request
	(*SomeActivity3Response)(nil), // 5: mycompany.simple.SomeActivity3Response
	(*SomeQuery1Response)(nil),    // 6: mycompany.simple.SomeQuery1Response
	(*SomeQuery2Request)(nil),     // 7: mycompany.simple.SomeQuery2Request
	(*SomeQuery2Response)(nil),    // 8: mycompany.simple.SomeQuery2Response
	(*SomeSignal2Request)(nil),    // 9: mycompany.simple.SomeSignal2Request
	(*SomeUpdate1Request)(nil),    // 10: mycompany.simple.SomeUpdate1Request
	(*SomeUpdate1Response)(nil),   // 11: mycompany.simple.SomeUpdate1Response
	(*emptypb.Empty)(nil),         // 12: google.protobuf.Empty
}
var file_test_simple_v1_simple_proto_depIdxs = []int32{
	0,  // 0: mycompany.simple.Simple.SomeWorkflow1:input_type -> mycompany.simple.SomeWorkflow1Request
	12, // 1: mycompany.simple.Simple.SomeWorkflow2:input_type -> google.protobuf.Empty
	2,  // 2: mycompany.simple.Simple.SomeWorkflow3:input_type -> mycompany.simple.SomeWorkflow3Request
	12, // 3: mycompany.simple.Simple.SomeActivity1:input_type -> google.protobuf.Empty
	3,  // 4: mycompany.simple.Simple.SomeActivity2:input_type -> mycompany.simple.SomeActivity2Request
	4,  // 5: mycompany.simple.Simple.SomeActivity3:input_type -> mycompany.simple.SomeActivity3Request
	12, // 6: mycompany.simple.Simple.SomeQuery1:input_type -> google.protobuf.Empty
	7,  // 7: mycompany.simple.Simple.SomeQuery2:input_type -> mycompany.simple.SomeQuery2Request
	12, // 8: mycompany.simple.Simple.SomeSignal1:input_type -> google.protobuf.Empty
	9,  // 9: mycompany.simple.Simple.SomeSignal2:input_type -> mycompany.simple.SomeSignal2Request
	10, // 10: mycompany.simple.Simple.SomeUpdate1:input_type -> mycompany.simple.SomeUpdate1Request
	1,  // 11: mycompany.simple.Simple.SomeWorkflow1:output_type -> mycompany.simple.SomeWorkflow1Response
	12, // 12: mycompany.simple.Simple.SomeWorkflow2:output_type -> google.protobuf.Empty
	12, // 13: mycompany.simple.Simple.SomeWorkflow3:output_type -> google.protobuf.Empty
	12, // 14: mycompany.simple.Simple.SomeActivity1:output_type -> google.protobuf.Empty
	12, // 15: mycompany.simple.Simple.SomeActivity2:output_type -> google.protobuf.Empty
	5,  // 16: mycompany.simple.Simple.SomeActivity3:output_type -> mycompany.simple.SomeActivity3Response
	6,  // 17: mycompany.simple.Simple.SomeQuery1:output_type -> mycompany.simple.SomeQuery1Response
	8,  // 18: mycompany.simple.Simple.SomeQuery2:output_type -> mycompany.simple.SomeQuery2Response
	12, // 19: mycompany.simple.Simple.SomeSignal1:output_type -> google.protobuf.Empty
	12, // 20: mycompany.simple.Simple.SomeSignal2:output_type -> google.protobuf.Empty
	11, // 21: mycompany.simple.Simple.SomeUpdate1:output_type -> mycompany.simple.SomeUpdate1Response
	11, // [11:22] is the sub-list for method output_type
	0,  // [0:11] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_test_simple_v1_simple_proto_init() }
func file_test_simple_v1_simple_proto_init() {
	if File_test_simple_v1_simple_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_simple_v1_simple_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SomeWorkflow1Request); i {
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
		file_test_simple_v1_simple_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SomeWorkflow1Response); i {
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
		file_test_simple_v1_simple_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SomeWorkflow3Request); i {
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
		file_test_simple_v1_simple_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SomeActivity2Request); i {
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
		file_test_simple_v1_simple_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SomeActivity3Request); i {
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
		file_test_simple_v1_simple_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SomeActivity3Response); i {
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
		file_test_simple_v1_simple_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SomeQuery1Response); i {
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
		file_test_simple_v1_simple_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SomeQuery2Request); i {
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
		file_test_simple_v1_simple_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SomeQuery2Response); i {
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
		file_test_simple_v1_simple_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SomeSignal2Request); i {
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
		file_test_simple_v1_simple_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SomeUpdate1Request); i {
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
		file_test_simple_v1_simple_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SomeUpdate1Response); i {
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
			RawDescriptor: file_test_simple_v1_simple_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_test_simple_v1_simple_proto_goTypes,
		DependencyIndexes: file_test_simple_v1_simple_proto_depIdxs,
		MessageInfos:      file_test_simple_v1_simple_proto_msgTypes,
	}.Build()
	File_test_simple_v1_simple_proto = out.File
	file_test_simple_v1_simple_proto_rawDesc = nil
	file_test_simple_v1_simple_proto_goTypes = nil
	file_test_simple_v1_simple_proto_depIdxs = nil
}
