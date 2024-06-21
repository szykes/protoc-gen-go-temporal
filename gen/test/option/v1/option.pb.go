// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: test/option/v1/option.proto

package optionv1

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

type ActivityWithInputRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ActivityWithInputRequest) Reset() {
	*x = ActivityWithInputRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_option_v1_option_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityWithInputRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityWithInputRequest) ProtoMessage() {}

func (x *ActivityWithInputRequest) ProtoReflect() protoreflect.Message {
	mi := &file_test_option_v1_option_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityWithInputRequest.ProtoReflect.Descriptor instead.
func (*ActivityWithInputRequest) Descriptor() ([]byte, []int) {
	return file_test_option_v1_option_proto_rawDescGZIP(), []int{0}
}

func (x *ActivityWithInputRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ActivityWithInputResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ActivityWithInputResponse) Reset() {
	*x = ActivityWithInputResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_option_v1_option_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityWithInputResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityWithInputResponse) ProtoMessage() {}

func (x *ActivityWithInputResponse) ProtoReflect() protoreflect.Message {
	mi := &file_test_option_v1_option_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityWithInputResponse.ProtoReflect.Descriptor instead.
func (*ActivityWithInputResponse) Descriptor() ([]byte, []int) {
	return file_test_option_v1_option_proto_rawDescGZIP(), []int{1}
}

func (x *ActivityWithInputResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type UpdateWithInputRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UpdateWithInputRequest) Reset() {
	*x = UpdateWithInputRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_option_v1_option_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateWithInputRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateWithInputRequest) ProtoMessage() {}

func (x *UpdateWithInputRequest) ProtoReflect() protoreflect.Message {
	mi := &file_test_option_v1_option_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateWithInputRequest.ProtoReflect.Descriptor instead.
func (*UpdateWithInputRequest) Descriptor() ([]byte, []int) {
	return file_test_option_v1_option_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateWithInputRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type WorkflowWithInputRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *WorkflowWithInputRequest) Reset() {
	*x = WorkflowWithInputRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_option_v1_option_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowWithInputRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowWithInputRequest) ProtoMessage() {}

func (x *WorkflowWithInputRequest) ProtoReflect() protoreflect.Message {
	mi := &file_test_option_v1_option_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowWithInputRequest.ProtoReflect.Descriptor instead.
func (*WorkflowWithInputRequest) Descriptor() ([]byte, []int) {
	return file_test_option_v1_option_proto_rawDescGZIP(), []int{3}
}

func (x *WorkflowWithInputRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_test_option_v1_option_proto protoreflect.FileDescriptor

var file_test_option_v1_option_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x74, 0x65, 0x6d, 0x70,
	0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x18, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x57, 0x69, 0x74, 0x68, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x33, 0x0a, 0x19, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x57, 0x69, 0x74, 0x68, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x2c, 0x0a, 0x16, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x69, 0x74, 0x68, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2e, 0x0a, 0x18, 0x57, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x57, 0x69, 0x74, 0x68, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0xa8, 0x04, 0x0a, 0x04, 0x54, 0x65,
	0x73, 0x74, 0x12, 0x91, 0x01, 0x0a, 0x11, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x57,
	0x69, 0x74, 0x68, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x28, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x57, 0x69, 0x74, 0x68, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x29, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x57, 0x69, 0x74, 0x68,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x92,
	0xc4, 0x03, 0x23, 0x0a, 0x09, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x76, 0x32, 0x12, 0x02,
	0x08, 0x78, 0x1a, 0x02, 0x08, 0x0a, 0x22, 0x02, 0x08, 0x3c, 0x2a, 0x02, 0x08, 0x1e, 0x32, 0x04,
	0x1a, 0x02, 0x08, 0x05, 0x40, 0x01, 0x12, 0x97, 0x01, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x57, 0x69, 0x74, 0x68, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x26, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x57, 0x69, 0x74, 0x68, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x44, 0xaa, 0xc4, 0x03, 0x40,
	0x0a, 0x3a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2d, 0x77, 0x69, 0x74, 0x68, 0x2d, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x3a, 0x24, 0x7b, 0x21, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x6f, 0x72, 0x28,
	0x74, 0x68, 0x72, 0x6f, 0x77, 0x28, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x20, 0x69, 0x73, 0x20, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x29, 0x29, 0x20, 0x7d, 0x18, 0x01, 0x38, 0x03,
	0x12, 0xe0, 0x01, 0x0a, 0x11, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x57, 0x69, 0x74,
	0x68, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x28, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x57, 0x69, 0x74, 0x68, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x88, 0x01, 0x8a, 0xc4, 0x03, 0x83, 0x01,
	0x1a, 0x11, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x69, 0x74, 0x68, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x22, 0x03, 0x08, 0xd8, 0x04, 0x2a, 0x3c, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x2d, 0x77, 0x69, 0x74, 0x68, 0x2d, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x3a, 0x24, 0x7b,
	0x21, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x6f, 0x72, 0x28, 0x74, 0x68, 0x72, 0x6f, 0x77, 0x28,
	0x22, 0x6e, 0x61, 0x6d, 0x65, 0x20, 0x69, 0x73, 0x20, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x64, 0x22, 0x29, 0x29, 0x20, 0x7d, 0x30, 0x02, 0x40, 0x03, 0x4a, 0x02, 0x20, 0x05, 0x52, 0x03,
	0x08, 0xac, 0x02, 0x5a, 0x09, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x76, 0x32, 0x62, 0x02,
	0x08, 0x0a, 0x68, 0x01, 0x7a, 0x0d, 0x6e, 0x61, 0x6d, 0x65, 0x20, 0x3d, 0x20, 0x6e, 0x61, 0x6d,
	0x65, 0x20, 0x0a, 0x1a, 0x0f, 0x8a, 0xc4, 0x03, 0x0b, 0x0a, 0x09, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x2d, 0x76, 0x31, 0x42, 0xc2, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x75, 0x64, 0x64, 0x65, 0x6e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x74, 0x65, 0x6d,
	0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x54, 0x4f, 0x58, 0xaa, 0x02, 0x0e, 0x54, 0x65, 0x73, 0x74, 0x2e, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0e, 0x54, 0x65, 0x73, 0x74, 0x5c,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1a, 0x54, 0x65, 0x73, 0x74,
	0x5c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x54, 0x65, 0x73, 0x74, 0x3a, 0x3a, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_test_option_v1_option_proto_rawDescOnce sync.Once
	file_test_option_v1_option_proto_rawDescData = file_test_option_v1_option_proto_rawDesc
)

func file_test_option_v1_option_proto_rawDescGZIP() []byte {
	file_test_option_v1_option_proto_rawDescOnce.Do(func() {
		file_test_option_v1_option_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_option_v1_option_proto_rawDescData)
	})
	return file_test_option_v1_option_proto_rawDescData
}

var file_test_option_v1_option_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_test_option_v1_option_proto_goTypes = []any{
	(*ActivityWithInputRequest)(nil),  // 0: test.option.v1.ActivityWithInputRequest
	(*ActivityWithInputResponse)(nil), // 1: test.option.v1.ActivityWithInputResponse
	(*UpdateWithInputRequest)(nil),    // 2: test.option.v1.UpdateWithInputRequest
	(*WorkflowWithInputRequest)(nil),  // 3: test.option.v1.WorkflowWithInputRequest
	(*emptypb.Empty)(nil),             // 4: google.protobuf.Empty
}
var file_test_option_v1_option_proto_depIdxs = []int32{
	0, // 0: test.option.v1.Test.ActivityWithInput:input_type -> test.option.v1.ActivityWithInputRequest
	2, // 1: test.option.v1.Test.UpdateWithInput:input_type -> test.option.v1.UpdateWithInputRequest
	3, // 2: test.option.v1.Test.WorkflowWithInput:input_type -> test.option.v1.WorkflowWithInputRequest
	1, // 3: test.option.v1.Test.ActivityWithInput:output_type -> test.option.v1.ActivityWithInputResponse
	4, // 4: test.option.v1.Test.UpdateWithInput:output_type -> google.protobuf.Empty
	4, // 5: test.option.v1.Test.WorkflowWithInput:output_type -> google.protobuf.Empty
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_test_option_v1_option_proto_init() }
func file_test_option_v1_option_proto_init() {
	if File_test_option_v1_option_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_option_v1_option_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ActivityWithInputRequest); i {
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
		file_test_option_v1_option_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ActivityWithInputResponse); i {
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
		file_test_option_v1_option_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateWithInputRequest); i {
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
		file_test_option_v1_option_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*WorkflowWithInputRequest); i {
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
			RawDescriptor: file_test_option_v1_option_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_test_option_v1_option_proto_goTypes,
		DependencyIndexes: file_test_option_v1_option_proto_depIdxs,
		MessageInfos:      file_test_option_v1_option_proto_msgTypes,
	}.Build()
	File_test_option_v1_option_proto = out.File
	file_test_option_v1_option_proto_rawDesc = nil
	file_test_option_v1_option_proto_goTypes = nil
	file_test_option_v1_option_proto_depIdxs = nil
}
