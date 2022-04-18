// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: iam.proto

package iam

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

type RoleBinding struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	ResourceId   string `protobuf:"bytes,2,opt,name=resourceId,proto3" json:"resourceId,omitempty"`
	ResourceType string `protobuf:"bytes,3,opt,name=resourceType,proto3" json:"resourceType,omitempty"`
	Role         string `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *RoleBinding) Reset() {
	*x = RoleBinding{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iam_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleBinding) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleBinding) ProtoMessage() {}

func (x *RoleBinding) ProtoReflect() protoreflect.Message {
	mi := &file_iam_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleBinding.ProtoReflect.Descriptor instead.
func (*RoleBinding) Descriptor() ([]byte, []int) {
	return file_iam_proto_rawDescGZIP(), []int{0}
}

func (x *RoleBinding) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RoleBinding) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *RoleBinding) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *RoleBinding) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type InAddMembership struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	ResourceType string `protobuf:"bytes,2,opt,name=resourceType,proto3" json:"resourceType,omitempty"`
	ResourceId   string `protobuf:"bytes,3,opt,name=resourceId,proto3" json:"resourceId,omitempty"`
	Role         string `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`
	Filter       string `protobuf:"bytes,5,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *InAddMembership) Reset() {
	*x = InAddMembership{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iam_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InAddMembership) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InAddMembership) ProtoMessage() {}

func (x *InAddMembership) ProtoReflect() protoreflect.Message {
	mi := &file_iam_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InAddMembership.ProtoReflect.Descriptor instead.
func (*InAddMembership) Descriptor() ([]byte, []int) {
	return file_iam_proto_rawDescGZIP(), []int{1}
}

func (x *InAddMembership) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *InAddMembership) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *InAddMembership) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *InAddMembership) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *InAddMembership) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

type OutAddMembership struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *OutAddMembership) Reset() {
	*x = OutAddMembership{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iam_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutAddMembership) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutAddMembership) ProtoMessage() {}

func (x *OutAddMembership) ProtoReflect() protoreflect.Message {
	mi := &file_iam_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutAddMembership.ProtoReflect.Descriptor instead.
func (*OutAddMembership) Descriptor() ([]byte, []int) {
	return file_iam_proto_rawDescGZIP(), []int{2}
}

func (x *OutAddMembership) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

type InRemoveMembership struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	ResourceId string `protobuf:"bytes,2,opt,name=resourceId,proto3" json:"resourceId,omitempty"`
}

func (x *InRemoveMembership) Reset() {
	*x = InRemoveMembership{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iam_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InRemoveMembership) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InRemoveMembership) ProtoMessage() {}

func (x *InRemoveMembership) ProtoReflect() protoreflect.Message {
	mi := &file_iam_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InRemoveMembership.ProtoReflect.Descriptor instead.
func (*InRemoveMembership) Descriptor() ([]byte, []int) {
	return file_iam_proto_rawDescGZIP(), []int{3}
}

func (x *InRemoveMembership) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *InRemoveMembership) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

type OutRemoveMembership struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *OutRemoveMembership) Reset() {
	*x = OutRemoveMembership{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iam_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutRemoveMembership) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutRemoveMembership) ProtoMessage() {}

func (x *OutRemoveMembership) ProtoReflect() protoreflect.Message {
	mi := &file_iam_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutRemoveMembership.ProtoReflect.Descriptor instead.
func (*OutRemoveMembership) Descriptor() ([]byte, []int) {
	return file_iam_proto_rawDescGZIP(), []int{4}
}

func (x *OutRemoveMembership) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

type InRemoveResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId string `protobuf:"bytes,1,opt,name=resourceId,proto3" json:"resourceId,omitempty"`
}

func (x *InRemoveResource) Reset() {
	*x = InRemoveResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iam_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InRemoveResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InRemoveResource) ProtoMessage() {}

