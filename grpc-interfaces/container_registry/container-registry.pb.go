// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.24.4
// source: container-registry.proto

package container_registry

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

type CreateReadOnlyCredentialIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountName      string `protobuf:"bytes,1,opt,name=accountName,proto3" json:"accountName,omitempty"`
	UserId           string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	CredentialName   string `protobuf:"bytes,3,opt,name=credentialName,proto3" json:"credentialName,omitempty"`
	RegistryUsername string `protobuf:"bytes,4,opt,name=registryUsername,proto3" json:"registryUsername,omitempty"`
}

func (x *CreateReadOnlyCredentialIn) Reset() {
	*x = CreateReadOnlyCredentialIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_container_registry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateReadOnlyCredentialIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReadOnlyCredentialIn) ProtoMessage() {}

func (x *CreateReadOnlyCredentialIn) ProtoReflect() protoreflect.Message {
	mi := &file_container_registry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReadOnlyCredentialIn.ProtoReflect.Descriptor instead.
func (*CreateReadOnlyCredentialIn) Descriptor() ([]byte, []int) {
	return file_container_registry_proto_rawDescGZIP(), []int{0}
}

func (x *CreateReadOnlyCredentialIn) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *CreateReadOnlyCredentialIn) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateReadOnlyCredentialIn) GetCredentialName() string {
	if x != nil {
		return x.CredentialName
	}
	return ""
}

func (x *CreateReadOnlyCredentialIn) GetRegistryUsername() string {
	if x != nil {
		return x.RegistryUsername
	}
	return ""
}

type CreateReadOnlyCredentialOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// dcokerconfigjson is as per format: https://kubernetes.io/docs/concepts/configuration/secret/#docker-config-secrets
	DockerConfigJson []byte `protobuf:"bytes,1,opt,name=dockerConfigJson,proto3" json:"dockerConfigJson,omitempty"`
}

func (x *CreateReadOnlyCredentialOut) Reset() {
	*x = CreateReadOnlyCredentialOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_container_registry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateReadOnlyCredentialOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReadOnlyCredentialOut) ProtoMessage() {}

func (x *CreateReadOnlyCredentialOut) ProtoReflect() protoreflect.Message {
	mi := &file_container_registry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReadOnlyCredentialOut.ProtoReflect.Descriptor instead.
func (*CreateReadOnlyCredentialOut) Descriptor() ([]byte, []int) {
	return file_container_registry_proto_rawDescGZIP(), []int{1}
}

func (x *CreateReadOnlyCredentialOut) GetDockerConfigJson() []byte {
	if x != nil {
		return x.DockerConfigJson
	}
	return nil
}

var File_container_registry_proto protoreflect.FileDescriptor

var file_container_registry_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2d, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x01, 0x0a, 0x1a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x61, 0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x43, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x49, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x55,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x49, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x61, 0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x4f, 0x75, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4a, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x10, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4a, 0x73,
	0x6f, 0x6e, 0x32, 0x6a, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x55, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x61, 0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x12, 0x1b, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x61, 0x64,
	0x4f, 0x6e, 0x6c, 0x79, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x49, 0x6e,
	0x1a, 0x1c, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x61, 0x64, 0x4f, 0x6e, 0x6c,
	0x79, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x4f, 0x75, 0x74, 0x42, 0x16,
	0x5a, 0x14, 0x2e, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_container_registry_proto_rawDescOnce sync.Once
	file_container_registry_proto_rawDescData = file_container_registry_proto_rawDesc
)

func file_container_registry_proto_rawDescGZIP() []byte {
	file_container_registry_proto_rawDescOnce.Do(func() {
		file_container_registry_proto_rawDescData = protoimpl.X.CompressGZIP(file_container_registry_proto_rawDescData)
	})
	return file_container_registry_proto_rawDescData
}

var file_container_registry_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_container_registry_proto_goTypes = []interface{}{
	(*CreateReadOnlyCredentialIn)(nil),  // 0: CreateReadOnlyCredentialIn
	(*CreateReadOnlyCredentialOut)(nil), // 1: CreateReadOnlyCredentialOut
}
var file_container_registry_proto_depIdxs = []int32{
	0, // 0: ContainerRegistry.CreateReadOnlyCredential:input_type -> CreateReadOnlyCredentialIn
	1, // 1: ContainerRegistry.CreateReadOnlyCredential:output_type -> CreateReadOnlyCredentialOut
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_container_registry_proto_init() }
func file_container_registry_proto_init() {
	if File_container_registry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_container_registry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateReadOnlyCredentialIn); i {
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
		file_container_registry_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateReadOnlyCredentialOut); i {
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
			RawDescriptor: file_container_registry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_container_registry_proto_goTypes,
		DependencyIndexes: file_container_registry_proto_depIdxs,
		MessageInfos:      file_container_registry_proto_msgTypes,
	}.Build()
	File_container_registry_proto = out.File
	file_container_registry_proto_rawDesc = nil
	file_container_registry_proto_goTypes = nil
	file_container_registry_proto_depIdxs = nil
}
