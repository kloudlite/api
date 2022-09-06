// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.5
// source: finance.proto

package finance

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

type ComputePlan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider       string  `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	Name           string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Desc           string  `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
	MemoryUnitSize float32 `protobuf:"fixed32,4,opt,name=memoryUnitSize,proto3" json:"memoryUnitSize,omitempty"`
	CpuUnitSize    float32 `protobuf:"fixed32,5,opt,name=cpuUnitSize,proto3" json:"cpuUnitSize,omitempty"`
}

func (x *ComputePlan) Reset() {
	*x = ComputePlan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_finance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputePlan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputePlan) ProtoMessage() {}

func (x *ComputePlan) ProtoReflect() protoreflect.Message {
	mi := &file_finance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputePlan.ProtoReflect.Descriptor instead.
func (*ComputePlan) Descriptor() ([]byte, []int) {
	return file_finance_proto_rawDescGZIP(), []int{0}
}

func (x *ComputePlan) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *ComputePlan) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ComputePlan) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *ComputePlan) GetMemoryUnitSize() float32 {
	if x != nil {
		return x.MemoryUnitSize
	}
	return 0
}

func (x *ComputePlan) GetCpuUnitSize() float32 {
	if x != nil {
		return x.CpuUnitSize
	}
	return 0
}

type StartBillableIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId    string  `protobuf:"bytes,1,opt,name=accountId,proto3" json:"accountId,omitempty"`
	BillableType string  `protobuf:"bytes,2,opt,name=BillableType,proto3" json:"BillableType,omitempty"`
	Quantity     float32 `protobuf:"fixed32,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *StartBillableIn) Reset() {
	*x = StartBillableIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_finance_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartBillableIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartBillableIn) ProtoMessage() {}

func (x *StartBillableIn) ProtoReflect() protoreflect.Message {
	mi := &file_finance_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartBillableIn.ProtoReflect.Descriptor instead.
func (*StartBillableIn) Descriptor() ([]byte, []int) {
	return file_finance_proto_rawDescGZIP(), []int{1}
}

func (x *StartBillableIn) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *StartBillableIn) GetBillableType() string {
	if x != nil {
		return x.BillableType
	}
	return ""
}

func (x *StartBillableIn) GetQuantity() float32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type StopBillableIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BillableId string `protobuf:"bytes,1,opt,name=BillableId,proto3" json:"BillableId,omitempty"`
}

func (x *StopBillableIn) Reset() {
	*x = StopBillableIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_finance_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopBillableIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopBillableIn) ProtoMessage() {}

func (x *StopBillableIn) ProtoReflect() protoreflect.Message {
	mi := &file_finance_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopBillableIn.ProtoReflect.Descriptor instead.
func (*StopBillableIn) Descriptor() ([]byte, []int) {
	return file_finance_proto_rawDescGZIP(), []int{2}
}

func (x *StopBillableIn) GetBillableId() string {
	if x != nil {
		return x.BillableId
	}
	return ""
}

type StartBillableOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BillingId string `protobuf:"bytes,1,opt,name=billingId,proto3" json:"billingId,omitempty"`
}

func (x *StartBillableOut) Reset() {
	*x = StartBillableOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_finance_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartBillableOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartBillableOut) ProtoMessage() {}

func (x *StartBillableOut) ProtoReflect() protoreflect.Message {
	mi := &file_finance_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartBillableOut.ProtoReflect.Descriptor instead.
func (*StartBillableOut) Descriptor() ([]byte, []int) {
	return file_finance_proto_rawDescGZIP(), []int{3}
}

func (x *StartBillableOut) GetBillingId() string {
	if x != nil {
		return x.BillingId
	}
	return ""
}

type StopBillableOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BillingId string `protobuf:"bytes,1,opt,name=billingId,proto3" json:"billingId,omitempty"`
}

func (x *StopBillableOut) Reset() {
	*x = StopBillableOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_finance_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopBillableOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopBillableOut) ProtoMessage() {}

func (x *StopBillableOut) ProtoReflect() protoreflect.Message {
	mi := &file_finance_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopBillableOut.ProtoReflect.Descriptor instead.
func (*StopBillableOut) Descriptor() ([]byte, []int) {
	return file_finance_proto_rawDescGZIP(), []int{4}
}

func (x *StopBillableOut) GetBillingId() string {
	if x != nil {
		return x.BillingId
	}
	return ""
}

var File_finance_proto protoreflect.FileDescriptor

var file_finance_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x9b, 0x01, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x65, 0x73, 0x63, 0x12, 0x26, 0x0a, 0x0e, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x55, 0x6e, 0x69,
	0x74, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e, 0x6d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x55, 0x6e, 0x69, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63,
	0x70, 0x75, 0x55, 0x6e, 0x69, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x0b, 0x63, 0x70, 0x75, 0x55, 0x6e, 0x69, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x6f, 0x0a,
	0x0f, 0x53, 0x74, 0x61, 0x72, 0x74, 0x42, 0x69, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x6e,
	0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x22,
	0x0a, 0x0c, 0x42, 0x69, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x42, 0x69, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x30,
	0x0a, 0x0e, 0x53, 0x74, 0x6f, 0x70, 0x42, 0x69, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x6e,
	0x12, 0x1e, 0x0a, 0x0a, 0x42, 0x69, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x42, 0x69, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64,
	0x22, 0x30, 0x0a, 0x10, 0x53, 0x74, 0x61, 0x72, 0x74, 0x42, 0x69, 0x6c, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x4f, 0x75, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x49, 0x64, 0x22, 0x2f, 0x0a, 0x0f, 0x53, 0x74, 0x6f, 0x70, 0x42, 0x69, 0x6c, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x4f, 0x75, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x49, 0x64, 0x32, 0x72, 0x0a, 0x07, 0x46, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x34,
	0x0a, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x42, 0x69, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12,
	0x10, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x42, 0x69, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x49,
	0x6e, 0x1a, 0x11, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x42, 0x69, 0x6c, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x4f, 0x75, 0x74, 0x12, 0x31, 0x0a, 0x0c, 0x73, 0x74, 0x6f, 0x70, 0x42, 0x69, 0x6c, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x12, 0x0f, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x42, 0x69, 0x6c, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x49, 0x6e, 0x1a, 0x10, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x42, 0x69, 0x6c, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x4f, 0x75, 0x74, 0x42, 0x1a, 0x5a, 0x18, 0x6b, 0x6c, 0x6f, 0x75, 0x64,
	0x6c, 0x69, 0x74, 0x65, 0x2e, 0x69, 0x6f, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x66, 0x69, 0x6e, 0x61,
	0x6e, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_finance_proto_rawDescOnce sync.Once
	file_finance_proto_rawDescData = file_finance_proto_rawDesc
)

func file_finance_proto_rawDescGZIP() []byte {
	file_finance_proto_rawDescOnce.Do(func() {
		file_finance_proto_rawDescData = protoimpl.X.CompressGZIP(file_finance_proto_rawDescData)
	})
	return file_finance_proto_rawDescData
}

var file_finance_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_finance_proto_goTypes = []interface{}{
	(*ComputePlan)(nil),      // 0: ComputePlan
	(*StartBillableIn)(nil),  // 1: StartBillableIn
	(*StopBillableIn)(nil),   // 2: StopBillableIn
	(*StartBillableOut)(nil), // 3: StartBillableOut
	(*StopBillableOut)(nil),  // 4: StopBillableOut
}
var file_finance_proto_depIdxs = []int32{
	1, // 0: Finance.startBillable:input_type -> StartBillableIn
	2, // 1: Finance.stopBillable:input_type -> StopBillableIn
	3, // 2: Finance.startBillable:output_type -> StartBillableOut
	4, // 3: Finance.stopBillable:output_type -> StopBillableOut
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_finance_proto_init() }
func file_finance_proto_init() {
	if File_finance_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_finance_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputePlan); i {
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
		file_finance_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartBillableIn); i {
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
		file_finance_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopBillableIn); i {
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
		file_finance_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartBillableOut); i {
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
		file_finance_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopBillableOut); i {
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
			RawDescriptor: file_finance_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_finance_proto_goTypes,
		DependencyIndexes: file_finance_proto_depIdxs,
		MessageInfos:      file_finance_proto_msgTypes,
	}.Build()
	File_finance_proto = out.File
	file_finance_proto_rawDesc = nil
	file_finance_proto_goTypes = nil
	file_finance_proto_depIdxs = nil
}
