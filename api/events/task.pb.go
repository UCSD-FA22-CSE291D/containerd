//
//Copyright The containerd Authors.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.1
// source: github.com/containerd/containerd/api/events/task.proto

package events

import (
	types "github.com/containerd/containerd/api/types"
	_ "github.com/containerd/containerd/protobuf/plugin"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type TaskCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContainerID string         `protobuf:"bytes,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	Bundle      string         `protobuf:"bytes,2,opt,name=bundle,proto3" json:"bundle,omitempty"`
	Rootfs      []*types.Mount `protobuf:"bytes,3,rep,name=rootfs,proto3" json:"rootfs,omitempty"`
	IO          *TaskIO        `protobuf:"bytes,4,opt,name=io,proto3" json:"io,omitempty"`
	Checkpoint  string         `protobuf:"bytes,5,opt,name=checkpoint,proto3" json:"checkpoint,omitempty"`
	Pid         uint32         `protobuf:"varint,6,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (x *TaskCreate) Reset() {
	*x = TaskCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_containerd_containerd_api_events_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskCreate) ProtoMessage() {}

func (x *TaskCreate) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_containerd_containerd_api_events_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskCreate.ProtoReflect.Descriptor instead.
func (*TaskCreate) Descriptor() ([]byte, []int) {
	return file_github_com_containerd_containerd_api_events_task_proto_rawDescGZIP(), []int{0}
}

func (x *TaskCreate) GetContainerID() string {
	if x != nil {
		return x.ContainerID
	}
	return ""
}

func (x *TaskCreate) GetBundle() string {
	if x != nil {
		return x.Bundle
	}
	return ""
}

func (x *TaskCreate) GetRootfs() []*types.Mount {
	if x != nil {
		return x.Rootfs
	}
	return nil
}

func (x *TaskCreate) GetIO() *TaskIO {
	if x != nil {
		return x.IO
	}
	return nil
}

func (x *TaskCreate) GetCheckpoint() string {
	if x != nil {
		return x.Checkpoint
	}
	return ""
}

func (x *TaskCreate) GetPid() uint32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

type TaskStart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContainerID string `protobuf:"bytes,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	Pid         uint32 `protobuf:"varint,2,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (x *TaskStart) Reset() {
	*x = TaskStart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_containerd_containerd_api_events_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskStart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskStart) ProtoMessage() {}

func (x *TaskStart) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_containerd_containerd_api_events_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskStart.ProtoReflect.Descriptor instead.
func (*TaskStart) Descriptor() ([]byte, []int) {
	return file_github_com_containerd_containerd_api_events_task_proto_rawDescGZIP(), []int{1}
}

func (x *TaskStart) GetContainerID() string {
	if x != nil {
		return x.ContainerID
	}
	return ""
}

func (x *TaskStart) GetPid() uint32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

type TaskDelete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContainerID string                 `protobuf:"bytes,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	Pid         uint32                 `protobuf:"varint,2,opt,name=pid,proto3" json:"pid,omitempty"`
	ExitStatus  uint32                 `protobuf:"varint,3,opt,name=exit_status,json=exitStatus,proto3" json:"exit_status,omitempty"`
	ExitedAt    *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=exited_at,json=exitedAt,proto3" json:"exited_at,omitempty"`
	// id is the specific exec. By default if omitted will be `""` thus matches
	// the init exec of the task matching `container_id`.
	ID string `protobuf:"bytes,5,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TaskDelete) Reset() {
	*x = TaskDelete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_containerd_containerd_api_events_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskDelete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskDelete) ProtoMessage() {}

func (x *TaskDelete) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_containerd_containerd_api_events_task_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskDelete.ProtoReflect.Descriptor instead.
func (*TaskDelete) Descriptor() ([]byte, []int) {
	return file_github_com_containerd_containerd_api_events_task_proto_rawDescGZIP(), []int{2}
}

func (x *TaskDelete) GetContainerID() string {
	if x != nil {
		return x.ContainerID
	}
	return ""
}

func (x *TaskDelete) GetPid() uint32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *TaskDelete) GetExitStatus() uint32 {
	if x != nil {
		return x.ExitStatus
	}
	return 0
}

func (x *TaskDelete) GetExitedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExitedAt
	}
	return nil
}

func (x *TaskDelete) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type TaskIO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stdin    string `protobuf:"bytes,1,opt,name=stdin,proto3" json:"stdin,omitempty"`
	Stdout   string `protobuf:"bytes,2,opt,name=stdout,proto3" json:"stdout,omitempty"`
	Stderr   string `protobuf:"bytes,3,opt,name=stderr,proto3" json:"stderr,omitempty"`
	Terminal bool   `protobuf:"varint,4,opt,name=terminal,proto3" json:"terminal,omitempty"`
}

func (x *TaskIO) Reset() {
	*x = TaskIO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_containerd_containerd_api_events_task_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskIO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskIO) ProtoMessage() {}

func (x *TaskIO) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_containerd_containerd_api_events_task_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskIO.ProtoReflect.Descriptor instead.
func (*TaskIO) Descriptor() ([]byte, []int) {
	return file_github_com_containerd_containerd_api_events_task_proto_rawDescGZIP(), []int{3}
}

func (x *TaskIO) GetStdin() string {
	if x != nil {
		return x.Stdin
	}
	return ""
}

func (x *TaskIO) GetStdout() string {
	if x != nil {
		return x.Stdout
	}
	return ""
}

func (x *TaskIO) GetStderr() string {
	if x != nil {
		return x.Stderr
	}
	return ""
}

func (x *TaskIO) GetTerminal() bool {
	if x != nil {
		return x.Terminal
	}
	return false
}

type TaskExit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContainerID string                 `protobuf:"bytes,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	ID          string                 `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Pid         uint32                 `protobuf:"varint,3,opt,name=pid,proto3" json:"pid,omitempty"`
	ExitStatus  uint32                 `protobuf:"varint,4,opt,name=exit_status,json=exitStatus,proto3" json:"exit_status,omitempty"`
	ExitedAt    *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=exited_at,json=exitedAt,proto3" json:"exited_at,omitempty"`
}

