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

type BedsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookedTime string `protobuf:"bytes,1,opt,name=BookedTime,proto3" json:"BookedTime,omitempty"`
	BookedEnd  string `protobuf:"bytes,2,opt,name=BookedEnd,proto3" json:"BookedEnd,omitempty"`
	HospitalId int32  `protobuf:"varint,3,opt,name=HospitalId,proto3" json:"HospitalId,omitempty"`
}

func (x *BedsRequest) Reset() {
	*x = BedsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_useradmin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BedsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BedsRequest) ProtoMessage() {}

func (x *BedsRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use BedsRequest.ProtoReflect.Descriptor instead.
func (*BedsRequest) Descriptor() ([]byte, []int) {
	return file_useradmin_proto_rawDescGZIP(), []int{2}
}

func (x *BedsRequest) GetBookedTime() string {
	if x != nil {
		return x.BookedTime
	}
	return ""
}

func (x *BedsRequest) GetBookedEnd() string {
	if x != nil {
		return x.BookedEnd
	}
	return ""
}

func (x *BedsRequest) GetHospitalId() int32 {
	if x != nil {
		return x.HospitalId
	}
	return 0
}

type UserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_useradmin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UserRequest.ProtoReflect.Descriptor instead.
func (*UserRequest) Descriptor() ([]byte, []int) {
	return file_useradmin_proto_rawDescGZIP(), []int{3}
}

func (x *UserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type UAUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address    string `protobuf:"bytes,1,opt,name=Address,proto3" json:"Address,omitempty"`
	City       string `protobuf:"bytes,2,opt,name=City,proto3" json:"City,omitempty"`
	Country    string `protobuf:"bytes,3,opt,name=Country,proto3" json:"Country,omitempty"`
	Name       string `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`
	PostalCode string `protobuf:"bytes,5,opt,name=PostalCode,proto3" json:"PostalCode,omitempty"`
	SocialIdNr string `protobuf:"bytes,6,opt,name=SocialIdNr,proto3" json:"SocialIdNr,omitempty"`
	Username   string `protobuf:"bytes,7,opt,name=Username,proto3" json:"Username,omitempty"`
}

func (x *UAUser) Reset() {
	*x = UAUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_useradmin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UAUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UAUser) ProtoMessage() {}

func (x *UAUser) ProtoReflect() protoreflect.Message {
	mi := &file_useradmin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UAUser.ProtoReflect.Descriptor instead.
func (*UAUser) Descriptor() ([]byte, []int) {
	return file_useradmin_proto_rawDescGZIP(), []int{4}
}

func (x *UAUser) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *UAUser) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *UAUser) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *UAUser) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UAUser) GetPostalCode() string {
	if x != nil {
		return x.PostalCode
	}
	return ""
}

func (x *UAUser) GetSocialIdNr() string {
	if x != nil {
		return x.SocialIdNr
	}
	return ""
}

func (x *UAUser) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type Users struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*UAUser `protobuf:"bytes,1,rep,name=Users,proto3" json:"Users,omitempty"`
}

func (x *Users) Reset() {
	*x = Users{}
	if protoimpl.UnsafeEnabled {
		mi := &file_useradmin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Users) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Users) ProtoMessage() {}

