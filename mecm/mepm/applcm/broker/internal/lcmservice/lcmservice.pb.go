// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0-devel
// 	protoc        v3.11.4
// source: lcmservice.proto

package lcmservice

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type InstantiateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//	*InstantiateRequest_HostIp
	//	*InstantiateRequest_Package
	Data isInstantiateRequest_Data `protobuf_oneof:"data"`
}

func (x *InstantiateRequest) Reset() {
	*x = InstantiateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lcmservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstantiateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstantiateRequest) ProtoMessage() {}

func (x *InstantiateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lcmservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstantiateRequest.ProtoReflect.Descriptor instead.
func (*InstantiateRequest) Descriptor() ([]byte, []int) {
	return file_lcmservice_proto_rawDescGZIP(), []int{0}
}

func (m *InstantiateRequest) GetData() isInstantiateRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *InstantiateRequest) GetHostIp() string {
	if x, ok := x.GetData().(*InstantiateRequest_HostIp); ok {
		return x.HostIp
	}
	return ""
}

func (x *InstantiateRequest) GetPackage() []byte {
	if x, ok := x.GetData().(*InstantiateRequest_Package); ok {
		return x.Package
	}
	return nil
}

type isInstantiateRequest_Data interface {
	isInstantiateRequest_Data()
}

type InstantiateRequest_HostIp struct {
	HostIp string `protobuf:"bytes,1,opt,name=hostIp,proto3,oneof"`
}

type InstantiateRequest_Package struct {
	Package []byte `protobuf:"bytes,2,opt,name=package,proto3,oneof"`
}

func (*InstantiateRequest_HostIp) isInstantiateRequest_Data() {}

func (*InstantiateRequest_Package) isInstantiateRequest_Data() {}

type InstantiateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkloadId string `protobuf:"bytes,1,opt,name=workloadId,proto3" json:"workloadId,omitempty"`
	Status     string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *InstantiateResponse) Reset() {
	*x = InstantiateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lcmservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstantiateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstantiateResponse) ProtoMessage() {}

