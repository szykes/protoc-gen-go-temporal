// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        (unknown)
// source: example/mutex/v1/mutex.proto

package mutexv1

import (
	_ "github.com/cludden/protoc-gen-go-temporal/gen/temporal/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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

type MutexInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ResourceId    string                 `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MutexInput) Reset() {
	*x = MutexInput{}
	mi := &file_example_mutex_v1_mutex_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MutexInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutexInput) ProtoMessage() {}

func (x *MutexInput) ProtoReflect() protoreflect.Message {
	mi := &file_example_mutex_v1_mutex_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutexInput.ProtoReflect.Descriptor instead.
func (*MutexInput) Descriptor() ([]byte, []int) {
	return file_example_mutex_v1_mutex_proto_rawDescGZIP(), []int{0}
}

func (x *MutexInput) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

type AcquireLockInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	WorkflowId    string                 `protobuf:"bytes,1,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id,omitempty"`
	Timeout       *durationpb.Duration   `protobuf:"bytes,2,opt,name=timeout,proto3" json:"timeout,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AcquireLockInput) Reset() {
	*x = AcquireLockInput{}
	mi := &file_example_mutex_v1_mutex_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AcquireLockInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcquireLockInput) ProtoMessage() {}

func (x *AcquireLockInput) ProtoReflect() protoreflect.Message {
	mi := &file_example_mutex_v1_mutex_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcquireLockInput.ProtoReflect.Descriptor instead.
func (*AcquireLockInput) Descriptor() ([]byte, []int) {
	return file_example_mutex_v1_mutex_proto_rawDescGZIP(), []int{1}
}

func (x *AcquireLockInput) GetWorkflowId() string {
	if x != nil {
		return x.WorkflowId
	}
	return ""
}

func (x *AcquireLockInput) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

type LockAcquiredInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	LeaseId       string                 `protobuf:"bytes,1,opt,name=lease_id,json=leaseId,proto3" json:"lease_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LockAcquiredInput) Reset() {
	*x = LockAcquiredInput{}
	mi := &file_example_mutex_v1_mutex_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LockAcquiredInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LockAcquiredInput) ProtoMessage() {}

func (x *LockAcquiredInput) ProtoReflect() protoreflect.Message {
	mi := &file_example_mutex_v1_mutex_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LockAcquiredInput.ProtoReflect.Descriptor instead.
func (*LockAcquiredInput) Descriptor() ([]byte, []int) {
	return file_example_mutex_v1_mutex_proto_rawDescGZIP(), []int{2}
}

func (x *LockAcquiredInput) GetLeaseId() string {
	if x != nil {
		return x.LeaseId
	}
	return ""
}

type ReleaseLockInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	LeaseId       string                 `protobuf:"bytes,1,opt,name=lease_id,json=leaseId,proto3" json:"lease_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReleaseLockInput) Reset() {
	*x = ReleaseLockInput{}
	mi := &file_example_mutex_v1_mutex_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReleaseLockInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleaseLockInput) ProtoMessage() {}

func (x *ReleaseLockInput) ProtoReflect() protoreflect.Message {
	mi := &file_example_mutex_v1_mutex_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleaseLockInput.ProtoReflect.Descriptor instead.
func (*ReleaseLockInput) Descriptor() ([]byte, []int) {
	return file_example_mutex_v1_mutex_proto_rawDescGZIP(), []int{3}
}

func (x *ReleaseLockInput) GetLeaseId() string {
	if x != nil {
		return x.LeaseId
	}
	return ""
}

type SampleWorkflowWithMutexInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ResourceId    string                 `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	Sleep         *durationpb.Duration   `protobuf:"bytes,2,opt,name=sleep,proto3" json:"sleep,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SampleWorkflowWithMutexInput) Reset() {
	*x = SampleWorkflowWithMutexInput{}
	mi := &file_example_mutex_v1_mutex_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SampleWorkflowWithMutexInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SampleWorkflowWithMutexInput) ProtoMessage() {}

func (x *SampleWorkflowWithMutexInput) ProtoReflect() protoreflect.Message {
	mi := &file_example_mutex_v1_mutex_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SampleWorkflowWithMutexInput.ProtoReflect.Descriptor instead.
func (*SampleWorkflowWithMutexInput) Descriptor() ([]byte, []int) {
	return file_example_mutex_v1_mutex_proto_rawDescGZIP(), []int{4}
}

func (x *SampleWorkflowWithMutexInput) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *SampleWorkflowWithMutexInput) GetSleep() *durationpb.Duration {
	if x != nil {
		return x.Sleep
	}
	return nil
}