func (x *InRemoveResource) ProtoReflect() protoreflect.Message {
	mi := &file_iam_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InRemoveResource.ProtoReflect.Descriptor instead.
func (*InRemoveResource) Descriptor() ([]byte, []int) {
	return file_iam_proto_rawDescGZIP(), []int{5}
}

func (x *InRemoveResource) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

type OutRemoveResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *OutRemoveResource) Reset() {
	*x = OutRemoveResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iam_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutRemoveResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutRemoveResource) ProtoMessage() {}

func (x *OutRemoveResource) ProtoReflect() protoreflect.Message {
	mi := &file_iam_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutRemoveResource.ProtoReflect.Descriptor instead.
func (*OutRemoveResource) Descriptor() ([]byte, []int) {
	return file_iam_proto_rawDescGZIP(), []int{6}
}

func (x *OutRemoveResource) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

type OutListMemberships struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleBindings []*RoleBinding `protobuf:"bytes,1,rep,name=roleBindings,proto3" json:"roleBindings,omitempty"`
}

func (x *OutListMemberships) Reset() {
	*x = OutListMemberships{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iam_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutListMemberships) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutListMemberships) ProtoMessage() {}

func (x *OutListMemberships) ProtoReflect() protoreflect.Message {
	mi := &file_iam_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutListMemberships.ProtoReflect.Descriptor instead.
func (*OutListMemberships) Descriptor() ([]byte, []int) {
	return file_iam_proto_rawDescGZIP(), []int{7}
}

func (x *OutListMemberships) GetRoleBindings() []*RoleBinding {
	if x != nil {
		return x.RoleBindings
	}
	return nil
}

type InUserMemberships struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *InUserMemberships) Reset() {
	*x = InUserMemberships{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iam_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InUserMemberships) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InUserMemberships) ProtoMessage() {}

func (x *InUserMemberships) ProtoReflect() protoreflect.Message {
	mi := &file_iam_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InUserMemberships.ProtoReflect.Descriptor instead.
func (*InUserMemberships) Descriptor() ([]byte, []int) {
	return file_iam_proto_rawDescGZIP(), []int{8}
}

func (x *InUserMemberships) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type InResourceMemberships struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceType string `protobuf:"bytes,1,opt,name=resourceType,proto3" json:"resourceType,omitempty"`
	ResourceId   string `protobuf:"bytes,2,opt,name=resourceId,proto3" json:"resourceId,omitempty"`
}

func (x *InResourceMemberships) Reset() {
	*x = InResourceMemberships{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iam_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InResourceMemberships) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InResourceMemberships) ProtoMessage() {}

func (x *InResourceMemberships) ProtoReflect() protoreflect.Message {
	mi := &file_iam_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InResourceMemberships.ProtoReflect.Descriptor instead.
func (*InResourceMemberships) Descriptor() ([]byte, []int) {
	return file_iam_proto_rawDescGZIP(), []int{9}
}

func (x *InResourceMemberships) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *InResourceMemberships) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

type OutResourceMemberships struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceType string `protobuf:"bytes,1,opt,name=resourceType,proto3" json:"resourceType,omitempty"`
	ResourceId   string `protobuf:"bytes,2,opt,name=resourceId,proto3" json:"resourceId,omitempty"`
}

func (x *OutResourceMemberships) Reset() {
	*x = OutResourceMemberships{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iam_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutResourceMemberships) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutResourceMemberships) ProtoMessage() {}

func (x *OutResourceMemberships) ProtoReflect() protoreflect.Message {
	mi := &file_iam_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutResourceMemberships.ProtoReflect.Descriptor instead.
func (*OutResourceMemberships) Descriptor() ([]byte, []int) {
	return file_iam_proto_rawDescGZIP(), []int{10}
}

func (x *OutResourceMemberships) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *OutResourceMemberships) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iam_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_iam_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_iam_proto_rawDescGZIP(), []int{11}
}

