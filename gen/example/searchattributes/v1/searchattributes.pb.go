// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        (unknown)
// source: example/searchattributes/v1/searchattributes.proto

package searchattributesv1

import (
	_ "github.com/cludden/protoc-gen-go-temporal/gen/temporal/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SearchAttributesInput struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	CustomKeywordField  string                 `protobuf:"bytes,1,opt,name=custom_keyword_field,json=customKeywordField,proto3" json:"custom_keyword_field,omitempty"`
	CustomTextField     string                 `protobuf:"bytes,2,opt,name=custom_text_field,json=customTextField,proto3" json:"custom_text_field,omitempty"`
	CustomIntField      int64                  `protobuf:"varint,3,opt,name=custom_int_field,json=customIntField,proto3" json:"custom_int_field,omitempty"`
	CustomDoubleField   float64                `protobuf:"fixed64,4,opt,name=custom_double_field,json=customDoubleField,proto3" json:"custom_double_field,omitempty"`
	CustomBoolField     bool                   `protobuf:"varint,5,opt,name=custom_bool_field,json=customBoolField,proto3" json:"custom_bool_field,omitempty"`
	CustomDatetimeField *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=custom_datetime_field,json=customDatetimeField,proto3" json:"custom_datetime_field,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *SearchAttributesInput) Reset() {
	*x = SearchAttributesInput{}
	mi := &file_example_searchattributes_v1_searchattributes_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchAttributesInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchAttributesInput) ProtoMessage() {}

func (x *SearchAttributesInput) ProtoReflect() protoreflect.Message {
	mi := &file_example_searchattributes_v1_searchattributes_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchAttributesInput.ProtoReflect.Descriptor instead.
func (*SearchAttributesInput) Descriptor() ([]byte, []int) {
	return file_example_searchattributes_v1_searchattributes_proto_rawDescGZIP(), []int{0}
}

func (x *SearchAttributesInput) GetCustomKeywordField() string {
	if x != nil {
		return x.CustomKeywordField
	}
	return ""
}

func (x *SearchAttributesInput) GetCustomTextField() string {
	if x != nil {
		return x.CustomTextField
	}
	return ""
}

func (x *SearchAttributesInput) GetCustomIntField() int64 {
	if x != nil {
		return x.CustomIntField
	}
	return 0
}

func (x *SearchAttributesInput) GetCustomDoubleField() float64 {
	if x != nil {
		return x.CustomDoubleField
	}
	return 0
}

func (x *SearchAttributesInput) GetCustomBoolField() bool {
	if x != nil {
		return x.CustomBoolField
	}
	return false
}

func (x *SearchAttributesInput) GetCustomDatetimeField() *timestamppb.Timestamp {
	if x != nil {
		return x.CustomDatetimeField
	}
	return nil
}

var File_example_searchattributes_v1_searchattributes_proto protoreflect.FileDescriptor

var file_example_searchattributes_v1_searchattributes_proto_rawDesc = []byte{
	0x0a, 0x32, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1a, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6d,
	0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcb, 0x02, 0x0a, 0x15,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f,
	0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x12, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x54, 0x65, 0x78, 0x74, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x69, 0x6e,
	0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x49, 0x6e, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x2e, 0x0a,
	0x13, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x11, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x2a, 0x0a,
	0x11, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x42, 0x6f, 0x6f, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x4e, 0x0a, 0x15, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x13, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x61, 0x74, 0x65,
	0x74, 0x69, 0x6d, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x32, 0xbb, 0x03, 0x0a, 0x07, 0x45, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x97, 0x03, 0x0a, 0x10, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x32, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x61, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0xb6, 0x02, 0x8a, 0xc4, 0x03, 0xb1, 0x02, 0x2a, 0x21,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x5f, 0x24, 0x7b, 0x21, 0x20, 0x75, 0x75, 0x69, 0x64, 0x5f, 0x76, 0x34, 0x28, 0x29, 0x20,
	0x7d, 0x7a, 0x8b, 0x02, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x20, 0x3d, 0x20, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x20, 0x0a, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x54, 0x65, 0x78, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x20, 0x3d, 0x20, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x54, 0x65, 0x78, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x20, 0x0a,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x49, 0x6e, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x20, 0x3d,
	0x20, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x49, 0x6e, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x2e,
	0x69, 0x6e, 0x74, 0x36, 0x34, 0x28, 0x29, 0x20, 0x0a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x20, 0x3d, 0x20, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x20, 0x0a,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x42, 0x6f, 0x6f, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x20,
	0x3d, 0x20, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x42, 0x6f, 0x6f, 0x6c, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x20, 0x0a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d,
	0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x20, 0x3d, 0x20, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44,
	0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x74, 0x73, 0x5f,
	0x70, 0x61, 0x72, 0x73, 0x65, 0x28, 0x22, 0x32, 0x30, 0x30, 0x36, 0x2d, 0x30, 0x31, 0x2d, 0x30,
	0x32, 0x54, 0x31, 0x35, 0x3a, 0x30, 0x34, 0x3a, 0x30, 0x35, 0x5a, 0x22, 0x29, 0x20, 0x0a, 0x1a,
	0x16, 0x8a, 0xc4, 0x03, 0x12, 0x0a, 0x10, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x61, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x42, 0xa4, 0x02, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x61, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x15, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6c, 0x75, 0x64, 0x64, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d,
	0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x45, 0x53, 0x58, 0xaa, 0x02, 0x1b, 0x45, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1b, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x5c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x27, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x1d, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x3a, 0x3a, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_example_searchattributes_v1_searchattributes_proto_rawDescOnce sync.Once
	file_example_searchattributes_v1_searchattributes_proto_rawDescData = file_example_searchattributes_v1_searchattributes_proto_rawDesc
)

func file_example_searchattributes_v1_searchattributes_proto_rawDescGZIP() []byte {
	file_example_searchattributes_v1_searchattributes_proto_rawDescOnce.Do(func() {
		file_example_searchattributes_v1_searchattributes_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_searchattributes_v1_searchattributes_proto_rawDescData)
	})
	return file_example_searchattributes_v1_searchattributes_proto_rawDescData
}

