// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: kubeagent.proto

package agent

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PayloadIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action      string                `protobuf:"bytes,1,opt,name=Action,proto3" json:"Action,omitempty"`
	Payload     map[string]*anypb.Any `protobuf:"bytes,2,rep,name=payload,proto3" json:"payload,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	AccountId   string                `protobuf:"bytes,3,opt,name=accountId,proto3" json:"accountId,omitempty"`
	ResourceRef string                `protobuf:"bytes,4,opt,name=ResourceRef,proto3" json:"ResourceRef,omitempty"`
}

func (x *PayloadIn) Reset() {
	*x = PayloadIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kubeagent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayloadIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadIn) ProtoMessage() {}

func (x *PayloadIn) ProtoReflect() protoreflect.Message {
	mi := &file_kubeagent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadIn.ProtoReflect.Descriptor instead.
func (*PayloadIn) Descriptor() ([]byte, []int) {
	return file_kubeagent_proto_rawDescGZIP(), []int{0}
}

func (x *PayloadIn) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *PayloadIn) GetPayload() map[string]*anypb.Any {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *PayloadIn) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *PayloadIn) GetResourceRef() string {
	if x != nil {
		return x.ResourceRef
	}
	return ""
}

type PayloadOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Stdout  string `protobuf:"bytes,2,opt,name=stdout,proto3" json:"stdout,omitempty"`
	Stderr  string `protobuf:"bytes,3,opt,name=stderr,proto3" json:"stderr,omitempty"`
	ExecErr string `protobuf:"bytes,4,opt,name=execErr,proto3" json:"execErr,omitempty"`
}

func (x *PayloadOut) Reset() {
	*x = PayloadOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kubeagent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayloadOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadOut) ProtoMessage() {}

func (x *PayloadOut) ProtoReflect() protoreflect.Message {
	mi := &file_kubeagent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadOut.ProtoReflect.Descriptor instead.
func (*PayloadOut) Descriptor() ([]byte, []int) {
	return file_kubeagent_proto_rawDescGZIP(), []int{1}
}

func (x *PayloadOut) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *PayloadOut) GetStdout() string {
	if x != nil {
		return x.Stdout
	}
	return ""
}

func (x *PayloadOut) GetStderr() string {
	if x != nil {
		return x.Stderr
	}
	return ""
}

func (x *PayloadOut) GetExecErr() string {
	if x != nil {
		return x.ExecErr
	}
	return ""
}

var File_kubeagent_proto protoreflect.FileDescriptor

var file_kubeagent_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6b, 0x75, 0x62, 0x65, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe8, 0x01, 0x0a,
	0x09, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x2e,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x65, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x52, 0x65, 0x66, 0x1a, 0x50, 0x0a, 0x0c, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x70, 0x0a, 0x0a, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x4f, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x64, 0x65, 0x72,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x64, 0x65, 0x72, 0x72, 0x12,
	0x18, 0x0a, 0x07, 0x65, 0x78, 0x65, 0x63, 0x45, 0x72, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x65, 0x78, 0x65, 0x63, 0x45, 0x72, 0x72, 0x32, 0x31, 0x0a, 0x09, 0x4b, 0x75, 0x62,
	0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x09, 0x4b, 0x75, 0x62, 0x65, 0x41, 0x70,
	0x70, 0x6c, 0x79, 0x12, 0x0a, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x1a,
	0x0b, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x4f, 0x75, 0x74, 0x42, 0x18, 0x5a, 0x16,
	0x6b, 0x6c, 0x6f, 0x75, 0x64, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x69, 0x6f, 0x2f, 0x72, 0x70, 0x63,
	0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kubeagent_proto_rawDescOnce sync.Once
	file_kubeagent_proto_rawDescData = file_kubeagent_proto_rawDesc
)

func file_kubeagent_proto_rawDescGZIP() []byte {
	file_kubeagent_proto_rawDescOnce.Do(func() {
		file_kubeagent_proto_rawDescData = protoimpl.X.CompressGZIP(file_kubeagent_proto_rawDescData)
	})
	return file_kubeagent_proto_rawDescData
}

var file_kubeagent_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_kubeagent_proto_goTypes = []interface{}{
	(*PayloadIn)(nil),  // 0: PayloadIn
	(*PayloadOut)(nil), // 1: PayloadOut
	nil,                // 2: PayloadIn.PayloadEntry
	(*anypb.Any)(nil),  // 3: google.protobuf.Any
}
var file_kubeagent_proto_depIdxs = []int32{
	2, // 0: PayloadIn.payload:type_name -> PayloadIn.PayloadEntry
	3, // 1: PayloadIn.PayloadEntry.value:type_name -> google.protobuf.Any
	0, // 2: KubeAgent.KubeApply:input_type -> PayloadIn
	1, // 3: KubeAgent.KubeApply:output_type -> PayloadOut
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_kubeagent_proto_init() }
func file_kubeagent_proto_init() {
	if File_kubeagent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kubeagent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayloadIn); i {
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
		file_kubeagent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayloadOut); i {
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
			RawDescriptor: file_kubeagent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kubeagent_proto_goTypes,
		DependencyIndexes: file_kubeagent_proto_depIdxs,
		MessageInfos:      file_kubeagent_proto_msgTypes,
	}.Build()
	File_kubeagent_proto = out.File
	file_kubeagent_proto_rawDesc = nil
	file_kubeagent_proto_goTypes = nil
	file_kubeagent_proto_depIdxs = nil
}