func (x *InstantiateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lcmservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstantiateResponse.ProtoReflect.Descriptor instead.
func (*InstantiateResponse) Descriptor() ([]byte, []int) {
	return file_lcmservice_proto_rawDescGZIP(), []int{1}
}

func (x *InstantiateResponse) GetWorkloadId() string {
	if x != nil {
		return x.WorkloadId
	}
	return ""
}

func (x *InstantiateResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type TerminateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostIp     string `protobuf:"bytes,1,opt,name=hostIp,proto3" json:"hostIp,omitempty"`
	WorkloadId string `protobuf:"bytes,2,opt,name=workloadId,proto3" json:"workloadId,omitempty"`
}

func (x *TerminateRequest) Reset() {
	*x = TerminateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lcmservice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerminateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerminateRequest) ProtoMessage() {}

func (x *TerminateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lcmservice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerminateRequest.ProtoReflect.Descriptor instead.
func (*TerminateRequest) Descriptor() ([]byte, []int) {
	return file_lcmservice_proto_rawDescGZIP(), []int{2}
}

func (x *TerminateRequest) GetHostIp() string {
	if x != nil {
		return x.HostIp
	}
	return ""
}

func (x *TerminateRequest) GetWorkloadId() string {
	if x != nil {
		return x.WorkloadId
	}
	return ""
}

type TerminateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *TerminateResponse) Reset() {
	*x = TerminateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lcmservice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerminateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerminateResponse) ProtoMessage() {}

func (x *TerminateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lcmservice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerminateResponse.ProtoReflect.Descriptor instead.
func (*TerminateResponse) Descriptor() ([]byte, []int) {
	return file_lcmservice_proto_rawDescGZIP(), []int{3}
}

func (x *TerminateResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type QueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostIp     string `protobuf:"bytes,1,opt,name=hostIp,proto3" json:"hostIp,omitempty"`
	WorkloadId string `protobuf:"bytes,2,opt,name=workloadId,proto3" json:"workloadId,omitempty"`
}

func (x *QueryRequest) Reset() {
	*x = QueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lcmservice_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRequest) ProtoMessage() {}

func (x *QueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lcmservice_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRequest.ProtoReflect.Descriptor instead.
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return file_lcmservice_proto_rawDescGZIP(), []int{4}
}

func (x *QueryRequest) GetHostIp() string {
	if x != nil {
		return x.HostIp
	}
	return ""
}

func (x *QueryRequest) GetWorkloadId() string {
	if x != nil {
		return x.WorkloadId
	}
	return ""
}

type QueryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *QueryResponse) Reset() {
	*x = QueryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lcmservice_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryResponse) ProtoMessage() {}

func (x *QueryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lcmservice_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryResponse.ProtoReflect.Descriptor instead.
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return file_lcmservice_proto_rawDescGZIP(), []int{5}
}

func (x *QueryResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_lcmservice_proto protoreflect.FileDescriptor

var file_lcmservice_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6c, 0x63, 0x6d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x6c, 0x63, 0x6d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x52,
	0x0a, 0x12, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x06, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x70, 0x12, 0x1a,
	0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48,
	0x00, 0x52, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x4d, 0x0a, 0x13, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x77, 0x6f, 0x72,
	0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77,
	0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x4a, 0x0a, 0x10, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x70, 0x12, 0x1e, 0x0a,
	0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x64, 0x22, 0x2b, 0x0a,
	0x11, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x46, 0x0a, 0x0c, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x6f,
	0x73, 0x74, 0x49, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x6f, 0x73, 0x74,
	0x49, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64,
	0x49, 0x64, 0x22, 0x27, 0x0a, 0x0d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xe8, 0x01, 0x0a, 0x06,
	0x41, 0x70, 0x70, 0x4c, 0x43, 0x4d, 0x12, 0x52, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x74, 0x69, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x6c, 0x63, 0x6d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6c, 0x63, 0x6d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x4a, 0x0a, 0x09, 0x74, 0x65,
	0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x6c, 0x63, 0x6d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6c, 0x63, 0x6d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12,
	0x18, 0x2e, 0x6c, 0x63, 0x6d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6c, 0x63, 0x6d, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lcmservice_proto_rawDescOnce sync.Once
	file_lcmservice_proto_rawDescData = file_lcmservice_proto_rawDesc
)

func file_lcmservice_proto_rawDescGZIP() []byte {
	file_lcmservice_proto_rawDescOnce.Do(func() {
		file_lcmservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_lcmservice_proto_rawDescData)
	})
	return file_lcmservice_proto_rawDescData
}

var file_lcmservice_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_lcmservice_proto_goTypes = []interface{}{
	(*InstantiateRequest)(nil),  // 0: lcmservice.InstantiateRequest
	(*InstantiateResponse)(nil), // 1: lcmservice.InstantiateResponse
	(*TerminateRequest)(nil),    // 2: lcmservice.TerminateRequest
	(*TerminateResponse)(nil),   // 3: lcmservice.TerminateResponse
	(*QueryRequest)(nil),        // 4: lcmservice.QueryRequest
	(*QueryResponse)(nil),       // 5: lcmservice.QueryResponse
}
var file_lcmservice_proto_depIdxs = []int32{
	0, // 0: lcmservice.AppLCM.instantiate:input_type -> lcmservice.InstantiateRequest
	2, // 1: lcmservice.AppLCM.terminate:input_type -> lcmservice.TerminateRequest
	4, // 2: lcmservice.AppLCM.query:input_type -> lcmservice.QueryRequest
	1, // 3: lcmservice.AppLCM.instantiate:output_type -> lcmservice.InstantiateResponse
	3, // 4: lcmservice.AppLCM.terminate:output_type -> lcmservice.TerminateResponse
	5, // 5: lcmservice.AppLCM.query:output_type -> lcmservice.QueryResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_lcmservice_proto_init() }
func file_lcmservice_proto_init() {
	if File_lcmservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lcmservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstantiateRequest); i {
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
		file_lcmservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstantiateResponse); i {
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
		file_lcmservice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerminateRequest); i {
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
		file_lcmservice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerminateResponse); i {
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
		file_lcmservice_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRequest); i {
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
		file_lcmservice_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryResponse); i {
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
	file_lcmservice_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*InstantiateRequest_HostIp)(nil),
		(*InstantiateRequest_Package)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_lcmservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_lcmservice_proto_goTypes,
		DependencyIndexes: file_lcmservice_proto_depIdxs,
		MessageInfos:      file_lcmservice_proto_msgTypes,
	}.Build()
	File_lcmservice_proto = out.File
	file_lcmservice_proto_rawDesc = nil
	file_lcmservice_proto_goTypes = nil
	file_lcmservice_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AppLCMClient is the client API for AppLCM service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AppLCMClient interface {
	Instantiate(ctx context.Context, opts ...grpc.CallOption) (AppLCM_InstantiateClient, error)
	Terminate(ctx context.Context, in *TerminateRequest, opts ...grpc.CallOption) (*TerminateResponse, error)
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
}