var File_example_mutex_v1_mutex_proto protoreflect.FileDescriptor

var file_example_mutex_v1_mutex_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x6d, 0x75, 0x74, 0x65, 0x78, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x75, 0x74, 0x65, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6d, 0x75, 0x74, 0x65, 0x78, 0x2e, 0x76, 0x31,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x74,
	0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6f,
	0x72, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x0a, 0x4d, 0x75, 0x74,
	0x65, 0x78, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x22, 0x68, 0x0a, 0x10, 0x41, 0x63, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x4c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x12, 0x33, 0x0a,
	0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x22, 0x2e, 0x0a, 0x11, 0x4c, 0x6f, 0x63, 0x6b, 0x41, 0x63, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x49, 0x64, 0x22, 0x2d, 0x0a, 0x10, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x4c, 0x6f, 0x63,
	0x6b, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x49,
	0x64, 0x22, 0x70, 0x0a, 0x1c, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x57, 0x69, 0x74, 0x68, 0x4d, 0x75, 0x74, 0x65, 0x78, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x2f, 0x0a, 0x05, 0x73, 0x6c, 0x65, 0x65, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x73, 0x6c,
	0x65, 0x65, 0x70, 0x32, 0xca, 0x04, 0x0a, 0x07, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12,
	0x4f, 0x0a, 0x0b, 0x41, 0x63, 0x71, 0x75, 0x69, 0x72, 0x65, 0x4c, 0x6f, 0x63, 0x6b, 0x12, 0x22,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6d, 0x75, 0x74, 0x65, 0x78, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x63, 0x71, 0x75, 0x69, 0x72, 0x65, 0x4c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x04, 0xa2, 0xc4, 0x03, 0x00,
	0x12, 0x51, 0x0a, 0x0c, 0x4c, 0x6f, 0x63, 0x6b, 0x41, 0x63, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64,
	0x12, 0x23, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6d, 0x75, 0x74, 0x65, 0x78,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x41, 0x63, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x04, 0xa2,
	0xc4, 0x03, 0x00, 0x12, 0x98, 0x01, 0x0a, 0x05, 0x4d, 0x75, 0x74, 0x65, 0x78, 0x12, 0x1c, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6d, 0x75, 0x74, 0x65, 0x78, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x75, 0x74, 0x65, 0x78, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x59, 0x8a, 0xc4, 0x03, 0x4d, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x63, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x4c, 0x6f, 0x63, 0x6b, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x0b, 0x52, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x4c, 0x6f, 0x63, 0x6b, 0x2a, 0x16, 0x6d, 0x75, 0x74, 0x65, 0x78,
	0x3a, 0x24, 0x7b, 0x21, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x20,
	0x7d, 0x4a, 0x13, 0x0a, 0x02, 0x08, 0x01, 0x11, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x40,
	0x1a, 0x02, 0x08, 0x3c, 0x20, 0x05, 0x92, 0xc4, 0x03, 0x04, 0x22, 0x02, 0x08, 0x0a, 0x12, 0x4f,
	0x0a, 0x0b, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x4c, 0x6f, 0x63, 0x6b, 0x12, 0x22, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6d, 0x75, 0x74, 0x65, 0x78, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x4c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x04, 0xa2, 0xc4, 0x03, 0x00, 0x12,
	0xa1, 0x01, 0x0a, 0x17, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x57, 0x69, 0x74, 0x68, 0x4d, 0x75, 0x74, 0x65, 0x78, 0x12, 0x2e, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6d, 0x75, 0x74, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x57, 0x69, 0x74,
	0x68, 0x4d, 0x75, 0x74, 0x65, 0x78, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x3e, 0x8a, 0xc4, 0x03, 0x3a, 0x12, 0x0e, 0x0a, 0x0c, 0x4c, 0x6f, 0x63,
	0x6b, 0x41, 0x63, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x2a, 0x28, 0x53, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x31, 0x57, 0x69, 0x74, 0x68, 0x4d, 0x75,
	0x74, 0x65, 0x78, 0x5f, 0x24, 0x7b, 0x21, 0x20, 0x75, 0x75, 0x69, 0x64, 0x5f, 0x76, 0x34, 0x28,
	0x29, 0x20, 0x7d, 0x1a, 0x0b, 0x8a, 0xc4, 0x03, 0x07, 0x0a, 0x05, 0x6d, 0x75, 0x74, 0x65, 0x78,
	0x42, 0xcc, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x6d, 0x75, 0x74, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x4d, 0x75, 0x74, 0x65, 0x78,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x75, 0x64, 0x64, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72,
	0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x6d,
	0x75, 0x74, 0x65, 0x78, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x75, 0x74, 0x65, 0x78, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x45, 0x4d, 0x58, 0xaa, 0x02, 0x10, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x4d, 0x75, 0x74, 0x65, 0x78, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x10, 0x45, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x5c, 0x4d, 0x75, 0x74, 0x65, 0x78, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1c, 0x45, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x4d, 0x75, 0x74, 0x65, 0x78, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x12, 0x45, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x3a, 0x3a, 0x4d, 0x75, 0x74, 0x65, 0x78, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_example_mutex_v1_mutex_proto_rawDescOnce sync.Once
	file_example_mutex_v1_mutex_proto_rawDescData = file_example_mutex_v1_mutex_proto_rawDesc
)

