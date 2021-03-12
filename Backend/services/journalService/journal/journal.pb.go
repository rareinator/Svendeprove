// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.4
// source: journal.proto

package journal

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

type JEmpty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *JEmpty) Reset() {
	*x = JEmpty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_journal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JEmpty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JEmpty) ProtoMessage() {}

func (x *JEmpty) ProtoReflect() protoreflect.Message {
	mi := &file_journal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JEmpty.ProtoReflect.Descriptor instead.
func (*JEmpty) Descriptor() ([]byte, []int) {
	return file_journal_proto_rawDescGZIP(), []int{0}
}

type Journal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JournalId    int32  `protobuf:"varint,1,opt,name=JournalId,proto3" json:"JournalId,omitempty"`
	CreationTime string `protobuf:"bytes,2,opt,name=CreationTime,proto3" json:"CreationTime,omitempty"`
	Intro        string `protobuf:"bytes,3,opt,name=Intro,proto3" json:"Intro,omitempty"`
	PatientId    int32  `protobuf:"varint,4,opt,name=PatientId,proto3" json:"PatientId,omitempty"`
	CreatedBy    int32  `protobuf:"varint,5,opt,name=CreatedBy,proto3" json:"CreatedBy,omitempty"`
}

func (x *Journal) Reset() {
	*x = Journal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_journal_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Journal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Journal) ProtoMessage() {}

func (x *Journal) ProtoReflect() protoreflect.Message {
	mi := &file_journal_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Journal.ProtoReflect.Descriptor instead.
func (*Journal) Descriptor() ([]byte, []int) {
	return file_journal_proto_rawDescGZIP(), []int{1}
}

func (x *Journal) GetJournalId() int32 {
	if x != nil {
		return x.JournalId
	}
	return 0
}

func (x *Journal) GetCreationTime() string {
	if x != nil {
		return x.CreationTime
	}
	return ""
}

func (x *Journal) GetIntro() string {
	if x != nil {
		return x.Intro
	}
	return ""
}

func (x *Journal) GetPatientId() int32 {
	if x != nil {
		return x.PatientId
	}
	return 0
}

func (x *Journal) GetCreatedBy() int32 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

type JournalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JournalId int32 `protobuf:"varint,1,opt,name=JournalId,proto3" json:"JournalId,omitempty"`
}

func (x *JournalRequest) Reset() {
	*x = JournalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_journal_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JournalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JournalRequest) ProtoMessage() {}

func (x *JournalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_journal_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JournalRequest.ProtoReflect.Descriptor instead.
func (*JournalRequest) Descriptor() ([]byte, []int) {
	return file_journal_proto_rawDescGZIP(), []int{2}
}

func (x *JournalRequest) GetJournalId() int32 {
	if x != nil {
		return x.JournalId
	}
	return 0
}

type PatientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PatientId int32 `protobuf:"varint,1,opt,name=PatientId,proto3" json:"PatientId,omitempty"`
}

func (x *PatientRequest) Reset() {
	*x = PatientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_journal_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatientRequest) ProtoMessage() {}

func (x *PatientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_journal_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatientRequest.ProtoReflect.Descriptor instead.
func (*PatientRequest) Descriptor() ([]byte, []int) {
	return file_journal_proto_rawDescGZIP(), []int{3}
}

func (x *PatientRequest) GetPatientId() int32 {
	if x != nil {
		return x.PatientId
	}
	return 0
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_journal_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_journal_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_journal_proto_rawDescGZIP(), []int{4}
}

func (x *Status) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type JHealth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *JHealth) Reset() {
	*x = JHealth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_journal_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JHealth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JHealth) ProtoMessage() {}

func (x *JHealth) ProtoReflect() protoreflect.Message {
	mi := &file_journal_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JHealth.ProtoReflect.Descriptor instead.
func (*JHealth) Descriptor() ([]byte, []int) {
	return file_journal_proto_rawDescGZIP(), []int{5}
}

func (x *JHealth) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Journals struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Journals []*Journal `protobuf:"bytes,1,rep,name=journals,proto3" json:"journals,omitempty"`
}

func (x *Journals) Reset() {
	*x = Journals{}
	if protoimpl.UnsafeEnabled {
		mi := &file_journal_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Journals) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Journals) ProtoMessage() {}

