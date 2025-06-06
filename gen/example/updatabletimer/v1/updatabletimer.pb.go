// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        (unknown)
// source: example/updatabletimer/v1/updatabletimer.proto

package updatabletimerv1

import (
	_ "github.com/cludden/protoc-gen-go-temporal/gen/temporal/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// GetWakeUpTimeOutput describes the input to a GetWakeUpTime query
type GetWakeUpTimeOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	WakeUpTime    *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=wake_up_time,json=wakeUpTime,proto3" json:"wake_up_time,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetWakeUpTimeOutput) Reset() {
	*x = GetWakeUpTimeOutput{}
	mi := &file_example_updatabletimer_v1_updatabletimer_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetWakeUpTimeOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWakeUpTimeOutput) ProtoMessage() {}

func (x *GetWakeUpTimeOutput) ProtoReflect() protoreflect.Message {
	mi := &file_example_updatabletimer_v1_updatabletimer_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWakeUpTimeOutput.ProtoReflect.Descriptor instead.
func (*GetWakeUpTimeOutput) Descriptor() ([]byte, []int) {
	return file_example_updatabletimer_v1_updatabletimer_proto_rawDescGZIP(), []int{0}
}

func (x *GetWakeUpTimeOutput) GetWakeUpTime() *timestamppb.Timestamp {
	if x != nil {
		return x.WakeUpTime
	}
	return nil
}

// UpdatableTimerInput describes the input to a UpdatableTimer workflow
type UpdatableTimerInput struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	InitialWakeUpTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=initial_wake_up_time,json=initialWakeUpTime,proto3" json:"initial_wake_up_time,omitempty"`
	Name              string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *UpdatableTimerInput) Reset() {
	*x = UpdatableTimerInput{}
	mi := &file_example_updatabletimer_v1_updatabletimer_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdatableTimerInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatableTimerInput) ProtoMessage() {}

func (x *UpdatableTimerInput) ProtoReflect() protoreflect.Message {
	mi := &file_example_updatabletimer_v1_updatabletimer_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatableTimerInput.ProtoReflect.Descriptor instead.
func (*UpdatableTimerInput) Descriptor() ([]byte, []int) {
	return file_example_updatabletimer_v1_updatabletimer_proto_rawDescGZIP(), []int{1}
}

func (x *UpdatableTimerInput) GetInitialWakeUpTime() *timestamppb.Timestamp {
	if x != nil {
		return x.InitialWakeUpTime
	}
	return nil
}

func (x *UpdatableTimerInput) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// UpdateWakeUpTimeInput describes the input to a UpdateWakeUpTime signal
type UpdateWakeUpTimeInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	WakeUpTime    *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=wake_up_time,json=wakeUpTime,proto3" json:"wake_up_time,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateWakeUpTimeInput) Reset() {
	*x = UpdateWakeUpTimeInput{}
	mi := &file_example_updatabletimer_v1_updatabletimer_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateWakeUpTimeInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateWakeUpTimeInput) ProtoMessage() {}