func (x *Users) ProtoReflect() protoreflect.Message {
	mi := &file_useradmin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Users.ProtoReflect.Descriptor instead.
func (*Users) Descriptor() ([]byte, []int) {
	return file_useradmin_proto_rawDescGZIP(), []int{5}
}

func (x *Users) GetUsers() []*UAUser {
	if x != nil {
		return x.Users
	}
	return nil
}

type UAHospital struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HospitalId int32  `protobuf:"varint,1,opt,name=HospitalId,proto3" json:"HospitalId,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Address    string `protobuf:"bytes,3,opt,name=Address,proto3" json:"Address,omitempty"`
	City       string `protobuf:"bytes,4,opt,name=City,proto3" json:"City,omitempty"`
	PostalCode string `protobuf:"bytes,5,opt,name=PostalCode,proto3" json:"PostalCode,omitempty"`
	Country    string `protobuf:"bytes,6,opt,name=Country,proto3" json:"Country,omitempty"`
}

func (x *UAHospital) Reset() {
	*x = UAHospital{}
	if protoimpl.UnsafeEnabled {
		mi := &file_useradmin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UAHospital) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UAHospital) ProtoMessage() {}

func (x *UAHospital) ProtoReflect() protoreflect.Message {
	mi := &file_useradmin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UAHospital.ProtoReflect.Descriptor instead.
func (*UAHospital) Descriptor() ([]byte, []int) {
	return file_useradmin_proto_rawDescGZIP(), []int{6}
}

func (x *UAHospital) GetHospitalId() int32 {
	if x != nil {
		return x.HospitalId
	}
	return 0
}

func (x *UAHospital) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UAHospital) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *UAHospital) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *UAHospital) GetPostalCode() string {
	if x != nil {
		return x.PostalCode
	}
	return ""
}

func (x *UAHospital) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

type Hospitals struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hospitals []*UAHospital `protobuf:"bytes,1,rep,name=Hospitals,proto3" json:"Hospitals,omitempty"`
}

func (x *Hospitals) Reset() {
	*x = Hospitals{}
	if protoimpl.UnsafeEnabled {
		mi := &file_useradmin_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hospitals) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hospitals) ProtoMessage() {}

func (x *Hospitals) ProtoReflect() protoreflect.Message {
	mi := &file_useradmin_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hospitals.ProtoReflect.Descriptor instead.
func (*Hospitals) Descriptor() ([]byte, []int) {
	return file_useradmin_proto_rawDescGZIP(), []int{7}
}

func (x *Hospitals) GetHospitals() []*UAHospital {
	if x != nil {
		return x.Hospitals
	}
	return nil
}

type Bed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BedId        int32       `protobuf:"varint,1,opt,name=BedId,proto3" json:"BedId,omitempty"`
	Name         string      `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Departmentid int32       `protobuf:"varint,3,opt,name=Departmentid,proto3" json:"Departmentid,omitempty"`
	IsAvailable  bool        `protobuf:"varint,4,opt,name=isAvailable,proto3" json:"isAvailable,omitempty"`
	Department   *Department `protobuf:"bytes,5,opt,name=Department,proto3" json:"Department,omitempty"`
}

func (x *Bed) Reset() {
	*x = Bed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_useradmin_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bed) ProtoMessage() {}

func (x *Bed) ProtoReflect() protoreflect.Message {
	mi := &file_useradmin_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bed.ProtoReflect.Descriptor instead.
func (*Bed) Descriptor() ([]byte, []int) {
	return file_useradmin_proto_rawDescGZIP(), []int{8}
}

func (x *Bed) GetBedId() int32 {
	if x != nil {
		return x.BedId
	}
	return 0
}

func (x *Bed) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Bed) GetDepartmentid() int32 {
	if x != nil {
		return x.Departmentid
	}
	return 0
}

func (x *Bed) GetIsAvailable() bool {
	if x != nil {
		return x.IsAvailable
	}
	return false
}

func (x *Bed) GetDepartment() *Department {
	if x != nil {
		return x.Department
	}
	return nil
}

type Beds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Beds []*Bed `protobuf:"bytes,1,rep,name=Beds,proto3" json:"Beds,omitempty"`
}

func (x *Beds) Reset() {
	*x = Beds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_useradmin_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Beds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Beds) ProtoMessage() {}

func (x *Beds) ProtoReflect() protoreflect.Message {
	mi := &file_useradmin_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Beds.ProtoReflect.Descriptor instead.
func (*Beds) Descriptor() ([]byte, []int) {
	return file_useradmin_proto_rawDescGZIP(), []int{9}
}

func (x *Beds) GetBeds() []*Bed {
	if x != nil {
		return x.Beds
	}
	return nil
}

type Department struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Departmentid int32  `protobuf:"varint,1,opt,name=Departmentid,proto3" json:"Departmentid,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Description  string `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	HospitalId   int32  `protobuf:"varint,4,opt,name=HospitalId,proto3" json:"HospitalId,omitempty"`
}