func (x *Journals) ProtoReflect() protoreflect.Message {
	mi := &file_journal_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Journals.ProtoReflect.Descriptor instead.
func (*Journals) Descriptor() ([]byte, []int) {
	return file_journal_proto_rawDescGZIP(), []int{6}
}

func (x *Journals) GetJournals() []*Journal {
	if x != nil {
		return x.Journals
	}
	return nil
}

type JournalDocument struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DocumentId   int32         `protobuf:"varint,1,opt,name=DocumentId,proto3" json:"DocumentId,omitempty"`
	Content      string        `protobuf:"bytes,2,opt,name=Content,proto3" json:"Content,omitempty"`
	JournalId    int32         `protobuf:"varint,3,opt,name=JournalId,proto3" json:"JournalId,omitempty"`
	CreatedBy    int32         `protobuf:"varint,4,opt,name=CreatedBy,proto3" json:"CreatedBy,omitempty"`
	Title        string        `protobuf:"bytes,5,opt,name=Title,proto3" json:"Title,omitempty"`
	Summary      string        `protobuf:"bytes,6,opt,name=Summary,proto3" json:"Summary,omitempty"`
	CreationTime string        `protobuf:"bytes,7,opt,name=CreationTime,proto3" json:"CreationTime,omitempty"`
	Attachments  []*Attachment `protobuf:"bytes,8,rep,name=Attachments,proto3" json:"Attachments,omitempty"`
}

func (x *JournalDocument) Reset() {
	*x = JournalDocument{}
	if protoimpl.UnsafeEnabled {
		mi := &file_journal_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JournalDocument) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JournalDocument) ProtoMessage() {}

func (x *JournalDocument) ProtoReflect() protoreflect.Message {
	mi := &file_journal_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JournalDocument.ProtoReflect.Descriptor instead.
func (*JournalDocument) Descriptor() ([]byte, []int) {
	return file_journal_proto_rawDescGZIP(), []int{7}
}

func (x *JournalDocument) GetDocumentId() int32 {
	if x != nil {
		return x.DocumentId
	}
	return 0
}

func (x *JournalDocument) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *JournalDocument) GetJournalId() int32 {
	if x != nil {
		return x.JournalId
	}
	return 0
}

func (x *JournalDocument) GetCreatedBy() int32 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

func (x *JournalDocument) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *JournalDocument) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *JournalDocument) GetCreationTime() string {
	if x != nil {
		return x.CreationTime
	}
	return ""
}

func (x *JournalDocument) GetAttachments() []*Attachment {
	if x != nil {
		return x.Attachments
	}
	return nil
}

type JournalDocuments struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JournalDocuments []*JournalDocument `protobuf:"bytes,1,rep,name=JournalDocuments,proto3" json:"JournalDocuments,omitempty"`
}

func (x *JournalDocuments) Reset() {
	*x = JournalDocuments{}
	if protoimpl.UnsafeEnabled {
		mi := &file_journal_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JournalDocuments) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JournalDocuments) ProtoMessage() {}

func (x *JournalDocuments) ProtoReflect() protoreflect.Message {
	mi := &file_journal_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JournalDocuments.ProtoReflect.Descriptor instead.
func (*JournalDocuments) Descriptor() ([]byte, []int) {
	return file_journal_proto_rawDescGZIP(), []int{8}
}

func (x *JournalDocuments) GetJournalDocuments() []*JournalDocument {
	if x != nil {
		return x.JournalDocuments
	}
	return nil
}

type JournalDocumentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JournalDocumentId int32 `protobuf:"varint,1,opt,name=JournalDocumentId,proto3" json:"JournalDocumentId,omitempty"`
}

func (x *JournalDocumentRequest) Reset() {
	*x = JournalDocumentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_journal_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JournalDocumentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JournalDocumentRequest) ProtoMessage() {}

func (x *JournalDocumentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_journal_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JournalDocumentRequest.ProtoReflect.Descriptor instead.
func (*JournalDocumentRequest) Descriptor() ([]byte, []int) {
	return file_journal_proto_rawDescGZIP(), []int{9}
}

func (x *JournalDocumentRequest) GetJournalDocumentId() int32 {
	if x != nil {
		return x.JournalDocumentId
	}
	return 0
}

type Attachment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AttachmentId int32   `protobuf:"varint,1,opt,name=AttachmentId,proto3" json:"AttachmentId,omitempty"`
	FileName     string  `protobuf:"bytes,2,opt,name=FileName,proto3" json:"FileName,omitempty"`
	FileStoreId  int32   `protobuf:"varint,3,opt,name=FileStoreId,proto3" json:"FileStoreId,omitempty"`
	DocumentId   int32   `protobuf:"varint,4,opt,name=DocumentId,proto3" json:"DocumentId,omitempty"`
	FileTypeId   int32   `protobuf:"varint,5,opt,name=FileTypeId,proto3" json:"FileTypeId,omitempty"`
	FileType     *string `protobuf:"bytes,6,opt,name=FileType,proto3,oneof" json:"FileType,omitempty"`
	Path         *string `protobuf:"bytes,7,opt,name=Path,proto3,oneof" json:"Path,omitempty"`
	Content      *string `protobuf:"bytes,8,opt,name=Content,proto3,oneof" json:"Content,omitempty"`
}