func (x *TaskExit) Reset() {
	*x = TaskExit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_containerd_containerd_api_events_task_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskExit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskExit) ProtoMessage() {}

func (x *TaskExit) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_containerd_containerd_api_events_task_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskExit.ProtoReflect.Descriptor instead.
func (*TaskExit) Descriptor() ([]byte, []int) {
	return file_github_com_containerd_containerd_api_events_task_proto_rawDescGZIP(), []int{4}
}

func (x *TaskExit) GetContainerID() string {
	if x != nil {
		return x.ContainerID
	}
	return ""
}

func (x *TaskExit) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *TaskExit) GetPid() uint32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *TaskExit) GetExitStatus() uint32 {
	if x != nil {
		return x.ExitStatus
	}
	return 0
}

func (x *TaskExit) GetExitedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExitedAt
	}
	return nil
}

type TaskOOM struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContainerID string `protobuf:"bytes,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
}

func (x *TaskOOM) Reset() {
	*x = TaskOOM{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_containerd_containerd_api_events_task_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskOOM) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskOOM) ProtoMessage() {}

func (x *TaskOOM) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_containerd_containerd_api_events_task_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskOOM.ProtoReflect.Descriptor instead.
func (*TaskOOM) Descriptor() ([]byte, []int) {
	return file_github_com_containerd_containerd_api_events_task_proto_rawDescGZIP(), []int{5}
}

func (x *TaskOOM) GetContainerID() string {
	if x != nil {
		return x.ContainerID
	}
	return ""
}

type TaskExecAdded struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContainerID string `protobuf:"bytes,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	ExecID      string `protobuf:"bytes,2,opt,name=exec_id,json=execId,proto3" json:"exec_id,omitempty"`
}

func (x *TaskExecAdded) Reset() {
	*x = TaskExecAdded{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_containerd_containerd_api_events_task_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskExecAdded) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskExecAdded) ProtoMessage() {}

func (x *TaskExecAdded) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_containerd_containerd_api_events_task_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskExecAdded.ProtoReflect.Descriptor instead.
func (*TaskExecAdded) Descriptor() ([]byte, []int) {
	return file_github_com_containerd_containerd_api_events_task_proto_rawDescGZIP(), []int{6}
}

func (x *TaskExecAdded) GetContainerID() string {
	if x != nil {
		return x.ContainerID
	}
	return ""
}

func (x *TaskExecAdded) GetExecID() string {
	if x != nil {
		return x.ExecID
	}
	return ""
}

type TaskExecStarted struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContainerID string `protobuf:"bytes,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	ExecID      string `protobuf:"bytes,2,opt,name=exec_id,json=execId,proto3" json:"exec_id,omitempty"`
	Pid         uint32 `protobuf:"varint,3,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (x *TaskExecStarted) Reset() {
	*x = TaskExecStarted{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_containerd_containerd_api_events_task_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskExecStarted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskExecStarted) ProtoMessage() {}

func (x *TaskExecStarted) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_containerd_containerd_api_events_task_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskExecStarted.ProtoReflect.Descriptor instead.
func (*TaskExecStarted) Descriptor() ([]byte, []int) {
	return file_github_com_containerd_containerd_api_events_task_proto_rawDescGZIP(), []int{7}
}