func file_example_mutex_v1_mutex_proto_rawDescGZIP() []byte {
	file_example_mutex_v1_mutex_proto_rawDescOnce.Do(func() {
		file_example_mutex_v1_mutex_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_mutex_v1_mutex_proto_rawDescData)
	})
	return file_example_mutex_v1_mutex_proto_rawDescData
}

var file_example_mutex_v1_mutex_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_example_mutex_v1_mutex_proto_goTypes = []any{
	(*MutexInput)(nil),                   // 0: example.mutex.v1.MutexInput
	(*AcquireLockInput)(nil),             // 1: example.mutex.v1.AcquireLockInput
	(*LockAcquiredInput)(nil),            // 2: example.mutex.v1.LockAcquiredInput
	(*ReleaseLockInput)(nil),             // 3: example.mutex.v1.ReleaseLockInput
	(*SampleWorkflowWithMutexInput)(nil), // 4: example.mutex.v1.SampleWorkflowWithMutexInput
	(*durationpb.Duration)(nil),          // 5: google.protobuf.Duration
	(*emptypb.Empty)(nil),                // 6: google.protobuf.Empty
}
var file_example_mutex_v1_mutex_proto_depIdxs = []int32{
	5, // 0: example.mutex.v1.AcquireLockInput.timeout:type_name -> google.protobuf.Duration
	5, // 1: example.mutex.v1.SampleWorkflowWithMutexInput.sleep:type_name -> google.protobuf.Duration
	1, // 2: example.mutex.v1.Example.AcquireLock:input_type -> example.mutex.v1.AcquireLockInput
	2, // 3: example.mutex.v1.Example.LockAcquired:input_type -> example.mutex.v1.LockAcquiredInput
	0, // 4: example.mutex.v1.Example.Mutex:input_type -> example.mutex.v1.MutexInput
	3, // 5: example.mutex.v1.Example.ReleaseLock:input_type -> example.mutex.v1.ReleaseLockInput
	4, // 6: example.mutex.v1.Example.SampleWorkflowWithMutex:input_type -> example.mutex.v1.SampleWorkflowWithMutexInput
	6, // 7: example.mutex.v1.Example.AcquireLock:output_type -> google.protobuf.Empty
	6, // 8: example.mutex.v1.Example.LockAcquired:output_type -> google.protobuf.Empty
	6, // 9: example.mutex.v1.Example.Mutex:output_type -> google.protobuf.Empty
	6, // 10: example.mutex.v1.Example.ReleaseLock:output_type -> google.protobuf.Empty
	6, // 11: example.mutex.v1.Example.SampleWorkflowWithMutex:output_type -> google.protobuf.Empty
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_example_mutex_v1_mutex_proto_init() }
func file_example_mutex_v1_mutex_proto_init() {
	if File_example_mutex_v1_mutex_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_example_mutex_v1_mutex_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_example_mutex_v1_mutex_proto_goTypes,
		DependencyIndexes: file_example_mutex_v1_mutex_proto_depIdxs,
		MessageInfos:      file_example_mutex_v1_mutex_proto_msgTypes,
	}.Build()
	File_example_mutex_v1_mutex_proto = out.File
	file_example_mutex_v1_mutex_proto_rawDesc = nil
	file_example_mutex_v1_mutex_proto_goTypes = nil
	file_example_mutex_v1_mutex_proto_depIdxs = nil
}