func (x *Message) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type InCan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	ResourceIds []string `protobuf:"bytes,2,rep,name=resourceIds,proto3" json:"resourceIds,omitempty"`
	Action      string   `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
}

func (x *InCan) Reset() {
	*x = InCan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iam_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InCan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InCan) ProtoMessage() {}

func (x *InCan) ProtoReflect() protoreflect.Message {
	mi := &file_iam_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InCan.ProtoReflect.Descriptor instead.
func (*InCan) Descriptor() ([]byte, []int) {
	return file_iam_proto_rawDescGZIP(), []int{12}
}

func (x *InCan) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *InCan) GetResourceIds() []string {
	if x != nil {
		return x.ResourceIds
	}
	return nil
}

func (x *InCan) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

type OutCan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *OutCan) Reset() {
	*x = OutCan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iam_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutCan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutCan) ProtoMessage() {}

func (x *OutCan) ProtoReflect() protoreflect.Message {
	mi := &file_iam_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutCan.ProtoReflect.Descriptor instead.
func (*OutCan) Descriptor() ([]byte, []int) {
	return file_iam_proto_rawDescGZIP(), []int{13}
}

func (x *OutCan) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

var File_iam_proto protoreflect.FileDescriptor

var file_iam_proto_rawDesc = []byte{
	0x0a, 0x09, 0x69, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7d, 0x0a, 0x0b, 0x52,
	0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x99, 0x01, 0x0a, 0x0f, 0x49,
	0x6e, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x2a, 0x0a, 0x10, 0x4f, 0x75, 0x74, 0x41, 0x64, 0x64,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0x4c, 0x0a, 0x12, 0x49, 0x6e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64,
	0x22, 0x2d, 0x0a, 0x13, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x32, 0x0a, 0x10, 0x49, 0x6e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x49, 0x64, 0x22, 0x2b, 0x0a, 0x11, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x22, 0x46, 0x0a, 0x12, 0x4f, 0x75, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x12, 0x30, 0x0a, 0x0c, 0x72, 0x6f, 0x6c, 0x65, 0x42, 0x69,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x52,
	0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x0c, 0x72, 0x6f, 0x6c, 0x65,
	0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x2b, 0x0a, 0x11, 0x49, 0x6e, 0x55, 0x73,
	0x65, 0x72, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x5b, 0x0a, 0x15, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x12, 0x22,
	0x0a, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x49, 0x64, 0x22, 0x5c, 0x0a, 0x16, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x12, 0x22, 0x0a, 0x0c,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64,
	0x22, 0x23, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x59, 0x0a, 0x05, 0x49, 0x6e, 0x43, 0x61, 0x6e, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x49, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x20, 0x0a, 0x06, 0x4f, 0x75, 0x74, 0x43, 0x61, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x32, 0xef, 0x02, 0x0a, 0x03, 0x49, 0x41, 0x4d, 0x12, 0x1a, 0x0a, 0x04, 0x50, 0x69,
	0x6e, 0x67, 0x12, 0x08, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x08, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x03, 0x43, 0x61, 0x6e, 0x12, 0x06, 0x2e,
	0x49, 0x6e, 0x43, 0x61, 0x6e, 0x1a, 0x07, 0x2e, 0x4f, 0x75, 0x74, 0x43, 0x61, 0x6e, 0x12, 0x3e,
	0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x68, 0x69, 0x70, 0x73, 0x12, 0x12, 0x2e, 0x49, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x1a, 0x13, 0x2e, 0x4f, 0x75, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x12, 0x46,
	0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x12, 0x16, 0x2e, 0x49, 0x6e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70,
	0x73, 0x1a, 0x13, 0x2e, 0x4f, 0x75, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x12, 0x34, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x12, 0x10, 0x2e, 0x49, 0x6e, 0x41, 0x64, 0x64, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x1a, 0x11, 0x2e, 0x4f, 0x75, 0x74, 0x41,
	0x64, 0x64, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x12, 0x3d, 0x0a, 0x10,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70,
	0x12, 0x13, 0x2e, 0x49, 0x6e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x68, 0x69, 0x70, 0x1a, 0x14, 0x2e, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x12, 0x37, 0x0a, 0x0e, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x11, 0x2e,
	0x49, 0x6e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x1a, 0x12, 0x2e, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x42, 0x16, 0x5a, 0x14, 0x6b, 0x6c, 0x6f, 0x75, 0x64, 0x6c, 0x69, 0x74,
	0x65, 0x2e, 0x69, 0x6f, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x69, 0x61, 0x6d, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_iam_proto_rawDescOnce sync.Once
	file_iam_proto_rawDescData = file_iam_proto_rawDesc
)

func file_iam_proto_rawDescGZIP() []byte {
	file_iam_proto_rawDescOnce.Do(func() {
		file_iam_proto_rawDescData = protoimpl.X.CompressGZIP(file_iam_proto_rawDescData)
	})
	return file_iam_proto_rawDescData
}

var file_iam_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_iam_proto_goTypes = []interface{}{
	(*RoleBinding)(nil),            // 0: RoleBinding
	(*InAddMembership)(nil),        // 1: InAddMembership
	(*OutAddMembership)(nil),       // 2: OutAddMembership
	(*InRemoveMembership)(nil),     // 3: InRemoveMembership
	(*OutRemoveMembership)(nil),    // 4: OutRemoveMembership
	(*InRemoveResource)(nil),       // 5: InRemoveResource
	(*OutRemoveResource)(nil),      // 6: OutRemoveResource
	(*OutListMemberships)(nil),     // 7: OutListMemberships
	(*InUserMemberships)(nil),      // 8: InUserMemberships
	(*InResourceMemberships)(nil),  // 9: InResourceMemberships
	(*OutResourceMemberships)(nil), // 10: OutResourceMemberships
	(*Message)(nil),                // 11: Message
	(*InCan)(nil),                  // 12: InCan
	(*OutCan)(nil),                 // 13: OutCan
}
var file_iam_proto_depIdxs = []int32{
	0,  // 0: OutListMemberships.roleBindings:type_name -> RoleBinding
	11, // 1: IAM.Ping:input_type -> Message
	12, // 2: IAM.Can:input_type -> InCan
	8,  // 3: IAM.ListUserMemberships:input_type -> InUserMemberships
	9,  // 4: IAM.ListResourceMemberships:input_type -> InResourceMemberships
	1,  // 5: IAM.AddMembership:input_type -> InAddMembership
	3,  // 6: IAM.RemoveMembership:input_type -> InRemoveMembership
	5,  // 7: IAM.RemoveResource:input_type -> InRemoveResource
	11, // 8: IAM.Ping:output_type -> Message
	13, // 9: IAM.Can:output_type -> OutCan
	7,  // 10: IAM.ListUserMemberships:output_type -> OutListMemberships
	7,  // 11: IAM.ListResourceMemberships:output_type -> OutListMemberships
	2,  // 12: IAM.AddMembership:output_type -> OutAddMembership
	4,  // 13: IAM.RemoveMembership:output_type -> OutRemoveMembership
	6,  // 14: IAM.RemoveResource:output_type -> OutRemoveResource
	8,  // [8:15] is the sub-list for method output_type
	1,  // [1:8] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_iam_proto_init() }
func file_iam_proto_init() {
	if File_iam_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_iam_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleBinding); i {
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
		file_iam_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InAddMembership); i {
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
		file_iam_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutAddMembership); i {
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
		file_iam_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InRemoveMembership); i {
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
		file_iam_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutRemoveMembership); i {
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
		file_iam_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InRemoveResource); i {
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
		file_iam_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutRemoveResource); i {
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
		file_iam_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutListMemberships); i {
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
		file_iam_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InUserMemberships); i {
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
		file_iam_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InResourceMemberships); i {
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
		file_iam_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutResourceMemberships); i {
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
		file_iam_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_iam_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InCan); i {
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
		file_iam_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutCan); i {
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
			RawDescriptor: file_iam_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_iam_proto_goTypes,
		DependencyIndexes: file_iam_proto_depIdxs,
		MessageInfos:      file_iam_proto_msgTypes,
	}.Build()
	File_iam_proto = out.File
	file_iam_proto_rawDesc = nil
	file_iam_proto_goTypes = nil
	file_iam_proto_depIdxs = nil
}