type appLCMClient struct {
	cc grpc.ClientConnInterface
}

func NewAppLCMClient(cc grpc.ClientConnInterface) AppLCMClient {
	return &appLCMClient{cc}
}

func (c *appLCMClient) Instantiate(ctx context.Context, opts ...grpc.CallOption) (AppLCM_InstantiateClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AppLCM_serviceDesc.Streams[0], "/lcmservice.AppLCM/instantiate", opts...)
	if err != nil {
		return nil, err
	}
	x := &appLCMInstantiateClient{stream}
	return x, nil
}

type AppLCM_InstantiateClient interface {
	Send(*InstantiateRequest) error
	CloseAndRecv() (*InstantiateResponse, error)
	grpc.ClientStream
}

type appLCMInstantiateClient struct {
	grpc.ClientStream
}

func (x *appLCMInstantiateClient) Send(m *InstantiateRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *appLCMInstantiateClient) CloseAndRecv() (*InstantiateResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(InstantiateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *appLCMClient) Terminate(ctx context.Context, in *TerminateRequest, opts ...grpc.CallOption) (*TerminateResponse, error) {
	out := new(TerminateResponse)
	err := c.cc.Invoke(ctx, "/lcmservice.AppLCM/terminate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appLCMClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, "/lcmservice.AppLCM/query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppLCMServer is the server API for AppLCM service.
type AppLCMServer interface {
	Instantiate(AppLCM_InstantiateServer) error
	Terminate(context.Context, *TerminateRequest) (*TerminateResponse, error)
	Query(context.Context, *QueryRequest) (*QueryResponse, error)
}

// UnimplementedAppLCMServer can be embedded to have forward compatible implementations.
type UnimplementedAppLCMServer struct {
}

func (*UnimplementedAppLCMServer) Instantiate(AppLCM_InstantiateServer) error {
	return status.Errorf(codes.Unimplemented, "method Instantiate not implemented")
}
func (*UnimplementedAppLCMServer) Terminate(context.Context, *TerminateRequest) (*TerminateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Terminate not implemented")
}
func (*UnimplementedAppLCMServer) Query(context.Context, *QueryRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}

func RegisterAppLCMServer(s *grpc.Server, srv AppLCMServer) {
	s.RegisterService(&_AppLCM_serviceDesc, srv)
}

func _AppLCM_Instantiate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AppLCMServer).Instantiate(&appLCMInstantiateServer{stream})
}

type AppLCM_InstantiateServer interface {
	SendAndClose(*InstantiateResponse) error
	Recv() (*InstantiateRequest, error)
	grpc.ServerStream
}

type appLCMInstantiateServer struct {
	grpc.ServerStream
}

func (x *appLCMInstantiateServer) SendAndClose(m *InstantiateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *appLCMInstantiateServer) Recv() (*InstantiateRequest, error) {
	m := new(InstantiateRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _AppLCM_Terminate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TerminateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppLCMServer).Terminate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lcmservice.AppLCM/Terminate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppLCMServer).Terminate(ctx, req.(*TerminateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppLCM_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppLCMServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lcmservice.AppLCM/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppLCMServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AppLCM_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lcmservice.AppLCM",
	HandlerType: (*AppLCMServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "terminate",
			Handler:    _AppLCM_Terminate_Handler,
		},
		{
			MethodName: "query",
			Handler:    _AppLCM_Query_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "instantiate",
			Handler:       _AppLCM_Instantiate_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "lcmservice.proto",
}
