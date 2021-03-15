// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.4
// source: useradmin.proto

package useradmin

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type UAEmpty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UAEmpty) Reset() {
	*x = UAEmpty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_useradmin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UAEmpty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UAEmpty) ProtoMessage() {}

func (x *UAEmpty) ProtoReflect() protoreflect.Message {
	mi := &file_useradmin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UAEmpty.ProtoReflect.Descriptor instead.
func (*UAEmpty) Descriptor() ([]byte, []int) {
	return file_useradmin_proto_rawDescGZIP(), []int{0}
}

type UAHealth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *UAHealth) Reset() {
	*x = UAHealth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_useradmin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UAHealth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UAHealth) ProtoMessage() {}

func (x *UAHealth) ProtoReflect() protoreflect.Message {
	mi := &file_useradmin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UAHealth.ProtoReflect.Descriptor instead.
func (*UAHealth) Descriptor() ([]byte, []int) {
	return file_useradmin_proto_rawDescGZIP(), []int{1}
}

func (x *UAHealth) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type EmployeeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Employee string `protobuf:"bytes,1,opt,name=Employee,proto3" json:"Employee,omitempty"`
}

func (x *EmployeeRequest) Reset() {
	*x = EmployeeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_useradmin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmployeeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmployeeRequest) ProtoMessage() {}

func (x *EmployeeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_useradmin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmployeeRequest.ProtoReflect.Descriptor instead.
func (*EmployeeRequest) Descriptor() ([]byte, []int) {
	return file_useradmin_proto_rawDescGZIP(), []int{2}
}

func (x *EmployeeRequest) GetEmployee() string {
	if x != nil {
		return x.Employee
	}
	return ""
}

type Employee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	WorktitleId  int32  `protobuf:"varint,2,opt,name=WorktitleId,proto3" json:"WorktitleId,omitempty"`
	DepartmentId int32  `protobuf:"varint,3,opt,name=DepartmentId,proto3" json:"DepartmentId,omitempty"`
	Username     string `protobuf:"bytes,4,opt,name=Username,proto3" json:"Username,omitempty"`
}

func (x *Employee) Reset() {
	*x = Employee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_useradmin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Employee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Employee) ProtoMessage() {}

func (x *Employee) ProtoReflect() protoreflect.Message {
	mi := &file_useradmin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Employee.ProtoReflect.Descriptor instead.
func (*Employee) Descriptor() ([]byte, []int) {
	return file_useradmin_proto_rawDescGZIP(), []int{3}
}

func (x *Employee) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Employee) GetWorktitleId() int32 {
	if x != nil {
		return x.WorktitleId
	}
	return 0
}

func (x *Employee) GetDepartmentId() int32 {
	if x != nil {
		return x.DepartmentId
	}
	return 0
}

func (x *Employee) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

var File_useradmin_proto protoreflect.FileDescriptor

var file_useradmin_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x09, 0x0a, 0x07, 0x55, 0x41, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x24, 0x0a, 0x08,
	0x55, 0x41, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x2d, 0x0a, 0x0f, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x65, 0x22, 0x80, 0x01, 0x0a, 0x08, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x57, 0x6f, 0x72, 0x6b, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x57, 0x6f, 0x72, 0x6b, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x44, 0x65, 0x70, 0x61,
	0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x32, 0x64, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x08, 0x2e, 0x55, 0x41, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x09, 0x2e, 0x55, 0x41, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x10, 0x2e, 0x45, 0x6d,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e,
	0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x3b,
	0x75, 0x73, 0x65, 0x72, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_useradmin_proto_rawDescOnce sync.Once
	file_useradmin_proto_rawDescData = file_useradmin_proto_rawDesc
)

func file_useradmin_proto_rawDescGZIP() []byte {
	file_useradmin_proto_rawDescOnce.Do(func() {
		file_useradmin_proto_rawDescData = protoimpl.X.CompressGZIP(file_useradmin_proto_rawDescData)
	})
	return file_useradmin_proto_rawDescData
}

var file_useradmin_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_useradmin_proto_goTypes = []interface{}{
	(*UAEmpty)(nil),         // 0: UAEmpty
	(*UAHealth)(nil),        // 1: UAHealth
	(*EmployeeRequest)(nil), // 2: EmployeeRequest
	(*Employee)(nil),        // 3: Employee
}
var file_useradmin_proto_depIdxs = []int32{
	0, // 0: UseradminService.GetHealth:input_type -> UAEmpty
	2, // 1: UseradminService.GetEmployee:input_type -> EmployeeRequest
	1, // 2: UseradminService.GetHealth:output_type -> UAHealth
	3, // 3: UseradminService.GetEmployee:output_type -> Employee
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_useradmin_proto_init() }
func file_useradmin_proto_init() {
	if File_useradmin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_useradmin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UAEmpty); i {
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
		file_useradmin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UAHealth); i {
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
		file_useradmin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmployeeRequest); i {
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
		file_useradmin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Employee); i {
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
			RawDescriptor: file_useradmin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_useradmin_proto_goTypes,
		DependencyIndexes: file_useradmin_proto_depIdxs,
		MessageInfos:      file_useradmin_proto_msgTypes,
	}.Build()
	File_useradmin_proto = out.File
	file_useradmin_proto_rawDesc = nil
	file_useradmin_proto_goTypes = nil
	file_useradmin_proto_depIdxs = nil
}
