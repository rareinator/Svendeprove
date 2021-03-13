// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package iot

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

// IotServiceClient is the client API for IotService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IotServiceClient interface {
	GetHealth(ctx context.Context, in *IOTEmpty, opts ...grpc.CallOption) (*IOTHealth, error)
	UploadData(ctx context.Context, in *IOTData, opts ...grpc.CallOption) (*IOTData, error)
	ReadData(ctx context.Context, in *IOTRequest, opts ...grpc.CallOption) (*IOTDatas, error)
	ReadDataInTimeFrame(ctx context.Context, in *IOTTimeframeRequest, opts ...grpc.CallOption) (*IOTDatas, error)
}

type iotServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIotServiceClient(cc grpc.ClientConnInterface) IotServiceClient {
	return &iotServiceClient{cc}
}

func (c *iotServiceClient) GetHealth(ctx context.Context, in *IOTEmpty, opts ...grpc.CallOption) (*IOTHealth, error) {
	out := new(IOTHealth)
	err := c.cc.Invoke(ctx, "/IotService/GetHealth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotServiceClient) UploadData(ctx context.Context, in *IOTData, opts ...grpc.CallOption) (*IOTData, error) {
	out := new(IOTData)
	err := c.cc.Invoke(ctx, "/IotService/UploadData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotServiceClient) ReadData(ctx context.Context, in *IOTRequest, opts ...grpc.CallOption) (*IOTDatas, error) {
	out := new(IOTDatas)
	err := c.cc.Invoke(ctx, "/IotService/ReadData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotServiceClient) ReadDataInTimeFrame(ctx context.Context, in *IOTTimeframeRequest, opts ...grpc.CallOption) (*IOTDatas, error) {
	out := new(IOTDatas)
	err := c.cc.Invoke(ctx, "/IotService/ReadDataInTimeFrame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IotServiceServer is the server API for IotService service.
// All implementations must embed UnimplementedIotServiceServer
// for forward compatibility
type IotServiceServer interface {
	GetHealth(context.Context, *IOTEmpty) (*IOTHealth, error)
	UploadData(context.Context, *IOTData) (*IOTData, error)
	ReadData(context.Context, *IOTRequest) (*IOTDatas, error)
	ReadDataInTimeFrame(context.Context, *IOTTimeframeRequest) (*IOTDatas, error)
	mustEmbedUnimplementedIotServiceServer()
}

// UnimplementedIotServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIotServiceServer struct {
}

func (UnimplementedIotServiceServer) GetHealth(context.Context, *IOTEmpty) (*IOTHealth, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHealth not implemented")
}
func (UnimplementedIotServiceServer) UploadData(context.Context, *IOTData) (*IOTData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadData not implemented")
}
func (UnimplementedIotServiceServer) ReadData(context.Context, *IOTRequest) (*IOTDatas, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadData not implemented")
}
func (UnimplementedIotServiceServer) ReadDataInTimeFrame(context.Context, *IOTTimeframeRequest) (*IOTDatas, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadDataInTimeFrame not implemented")
}
func (UnimplementedIotServiceServer) mustEmbedUnimplementedIotServiceServer() {}

// UnsafeIotServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IotServiceServer will
// result in compilation errors.
type UnsafeIotServiceServer interface {
	mustEmbedUnimplementedIotServiceServer()
}

func RegisterIotServiceServer(s grpc.ServiceRegistrar, srv IotServiceServer) {
	s.RegisterService(&IotService_ServiceDesc, srv)
}

func _IotService_GetHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IOTEmpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IotServiceServer).GetHealth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/IotService/GetHealth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IotServiceServer).GetHealth(ctx, req.(*IOTEmpty))
	}
	return interceptor(ctx, in, info, handler)
}

func _IotService_UploadData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IOTData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IotServiceServer).UploadData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/IotService/UploadData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IotServiceServer).UploadData(ctx, req.(*IOTData))
	}
	return interceptor(ctx, in, info, handler)
}

func _IotService_ReadData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IOTRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IotServiceServer).ReadData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/IotService/ReadData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IotServiceServer).ReadData(ctx, req.(*IOTRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IotService_ReadDataInTimeFrame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IOTTimeframeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IotServiceServer).ReadDataInTimeFrame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/IotService/ReadDataInTimeFrame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IotServiceServer).ReadDataInTimeFrame(ctx, req.(*IOTTimeframeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IotService_ServiceDesc is the grpc.ServiceDesc for IotService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IotService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "IotService",
	HandlerType: (*IotServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHealth",
			Handler:    _IotService_GetHealth_Handler,
		},
		{
			MethodName: "UploadData",
			Handler:    _IotService_UploadData_Handler,
		},
		{
			MethodName: "ReadData",
			Handler:    _IotService_ReadData_Handler,
		},
		{
			MethodName: "ReadDataInTimeFrame",
			Handler:    _IotService_ReadDataInTimeFrame_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "iot.proto",
}