var file_example_searchattributes_v1_searchattributes_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_example_searchattributes_v1_searchattributes_proto_goTypes = []any{
	(*SearchAttributesInput)(nil), // 0: example.searchattributes.v1.SearchAttributesInput
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 2: google.protobuf.Empty
}
var file_example_searchattributes_v1_searchattributes_proto_depIdxs = []int32{
	1, // 0: example.searchattributes.v1.SearchAttributesInput.custom_datetime_field:type_name -> google.protobuf.Timestamp
	0, // 1: example.searchattributes.v1.Example.SearchAttributes:input_type -> example.searchattributes.v1.SearchAttributesInput
	2, // 2: example.searchattributes.v1.Example.SearchAttributes:output_type -> google.protobuf.Empty
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_example_searchattributes_v1_searchattributes_proto_init() }
func file_example_searchattributes_v1_searchattributes_proto_init() {
	if File_example_searchattributes_v1_searchattributes_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_example_searchattributes_v1_searchattributes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_example_searchattributes_v1_searchattributes_proto_goTypes,
		DependencyIndexes: file_example_searchattributes_v1_searchattributes_proto_depIdxs,
		MessageInfos:      file_example_searchattributes_v1_searchattributes_proto_msgTypes,
	}.Build()
	File_example_searchattributes_v1_searchattributes_proto = out.File
	file_example_searchattributes_v1_searchattributes_proto_rawDesc = nil
	file_example_searchattributes_v1_searchattributes_proto_goTypes = nil
	file_example_searchattributes_v1_searchattributes_proto_depIdxs = nil
}