func (x *Department) Reset() {
	*x = Department{}
	if protoimpl.UnsafeEnabled {
		mi := &file_useradmin_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Department) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Department) ProtoMessage() {}

func (x *Department) ProtoReflect() protoreflect.Message {
	mi := &file_useradmin_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Department.ProtoReflect.Descriptor instead.
func (*Department) Descriptor() ([]byte, []int) {
	return file_useradmin_proto_rawDescGZIP(), []int{10}
}

func (x *Department) GetDepartmentid() int32 {
	if x != nil {
		return x.Departmentid
	}
	return 0
}

func (x *Department) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Department) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Department) GetHospitalId() int32 {
	if x != nil {
		return x.HospitalId
	}
	return 0
}

type Departments struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Departments []*Department `protobuf:"bytes,1,rep,name=Departments,proto3" json:"Departments,omitempty"`
}

func (x *Departments) Reset() {
	*x = Departments{}
	if protoimpl.UnsafeEnabled {
		mi := &file_useradmin_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Departments) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Departments) ProtoMessage() {}

func (x *Departments) ProtoReflect() protoreflect.Message {
	mi := &file_useradmin_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Departments.ProtoReflect.Descriptor instead.
func (*Departments) Descriptor() ([]byte, []int) {
	return file_useradmin_proto_rawDescGZIP(), []int{11}
}

func (x *Departments) GetDepartments() []*Department {
	if x != nil {
		return x.Departments
	}
	return nil
}

var File_useradmin_proto protoreflect.FileDescriptor

var file_useradmin_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x09, 0x0a, 0x07, 0x55, 0x41, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x24, 0x0a, 0x08,
	0x55, 0x41, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x6b, 0x0a, 0x0b, 0x42, 0x65, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x42, 0x6f, 0x6f, 0x6b, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x42, 0x6f, 0x6f, 0x6b, 0x65, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x6b, 0x65, 0x64, 0x45, 0x6e, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x42, 0x6f, 0x6f, 0x6b, 0x65, 0x64, 0x45, 0x6e, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x48, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x48, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x49, 0x64, 0x22,
	0x29, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xc0, 0x01, 0x0a, 0x06, 0x55,
	0x41, 0x55, 0x73, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x43, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43,
	0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x49, 0x64, 0x4e, 0x72, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x49, 0x64, 0x4e,
	0x72, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x26, 0x0a,
	0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x1d, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x55, 0x41, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x22, 0xa8, 0x01, 0x0a, 0x0a, 0x55, 0x41, 0x48, 0x6f, 0x73, 0x70,
	0x69, 0x74, 0x61, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x48, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x48, 0x6f, 0x73, 0x70, 0x69, 0x74,
	0x61, 0x6c, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x43, 0x69, 0x74, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x6f, 0x73, 0x74, 0x61, 0x6c,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x6f, 0x73, 0x74,
	0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x22, 0x36, 0x0a, 0x09, 0x48, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x73, 0x12, 0x29, 0x0a,
	0x09, 0x48, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x55, 0x41, 0x48, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x52, 0x09, 0x48,
	0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x73, 0x22, 0xa2, 0x01, 0x0a, 0x03, 0x42, 0x65, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x42, 0x65, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x42, 0x65, 0x64, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x44, 0x65,
	0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0c, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x69, 0x64, 0x12, 0x20,
	0x0a, 0x0b, 0x69, 0x73, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x12, 0x2b, 0x0a, 0x0a, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x0a, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x20, 0x0a,
	0x04, 0x42, 0x65, 0x64, 0x73, 0x12, 0x18, 0x0a, 0x04, 0x42, 0x65, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x04, 0x2e, 0x42, 0x65, 0x64, 0x52, 0x04, 0x42, 0x65, 0x64, 0x73, 0x22,
	0x86, 0x01, 0x0a, 0x0a, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x22,
	0x0a, 0x0c, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x48, 0x6f, 0x73, 0x70,
	0x69, 0x74, 0x61, 0x6c, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x48, 0x6f,
	0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x49, 0x64, 0x22, 0x3c, 0x0a, 0x0b, 0x44, 0x65, 0x70, 0x61,
	0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x2d, 0x0a, 0x0b, 0x44, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x44,
	0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x44, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x32, 0x9e, 0x02, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x08, 0x2e, 0x55, 0x41, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x09, 0x2e, 0x55, 0x41, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x22, 0x00, 0x12,
	0x26, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x0c,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x55,
	0x41, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x12, 0x21, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x61,
	0x74, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x08, 0x2e, 0x55, 0x41, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x06, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x08, 0x2e, 0x55,
	0x41, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0c, 0x2e, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x22, 0x00, 0x12, 0x1c, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x65, 0x64,
	0x73, 0x12, 0x08, 0x2e, 0x55, 0x41, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x05, 0x2e, 0x42, 0x65,
	0x64, 0x73, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x42, 0x65, 0x64, 0x73, 0x12, 0x0c, 0x2e, 0x42, 0x65, 0x64, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x05, 0x2e, 0x42, 0x65, 0x64, 0x73, 0x22, 0x00, 0x12,
	0x26, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x73, 0x12,
	0x08, 0x2e, 0x55, 0x41, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0a, 0x2e, 0x48, 0x6f, 0x73, 0x70,
	0x69, 0x74, 0x61, 0x6c, 0x73, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x3b, 0x75, 0x73, 0x65,
	0x72, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_useradmin_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_useradmin_proto_goTypes = []interface{}{
	(*UAEmpty)(nil),     // 0: UAEmpty
	(*UAHealth)(nil),    // 1: UAHealth
	(*BedsRequest)(nil), // 2: BedsRequest
	(*UserRequest)(nil), // 3: UserRequest
	(*UAUser)(nil),      // 4: UAUser
	(*Users)(nil),       // 5: Users
	(*UAHospital)(nil),  // 6: UAHospital
	(*Hospitals)(nil),   // 7: Hospitals
	(*Bed)(nil),         // 8: Bed
	(*Beds)(nil),        // 9: Beds
	(*Department)(nil),  // 10: Department
	(*Departments)(nil), // 11: Departments
}
var file_useradmin_proto_depIdxs = []int32{
	4,  // 0: Users.Users:type_name -> UAUser
	6,  // 1: Hospitals.Hospitals:type_name -> UAHospital
	10, // 2: Bed.Department:type_name -> Department
	8,  // 3: Beds.Beds:type_name -> Bed
	10, // 4: Departments.Departments:type_name -> Department
	0,  // 5: UseradminService.GetHealth:input_type -> UAEmpty
	3,  // 6: UseradminService.GetEmployee:input_type -> UserRequest
	0,  // 7: UseradminService.GetPatients:input_type -> UAEmpty
	0,  // 8: UseradminService.GetDepartments:input_type -> UAEmpty
	0,  // 9: UseradminService.GetBeds:input_type -> UAEmpty
	2,  // 10: UseradminService.GetAvailableBeds:input_type -> BedsRequest
	0,  // 11: UseradminService.GetHospitals:input_type -> UAEmpty
	1,  // 12: UseradminService.GetHealth:output_type -> UAHealth
	4,  // 13: UseradminService.GetEmployee:output_type -> UAUser
	5,  // 14: UseradminService.GetPatients:output_type -> Users
	11, // 15: UseradminService.GetDepartments:output_type -> Departments
	9,  // 16: UseradminService.GetBeds:output_type -> Beds
	9,  // 17: UseradminService.GetAvailableBeds:output_type -> Beds
	7,  // 18: UseradminService.GetHospitals:output_type -> Hospitals
	12, // [12:19] is the sub-list for method output_type
	5,  // [5:12] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
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
			switch v := v.(*BedsRequest); i {
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
			switch v := v.(*UserRequest); i {
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
		file_useradmin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UAUser); i {
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
		file_useradmin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Users); i {
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
		file_useradmin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UAHospital); i {
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
		file_useradmin_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hospitals); i {
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
		file_useradmin_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bed); i {
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
		file_useradmin_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Beds); i {
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
		file_useradmin_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Department); i {
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
		file_useradmin_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Departments); i {
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
			NumMessages:   12,
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