func (x *Attachment) Reset() {
	*x = Attachment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_journal_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attachment) ProtoMessage() {}

func (x *Attachment) ProtoReflect() protoreflect.Message {
	mi := &file_journal_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attachment.ProtoReflect.Descriptor instead.
func (*Attachment) Descriptor() ([]byte, []int) {
	return file_journal_proto_rawDescGZIP(), []int{10}
}

func (x *Attachment) GetAttachmentId() int32 {
	if x != nil {
		return x.AttachmentId
	}
	return 0
}

func (x *Attachment) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *Attachment) GetFileStoreId() int32 {
	if x != nil {
		return x.FileStoreId
	}
	return 0
}

func (x *Attachment) GetDocumentId() int32 {
	if x != nil {
		return x.DocumentId
	}
	return 0
}

func (x *Attachment) GetFileTypeId() int32 {
	if x != nil {
		return x.FileTypeId
	}
	return 0
}

func (x *Attachment) GetFileType() string {
	if x != nil && x.FileType != nil {
		return *x.FileType
	}
	return ""
}

func (x *Attachment) GetPath() string {
	if x != nil && x.Path != nil {
		return *x.Path
	}
	return ""
}

func (x *Attachment) GetContent() string {
	if x != nil && x.Content != nil {
		return *x.Content
	}
	return ""
}

var File_journal_proto protoreflect.FileDescriptor

var file_journal_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x08, 0x0a, 0x06, 0x4a, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x9d, 0x01, 0x0a, 0x07, 0x4a, 0x6f,
	0x75, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61,
	0x6c, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6e, 0x74, 0x72, 0x6f,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x12, 0x1c, 0x0a,
	0x09, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x22, 0x2e, 0x0a, 0x0e, 0x4a, 0x6f, 0x75,
	0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x4a,
	0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x0e, 0x50, 0x61, 0x74,
	0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x50,
	0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x22, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x23, 0x0a,
	0x07, 0x4a, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x30, 0x0a, 0x08, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x12, 0x24,
	0x0a, 0x08, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x08, 0x2e, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x08, 0x6a, 0x6f, 0x75, 0x72,
	0x6e, 0x61, 0x6c, 0x73, 0x22, 0x8a, 0x02, 0x0a, 0x0f, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c,
	0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x44, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x22,
	0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x0b, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x22, 0x50, 0x0a, 0x10, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x3c, 0x0a, 0x10, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c,
	0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x10, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x22, 0x46, 0x0a, 0x16, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a,
	0x11, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61,
	0x6c, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0xa9, 0x02, 0x0a, 0x0a,
	0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x41, 0x74,
	0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0c, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x46, 0x69,
	0x6c, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a,
	0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a,
	0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x08,
	0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a,
	0x04, 0x50, 0x61, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x50,
	0x61, 0x74, 0x68, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x50, 0x61, 0x74, 0x68, 0x42, 0x0a, 0x0a, 0x08, 0x5f,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x32, 0xd2, 0x04, 0x0a, 0x0e, 0x4a, 0x6f, 0x75, 0x72,
	0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x29, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x0f, 0x2e, 0x4a, 0x6f, 0x75, 0x72, 0x6e,
	0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e, 0x4a, 0x6f, 0x75, 0x72,
	0x6e, 0x61, 0x6c, 0x22, 0x00, 0x12, 0x25, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a,
	0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x08, 0x2e, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c,
	0x1a, 0x08, 0x2e, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x0d,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x0f, 0x2e,
	0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x20, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x07, 0x2e, 0x4a, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x08, 0x2e, 0x4a, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x22, 0x00, 0x12, 0x25, 0x0a, 0x0d, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x08, 0x2e, 0x4a,
	0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x1a, 0x08, 0x2e, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c,
	0x22, 0x00, 0x12, 0x34, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c,
	0x73, 0x42, 0x79, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x0f, 0x2e, 0x50, 0x61, 0x74,
	0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x4a, 0x6f,
	0x75, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x17, 0x2e, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a,
	0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x10,
	0x2e, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x1a, 0x10, 0x2e, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x75, 0x72, 0x6e,
	0x61, 0x6c, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x4a, 0x6f, 0x75,
	0x72, 0x6e, 0x61, 0x6c, 0x12, 0x0f, 0x2e, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x44,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x17, 0x2e, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x4a, 0x6f, 0x75, 0x72,
	0x6e, 0x61, 0x6c, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x3d, 0x0a,
	0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x2e, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c,
	0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x10, 0x2e, 0x4a, 0x6f, 0x75, 0x72, 0x6e,
	0x61, 0x6c, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09,
	0x2e, 0x3b, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_journal_proto_rawDescOnce sync.Once
	file_journal_proto_rawDescData = file_journal_proto_rawDesc
)