func (x *TaskExecStarted) GetContainerID() string {
	if x != nil {
		return x.ContainerID
	}
	return ""
}

func (x *TaskExecStarted) GetExecID() string {
	if x != nil {
		return x.ExecID
	}
	return ""
}

func (x *TaskExecStarted) GetPid() uint32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

type TaskPaused struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContainerID string `protobuf:"bytes,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
}

func (x *TaskPaused) Reset() {
	*x = TaskPaused{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_containerd_containerd_api_events_task_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskPaused) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskPaused) ProtoMessage() {}

func (x *TaskPaused) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_containerd_containerd_api_events_task_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskPaused.ProtoReflect.Descriptor instead.
func (*TaskPaused) Descriptor() ([]byte, []int) {
	return file_github_com_containerd_containerd_api_events_task_proto_rawDescGZIP(), []int{8}
}

func (x *TaskPaused) GetContainerID() string {
	if x != nil {
		return x.ContainerID
	}
	return ""
}

type TaskResumed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContainerID string `protobuf:"bytes,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
}

func (x *TaskResumed) Reset() {
	*x = TaskResumed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_containerd_containerd_api_events_task_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskResumed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskResumed) ProtoMessage() {}

func (x *TaskResumed) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_containerd_containerd_api_events_task_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskResumed.ProtoReflect.Descriptor instead.
func (*TaskResumed) Descriptor() ([]byte, []int) {
	return file_github_com_containerd_containerd_api_events_task_proto_rawDescGZIP(), []int{9}
}

func (x *TaskResumed) GetContainerID() string {
	if x != nil {
		return x.ContainerID
	}
	return ""
}