func (x *UpdateWakeUpTimeInput) ProtoReflect() protoreflect.Message {
	mi := &file_example_updatabletimer_v1_updatabletimer_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateWakeUpTimeInput.ProtoReflect.Descriptor instead.
func (*UpdateWakeUpTimeInput) Descriptor() ([]byte, []int) {
	return file_example_updatabletimer_v1_updatabletimer_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateWakeUpTimeInput) GetWakeUpTime() *timestamppb.Timestamp {
	if x != nil {
		return x.WakeUpTime
	}
	return nil
}

var File_example_updatabletimer_v1_updatabletimer_proto protoreflect.FileDescriptor

var file_example_updatabletimer_v1_updatabletimer_proto_rawDesc = string([]byte{
	0x0a, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x19, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x74, 0x65, 0x6d, 0x70, 0x6f,
	0x72, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x53, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x57, 0x61, 0x6b, 0x65,
	0x55, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x3c, 0x0a, 0x0c,
	0x77, 0x61, 0x6b, 0x65, 0x5f, 0x75, 0x70, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a,
	0x77, 0x61, 0x6b, 0x65, 0x55, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x76, 0x0a, 0x13, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x12, 0x4b, 0x0a, 0x14, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x77, 0x61, 0x6b,
	0x65, 0x5f, 0x75, 0x70, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x11, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x6c, 0x57, 0x61, 0x6b, 0x65, 0x55, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x55, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6b, 0x65,
	0x55, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x3c, 0x0a, 0x0c, 0x77,
	0x61, 0x6b, 0x65, 0x5f, 0x75, 0x70, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x77,
	0x61, 0x6b, 0x65, 0x55, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x32, 0xa3, 0x03, 0x0a, 0x07, 0x45, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x5d, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x57, 0x61, 0x6b, 0x65,
	0x55, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2e,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x61,
	0x6b, 0x65, 0x55, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x04,
	0x9a, 0xc4, 0x03, 0x00, 0x12, 0xbd, 0x01, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x12, 0x2e, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x63, 0x8a, 0xc4, 0x03, 0x5f, 0x0a, 0x0f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x57, 0x61, 0x6b, 0x65,
	0x55, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x57, 0x61, 0x6b, 0x65, 0x55, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x2a, 0x28, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x2d, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x2f, 0x24, 0x7b, 0x21, 0x20,
	0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x6f, 0x72, 0x28, 0x75, 0x75, 0x69, 0x64, 0x5f, 0x76, 0x34, 0x28,
	0x29, 0x29, 0x20, 0x7d, 0x72, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x72, 0x12, 0x62, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x61,
	0x6b, 0x65, 0x55, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x30, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x74, 0x69, 0x6d, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6b, 0x65, 0x55,
	0x70, 0x54, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x04, 0xa2, 0xc4, 0x03, 0x00, 0x1a, 0x15, 0x8a, 0xc4, 0x03, 0x11, 0x0a, 0x0f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2d, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x42,
	0x94, 0x02, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x42, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x74, 0x69, 0x6d, 0x65,
	0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x58, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x75, 0x64, 0x64, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x74, 0x65, 0x6d, 0x70, 0x6f,
	0x72, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x3b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x72,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x45, 0x55, 0x58, 0xaa, 0x02, 0x19, 0x45, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x74, 0x69, 0x6d, 0x65,
	0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x19, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x25, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1b, 0x45, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x3a, 0x3a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x74, 0x69, 0x6d,
	0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_example_updatabletimer_v1_updatabletimer_proto_rawDescOnce sync.Once
	file_example_updatabletimer_v1_updatabletimer_proto_rawDescData []byte
)

func file_example_updatabletimer_v1_updatabletimer_proto_rawDescGZIP() []byte {
	file_example_updatabletimer_v1_updatabletimer_proto_rawDescOnce.Do(func() {
		file_example_updatabletimer_v1_updatabletimer_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_example_updatabletimer_v1_updatabletimer_proto_rawDesc), len(file_example_updatabletimer_v1_updatabletimer_proto_rawDesc)))
	})
	return file_example_updatabletimer_v1_updatabletimer_proto_rawDescData
}

var file_example_updatabletimer_v1_updatabletimer_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_example_updatabletimer_v1_updatabletimer_proto_goTypes = []any{
	(*GetWakeUpTimeOutput)(nil),   // 0: example.updatabletimer.v1.GetWakeUpTimeOutput
	(*UpdatableTimerInput)(nil),   // 1: example.updatabletimer.v1.UpdatableTimerInput
	(*UpdateWakeUpTimeInput)(nil), // 2: example.updatabletimer.v1.UpdateWakeUpTimeInput
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 4: google.protobuf.Empty
}
var file_example_updatabletimer_v1_updatabletimer_proto_depIdxs = []int32{
	3, // 0: example.updatabletimer.v1.GetWakeUpTimeOutput.wake_up_time:type_name -> google.protobuf.Timestamp
	3, // 1: example.updatabletimer.v1.UpdatableTimerInput.initial_wake_up_time:type_name -> google.protobuf.Timestamp
	3, // 2: example.updatabletimer.v1.UpdateWakeUpTimeInput.wake_up_time:type_name -> google.protobuf.Timestamp
	4, // 3: example.updatabletimer.v1.Example.GetWakeUpTime:input_type -> google.protobuf.Empty
	1, // 4: example.updatabletimer.v1.Example.UpdatableTimer:input_type -> example.updatabletimer.v1.UpdatableTimerInput
	2, // 5: example.updatabletimer.v1.Example.UpdateWakeUpTime:input_type -> example.updatabletimer.v1.UpdateWakeUpTimeInput
	0, // 6: example.updatabletimer.v1.Example.GetWakeUpTime:output_type -> example.updatabletimer.v1.GetWakeUpTimeOutput
	4, // 7: example.updatabletimer.v1.Example.UpdatableTimer:output_type -> google.protobuf.Empty
	4, // 8: example.updatabletimer.v1.Example.UpdateWakeUpTime:output_type -> google.protobuf.Empty
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_example_updatabletimer_v1_updatabletimer_proto_init() }
func file_example_updatabletimer_v1_updatabletimer_proto_init() {
	if File_example_updatabletimer_v1_updatabletimer_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_example_updatabletimer_v1_updatabletimer_proto_rawDesc), len(file_example_updatabletimer_v1_updatabletimer_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_example_updatabletimer_v1_updatabletimer_proto_goTypes,
		DependencyIndexes: file_example_updatabletimer_v1_updatabletimer_proto_depIdxs,
		MessageInfos:      file_example_updatabletimer_v1_updatabletimer_proto_msgTypes,
	}.Build()
	File_example_updatabletimer_v1_updatabletimer_proto = out.File
	file_example_updatabletimer_v1_updatabletimer_proto_goTypes = nil
	file_example_updatabletimer_v1_updatabletimer_proto_depIdxs = nil
}