func file_journal_proto_rawDescGZIP() []byte {
	file_journal_proto_rawDescOnce.Do(func() {
		file_journal_proto_rawDescData = protoimpl.X.CompressGZIP(file_journal_proto_rawDescData)
	})
	return file_journal_proto_rawDescData
}

var file_journal_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_journal_proto_goTypes = []interface{}{
	(*JEmpty)(nil),                 // 0: JEmpty
	(*Journal)(nil),                // 1: Journal
	(*JournalRequest)(nil),         // 2: JournalRequest
	(*PatientRequest)(nil),         // 3: PatientRequest
	(*Status)(nil),                 // 4: Status
	(*JHealth)(nil),                // 5: JHealth
	(*Journals)(nil),               // 6: Journals
	(*JournalDocument)(nil),        // 7: JournalDocument
	(*JournalDocuments)(nil),       // 8: JournalDocuments
	(*JournalDocumentRequest)(nil), // 9: JournalDocumentRequest
	(*Attachment)(nil),             // 10: Attachment
}
var file_journal_proto_depIdxs = []int32{
	1,  // 0: Journals.journals:type_name -> Journal
	10, // 1: JournalDocument.Attachments:type_name -> Attachment
	7,  // 2: JournalDocuments.JournalDocuments:type_name -> JournalDocument
	2,  // 3: JournalService.GetJournal:input_type -> JournalRequest
	1,  // 4: JournalService.CreateJournal:input_type -> Journal
	2,  // 5: JournalService.DeleteJournal:input_type -> JournalRequest
	0,  // 6: JournalService.GetHealth:input_type -> JEmpty
	1,  // 7: JournalService.UpdateJournal:input_type -> Journal
	3,  // 8: JournalService.GetJournalsByPatient:input_type -> PatientRequest
	9,  // 9: JournalService.DeleteJournalDocument:input_type -> JournalDocumentRequest
	7,  // 10: JournalService.UpdateJournalDocument:input_type -> JournalDocument
	2,  // 11: JournalService.GetJournalDocumentsByJournal:input_type -> JournalRequest
	9,  // 12: JournalService.GetJournalDocument:input_type -> JournalDocumentRequest
	7,  // 13: JournalService.CreateJournalDocument:input_type -> JournalDocument
	1,  // 14: JournalService.GetJournal:output_type -> Journal
	1,  // 15: JournalService.CreateJournal:output_type -> Journal
	4,  // 16: JournalService.DeleteJournal:output_type -> Status
	5,  // 17: JournalService.GetHealth:output_type -> JHealth
	1,  // 18: JournalService.UpdateJournal:output_type -> Journal
	6,  // 19: JournalService.GetJournalsByPatient:output_type -> Journals
	4,  // 20: JournalService.DeleteJournalDocument:output_type -> Status
	7,  // 21: JournalService.UpdateJournalDocument:output_type -> JournalDocument
	8,  // 22: JournalService.GetJournalDocumentsByJournal:output_type -> JournalDocuments
	7,  // 23: JournalService.GetJournalDocument:output_type -> JournalDocument
	7,  // 24: JournalService.CreateJournalDocument:output_type -> JournalDocument
	14, // [14:25] is the sub-list for method output_type
	3,  // [3:14] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_journal_proto_init() }
func file_journal_proto_init() {
	if File_journal_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_journal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JEmpty); i {
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
		file_journal_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Journal); i {
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
		file_journal_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JournalRequest); i {
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
		file_journal_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatientRequest); i {
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
		file_journal_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_journal_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JHealth); i {
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
		file_journal_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Journals); i {
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
		file_journal_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JournalDocument); i {
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
		file_journal_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JournalDocuments); i {
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
		file_journal_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JournalDocumentRequest); i {
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
		file_journal_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attachment); i {
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
	file_journal_proto_msgTypes[10].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_journal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_journal_proto_goTypes,
		DependencyIndexes: file_journal_proto_depIdxs,
		MessageInfos:      file_journal_proto_msgTypes,
	}.Build()
	File_journal_proto = out.File
	file_journal_proto_rawDesc = nil
	file_journal_proto_goTypes = nil
	file_journal_proto_depIdxs = nil
}