type TaskCheckpointed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContainerID string `protobuf:"bytes,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	Checkpoint  string `protobuf:"bytes,2,opt,name=checkpoint,proto3" json:"checkpoint,omitempty"`
}

func (x *TaskCheckpointed) Reset() {
	*x = TaskCheckpointed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_containerd_containerd_api_events_task_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskCheckpointed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskCheckpointed) ProtoMessage() {}

func (x *TaskCheckpointed) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_containerd_containerd_api_events_task_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskCheckpointed.ProtoReflect.Descriptor instead.
func (*TaskCheckpointed) Descriptor() ([]byte, []int) {
	return file_github_com_containerd_containerd_api_events_task_proto_rawDescGZIP(), []int{10}
}

func (x *TaskCheckpointed) GetContainerID() string {
	if x != nil {
		return x.ContainerID
	}
	return ""
}

func (x *TaskCheckpointed) GetCheckpoint() string {
	if x != nil {
		return x.Checkpoint
	}
	return ""
}

var File_github_com_containerd_containerd_api_events_task_proto protoreflect.FileDescriptor

var file_github_com_containerd_containerd_api_events_task_proto_rawDesc = []byte{
	0x0a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x74, 0x61,
	0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x64, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x64, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x2f, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x70, 0x61, 0x74, 0x68,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd5, 0x01, 0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x6e, 0x64,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65,
	0x12, 0x2f, 0x0a, 0x06, 0x72, 0x6f, 0x6f, 0x74, 0x66, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x74, 0x66,
	0x73, 0x12, 0x29, 0x0a, 0x02, 0x69, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x4f, 0x52, 0x02, 0x69, 0x6f, 0x12, 0x1e, 0x0a, 0x0a,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x70, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x70, 0x69, 0x64, 0x22, 0x40,
	0x0a, 0x09, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x70, 0x69, 0x64,
	0x22, 0xab, 0x01, 0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x03, 0x70, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x65, 0x78, 0x69, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x37, 0x0a, 0x09, 0x65, 0x78, 0x69, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x65, 0x78, 0x69, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x6a,
	0x0a, 0x06, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x4f, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x64, 0x69,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x64, 0x69, 0x6e, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x64, 0x65, 0x72, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x64, 0x65, 0x72, 0x72, 0x12, 0x1a,
	0x0a, 0x08, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x22, 0xa9, 0x01, 0x0a, 0x08, 0x54,
	0x61, 0x73, 0x6b, 0x45, 0x78, 0x69, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x65, 0x78, 0x69, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x65, 0x78, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x37, 0x0a,
	0x09, 0x65, 0x78, 0x69, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x65, 0x78,
	0x69, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x2c, 0x0a, 0x07, 0x54, 0x61, 0x73, 0x6b, 0x4f, 0x4f,
	0x4d, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x4b, 0x0a, 0x0d, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x78, 0x65, 0x63,
	0x41, 0x64, 0x64, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x65, 0x78, 0x65, 0x63,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x78, 0x65, 0x63, 0x49,
	0x64, 0x22, 0x5f, 0x0a, 0x0f, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x78, 0x65, 0x63, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x65, 0x78, 0x65, 0x63, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x78, 0x65, 0x63, 0x49, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x70,
	0x69, 0x64, 0x22, 0x2f, 0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x50, 0x61, 0x75, 0x73, 0x65, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x30, 0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6d,
	0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x49, 0x64, 0x22, 0x55, 0x0a, 0x10, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x42, 0x38, 0x5a, 0x32,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x64, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x3b, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0xa0, 0xf4, 0x1e, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_containerd_containerd_api_events_task_proto_rawDescOnce sync.Once
	file_github_com_containerd_containerd_api_events_task_proto_rawDescData = file_github_com_containerd_containerd_api_events_task_proto_rawDesc
)

func file_github_com_containerd_containerd_api_events_task_proto_rawDescGZIP() []byte {
	file_github_com_containerd_containerd_api_events_task_proto_rawDescOnce.Do(func() {
		file_github_com_containerd_containerd_api_events_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_containerd_containerd_api_events_task_proto_rawDescData)
	})
	return file_github_com_containerd_containerd_api_events_task_proto_rawDescData
}

var file_github_com_containerd_containerd_api_events_task_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_github_com_containerd_containerd_api_events_task_proto_goTypes = []interface{}{
	(*TaskCreate)(nil),            // 0: containerd.events.TaskCreate
	(*TaskStart)(nil),             // 1: containerd.events.TaskStart
	(*TaskDelete)(nil),            // 2: containerd.events.TaskDelete
	(*TaskIO)(nil),                // 3: containerd.events.TaskIO
	(*TaskExit)(nil),              // 4: containerd.events.TaskExit
	(*TaskOOM)(nil),               // 5: containerd.events.TaskOOM
	(*TaskExecAdded)(nil),         // 6: containerd.events.TaskExecAdded
	(*TaskExecStarted)(nil),       // 7: containerd.events.TaskExecStarted
	(*TaskPaused)(nil),            // 8: containerd.events.TaskPaused
	(*TaskResumed)(nil),           // 9: containerd.events.TaskResumed
	(*TaskCheckpointed)(nil),      // 10: containerd.events.TaskCheckpointed
	(*types.Mount)(nil),           // 11: containerd.types.Mount
	(*timestamppb.Timestamp)(nil), // 12: google.protobuf.Timestamp
}
var file_github_com_containerd_containerd_api_events_task_proto_depIdxs = []int32{
	11, // 0: containerd.events.TaskCreate.rootfs:type_name -> containerd.types.Mount
	3,  // 1: containerd.events.TaskCreate.io:type_name -> containerd.events.TaskIO
	12, // 2: containerd.events.TaskDelete.exited_at:type_name -> google.protobuf.Timestamp
	12, // 3: containerd.events.TaskExit.exited_at:type_name -> google.protobuf.Timestamp
	4,  // [4:4] is the sub-list for method output_type
	4,  // [4:4] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_github_com_containerd_containerd_api_events_task_proto_init() }
func file_github_com_containerd_containerd_api_events_task_proto_init() {
	if File_github_com_containerd_containerd_api_events_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_containerd_containerd_api_events_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskCreate); i {
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
		file_github_com_containerd_containerd_api_events_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskStart); i {
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
		file_github_com_containerd_containerd_api_events_task_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskDelete); i {
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
		file_github_com_containerd_containerd_api_events_task_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskIO); i {
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
		file_github_com_containerd_containerd_api_events_task_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskExit); i {
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
		file_github_com_containerd_containerd_api_events_task_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskOOM); i {
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
		file_github_com_containerd_containerd_api_events_task_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskExecAdded); i {
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
		file_github_com_containerd_containerd_api_events_task_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskExecStarted); i {
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
		file_github_com_containerd_containerd_api_events_task_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskPaused); i {
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
		file_github_com_containerd_containerd_api_events_task_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskResumed); i {
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
		file_github_com_containerd_containerd_api_events_task_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskCheckpointed); i {
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
			RawDescriptor: file_github_com_containerd_containerd_api_events_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_containerd_containerd_api_events_task_proto_goTypes,
		DependencyIndexes: file_github_com_containerd_containerd_api_events_task_proto_depIdxs,
		MessageInfos:      file_github_com_containerd_containerd_api_events_task_proto_msgTypes,
	}.Build()
	File_github_com_containerd_containerd_api_events_task_proto = out.File
	file_github_com_containerd_containerd_api_events_task_proto_rawDesc = nil
	file_github_com_containerd_containerd_api_events_task_proto_goTypes = nil
	file_github_com_containerd_containerd_api_events_task_proto_depIdxs = nil
}
