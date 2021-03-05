// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package journal

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// JournalServiceClient is the client API for JournalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JournalServiceClient interface {
	GetJournal(ctx context.Context, in *JournalRequest, opts ...grpc.CallOption) (*Journal, error)
	CreateJournal(ctx context.Context, in *Journal, opts ...grpc.CallOption) (*Journal, error)
	DeleteJournal(ctx context.Context, in *JournalRequest, opts ...grpc.CallOption) (*Status, error)
	GetHealth(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Health, error)
	UpdateJournal(ctx context.Context, in *Journal, opts ...grpc.CallOption) (*Journal, error)
	GetJournalsByPatient(ctx context.Context, in *PatientRequest, opts ...grpc.CallOption) (*Journals, error)
}

type journalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJournalServiceClient(cc grpc.ClientConnInterface) JournalServiceClient {
	return &journalServiceClient{cc}
}

func (c *journalServiceClient) GetJournal(ctx context.Context, in *JournalRequest, opts ...grpc.CallOption) (*Journal, error) {
	out := new(Journal)
	err := c.cc.Invoke(ctx, "/JournalService/GetJournal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *journalServiceClient) CreateJournal(ctx context.Context, in *Journal, opts ...grpc.CallOption) (*Journal, error) {
	out := new(Journal)
	err := c.cc.Invoke(ctx, "/JournalService/CreateJournal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *journalServiceClient) DeleteJournal(ctx context.Context, in *JournalRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/JournalService/DeleteJournal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *journalServiceClient) GetHealth(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Health, error) {
	out := new(Health)
	err := c.cc.Invoke(ctx, "/JournalService/GetHealth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *journalServiceClient) UpdateJournal(ctx context.Context, in *Journal, opts ...grpc.CallOption) (*Journal, error) {
	out := new(Journal)
	err := c.cc.Invoke(ctx, "/JournalService/UpdateJournal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *journalServiceClient) GetJournalsByPatient(ctx context.Context, in *PatientRequest, opts ...grpc.CallOption) (*Journals, error) {
	out := new(Journals)
	err := c.cc.Invoke(ctx, "/JournalService/GetJournalsByPatient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JournalServiceServer is the server API for JournalService service.
// All implementations must embed UnimplementedJournalServiceServer
// for forward compatibility
type JournalServiceServer interface {
	GetJournal(context.Context, *JournalRequest) (*Journal, error)
	CreateJournal(context.Context, *Journal) (*Journal, error)
	DeleteJournal(context.Context, *JournalRequest) (*Status, error)
	GetHealth(context.Context, *Empty) (*Health, error)
	UpdateJournal(context.Context, *Journal) (*Journal, error)
	GetJournalsByPatient(context.Context, *PatientRequest) (*Journals, error)
	mustEmbedUnimplementedJournalServiceServer()
}

// UnimplementedJournalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedJournalServiceServer struct {
}

func (UnimplementedJournalServiceServer) GetJournal(context.Context, *JournalRequest) (*Journal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJournal not implemented")
}
func (UnimplementedJournalServiceServer) CreateJournal(context.Context, *Journal) (*Journal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateJournal not implemented")
}
func (UnimplementedJournalServiceServer) DeleteJournal(context.Context, *JournalRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteJournal not implemented")
}
func (UnimplementedJournalServiceServer) GetHealth(context.Context, *Empty) (*Health, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHealth not implemented")
}
func (UnimplementedJournalServiceServer) UpdateJournal(context.Context, *Journal) (*Journal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateJournal not implemented")
}
func (UnimplementedJournalServiceServer) GetJournalsByPatient(context.Context, *PatientRequest) (*Journals, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJournalsByPatient not implemented")
}
func (UnimplementedJournalServiceServer) mustEmbedUnimplementedJournalServiceServer() {}

// UnsafeJournalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JournalServiceServer will
// result in compilation errors.
type UnsafeJournalServiceServer interface {
	mustEmbedUnimplementedJournalServiceServer()
}

func RegisterJournalServiceServer(s grpc.ServiceRegistrar, srv JournalServiceServer) {
	s.RegisterService(&JournalService_ServiceDesc, srv)
}

func _JournalService_GetJournal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JournalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JournalServiceServer).GetJournal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/JournalService/GetJournal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JournalServiceServer).GetJournal(ctx, req.(*JournalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JournalService_CreateJournal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Journal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JournalServiceServer).CreateJournal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/JournalService/CreateJournal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JournalServiceServer).CreateJournal(ctx, req.(*Journal))
	}
	return interceptor(ctx, in, info, handler)
}

func _JournalService_DeleteJournal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JournalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JournalServiceServer).DeleteJournal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/JournalService/DeleteJournal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JournalServiceServer).DeleteJournal(ctx, req.(*JournalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JournalService_GetHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JournalServiceServer).GetHealth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/JournalService/GetHealth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JournalServiceServer).GetHealth(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _JournalService_UpdateJournal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Journal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JournalServiceServer).UpdateJournal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/JournalService/UpdateJournal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JournalServiceServer).UpdateJournal(ctx, req.(*Journal))
	}
	return interceptor(ctx, in, info, handler)
}

func _JournalService_GetJournalsByPatient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JournalServiceServer).GetJournalsByPatient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/JournalService/GetJournalsByPatient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JournalServiceServer).GetJournalsByPatient(ctx, req.(*PatientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// JournalService_ServiceDesc is the grpc.ServiceDesc for JournalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JournalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "JournalService",
	HandlerType: (*JournalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetJournal",
			Handler:    _JournalService_GetJournal_Handler,
		},
		{
			MethodName: "CreateJournal",
			Handler:    _JournalService_CreateJournal_Handler,
		},
		{
			MethodName: "DeleteJournal",
			Handler:    _JournalService_DeleteJournal_Handler,
		},
		{
			MethodName: "GetHealth",
			Handler:    _JournalService_GetHealth_Handler,
		},
		{
			MethodName: "UpdateJournal",
			Handler:    _JournalService_UpdateJournal_Handler,
		},
		{
			MethodName: "GetJournalsByPatient",
			Handler:    _JournalService_GetJournalsByPatient_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "journal.proto",
}
