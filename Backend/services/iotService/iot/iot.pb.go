// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.4
// source: iot.proto

package iot

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

type IOTEmpty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *IOTEmpty) Reset() {
	*x = IOTEmpty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iot_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IOTEmpty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IOTEmpty) ProtoMessage() {}

func (x *IOTEmpty) ProtoReflect() protoreflect.Message {
	mi := &file_iot_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IOTEmpty.ProtoReflect.Descriptor instead.
func (*IOTEmpty) Descriptor() ([]byte, []int) {
	return file_iot_proto_rawDescGZIP(), []int{0}
}

type IOTHealth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *IOTHealth) Reset() {
	*x = IOTHealth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iot_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IOTHealth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IOTHealth) ProtoMessage() {}

func (x *IOTHealth) ProtoReflect() protoreflect.Message {
	mi := &file_iot_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IOTHealth.ProtoReflect.Descriptor instead.
func (*IOTHealth) Descriptor() ([]byte, []int) {
	return file_iot_proto_rawDescGZIP(), []int{1}
}

func (x *IOTHealth) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type IOTData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	SensorID  int32  `protobuf:"varint,3,opt,name=SensorID,proto3" json:"SensorID,omitempty"`
	Data      string `protobuf:"bytes,4,opt,name=Data,proto3" json:"Data,omitempty"`
	Timestamp string `protobuf:"bytes,5,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
}

func (x *IOTData) Reset() {
	*x = IOTData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iot_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IOTData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IOTData) ProtoMessage() {}

func (x *IOTData) ProtoReflect() protoreflect.Message {
	mi := &file_iot_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IOTData.ProtoReflect.Descriptor instead.
func (*IOTData) Descriptor() ([]byte, []int) {
	return file_iot_proto_rawDescGZIP(), []int{2}
}

func (x *IOTData) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *IOTData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *IOTData) GetSensorID() int32 {
	if x != nil {
		return x.SensorID
	}
	return 0
}

func (x *IOTData) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *IOTData) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

type IOTDatas struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IOTDatas []*IOTData `protobuf:"bytes,1,rep,name=IOTDatas,proto3" json:"IOTDatas,omitempty"`
}

func (x *IOTDatas) Reset() {
	*x = IOTDatas{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iot_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IOTDatas) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IOTDatas) ProtoMessage() {}

func (x *IOTDatas) ProtoReflect() protoreflect.Message {
	mi := &file_iot_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IOTDatas.ProtoReflect.Descriptor instead.
func (*IOTDatas) Descriptor() ([]byte, []int) {
	return file_iot_proto_rawDescGZIP(), []int{3}
}

func (x *IOTDatas) GetIOTDatas() []*IOTData {
	if x != nil {
		return x.IOTDatas
	}
	return nil
}

type IOTTimeframeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimeStart string `protobuf:"bytes,1,opt,name=TimeStart,proto3" json:"TimeStart,omitempty"`
	TimeEnd   string `protobuf:"bytes,2,opt,name=TimeEnd,proto3" json:"TimeEnd,omitempty"`
}

func (x *IOTTimeframeRequest) Reset() {
	*x = IOTTimeframeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iot_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IOTTimeframeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IOTTimeframeRequest) ProtoMessage() {}

func (x *IOTTimeframeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iot_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IOTTimeframeRequest.ProtoReflect.Descriptor instead.
func (*IOTTimeframeRequest) Descriptor() ([]byte, []int) {
	return file_iot_proto_rawDescGZIP(), []int{4}
}

func (x *IOTTimeframeRequest) GetTimeStart() string {
	if x != nil {
		return x.TimeStart
	}
	return ""
}

func (x *IOTTimeframeRequest) GetTimeEnd() string {
	if x != nil {
		return x.TimeEnd
	}
	return ""
}

type IOTRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *IOTRequest) Reset() {
	*x = IOTRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iot_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IOTRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IOTRequest) ProtoMessage() {}

func (x *IOTRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iot_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IOTRequest.ProtoReflect.Descriptor instead.
func (*IOTRequest) Descriptor() ([]byte, []int) {
	return file_iot_proto_rawDescGZIP(), []int{5}
}

func (x *IOTRequest) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

var File_iot_proto protoreflect.FileDescriptor

var file_iot_proto_rawDesc = []byte{
	0x0a, 0x09, 0x69, 0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0a, 0x0a, 0x08, 0x49,
	0x4f, 0x54, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x25, 0x0a, 0x09, 0x49, 0x4f, 0x54, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x7b,
	0x0a, 0x07, 0x49, 0x4f, 0x54, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74,
	0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a,
	0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x30, 0x0a, 0x08, 0x49,
	0x4f, 0x54, 0x44, 0x61, 0x74, 0x61, 0x73, 0x12, 0x24, 0x0a, 0x08, 0x49, 0x4f, 0x54, 0x44, 0x61,
	0x74, 0x61, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x49, 0x4f, 0x54, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x08, 0x49, 0x4f, 0x54, 0x44, 0x61, 0x74, 0x61, 0x73, 0x22, 0x4d, 0x0a,
	0x13, 0x49, 0x4f, 0x54, 0x54, 0x69, 0x6d, 0x65, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x69, 0x6d, 0x65, 0x45, 0x6e, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x54, 0x69, 0x6d, 0x65, 0x45, 0x6e, 0x64, 0x22, 0x1c, 0x0a, 0x0a,
	0x49, 0x4f, 0x54, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x32, 0xb6, 0x01, 0x0a, 0x0a, 0x49,
	0x6f, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x09, 0x2e, 0x49, 0x4f, 0x54, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x0a, 0x2e, 0x49, 0x4f, 0x54, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x22, 0x00, 0x12,
	0x22, 0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x08, 0x2e,
	0x49, 0x4f, 0x54, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x08, 0x2e, 0x49, 0x4f, 0x54, 0x44, 0x61, 0x74,
	0x61, 0x22, 0x00, 0x12, 0x24, 0x0a, 0x08, 0x52, 0x65, 0x61, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x0b, 0x2e, 0x49, 0x4f, 0x54, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x49,
	0x4f, 0x54, 0x44, 0x61, 0x74, 0x61, 0x73, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x13, 0x52, 0x65, 0x61,
	0x64, 0x44, 0x61, 0x74, 0x61, 0x49, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x46, 0x72, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x2e, 0x49, 0x4f, 0x54, 0x54, 0x69, 0x6d, 0x65, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x49, 0x4f, 0x54, 0x44, 0x61, 0x74, 0x61,
	0x73, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x3b, 0x69, 0x6f, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_iot_proto_rawDescOnce sync.Once
	file_iot_proto_rawDescData = file_iot_proto_rawDesc
)

func file_iot_proto_rawDescGZIP() []byte {
	file_iot_proto_rawDescOnce.Do(func() {
		file_iot_proto_rawDescData = protoimpl.X.CompressGZIP(file_iot_proto_rawDescData)
	})
	return file_iot_proto_rawDescData
}

var file_iot_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_iot_proto_goTypes = []interface{}{
	(*IOTEmpty)(nil),            // 0: IOTEmpty
	(*IOTHealth)(nil),           // 1: IOTHealth
	(*IOTData)(nil),             // 2: IOTData
	(*IOTDatas)(nil),            // 3: IOTDatas
	(*IOTTimeframeRequest)(nil), // 4: IOTTimeframeRequest
	(*IOTRequest)(nil),          // 5: IOTRequest
}
var file_iot_proto_depIdxs = []int32{
	2, // 0: IOTDatas.IOTDatas:type_name -> IOTData
	0, // 1: IotService.GetHealth:input_type -> IOTEmpty
	2, // 2: IotService.UploadData:input_type -> IOTData
	5, // 3: IotService.ReadData:input_type -> IOTRequest
	4, // 4: IotService.ReadDataInTimeFrame:input_type -> IOTTimeframeRequest
	1, // 5: IotService.GetHealth:output_type -> IOTHealth
	2, // 6: IotService.UploadData:output_type -> IOTData
	3, // 7: IotService.ReadData:output_type -> IOTDatas
	3, // 8: IotService.ReadDataInTimeFrame:output_type -> IOTDatas
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_iot_proto_init() }
func file_iot_proto_init() {
	if File_iot_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_iot_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IOTEmpty); i {
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
		file_iot_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IOTHealth); i {
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
		file_iot_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IOTData); i {
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
		file_iot_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IOTDatas); i {
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
		file_iot_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IOTTimeframeRequest); i {
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
		file_iot_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IOTRequest); i {
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
			RawDescriptor: file_iot_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_iot_proto_goTypes,
		DependencyIndexes: file_iot_proto_depIdxs,
		MessageInfos:      file_iot_proto_msgTypes,
	}.Build()
	File_iot_proto = out.File
	file_iot_proto_rawDesc = nil
	file_iot_proto_goTypes = nil
	file_iot_proto_depIdxs = nil
}
