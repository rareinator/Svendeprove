// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package useradmin

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

// UseradminServiceClient is the client API for UseradminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UseradminServiceClient interface {
	GetHealth(ctx context.Context, in *UAEmpty, opts ...grpc.CallOption) (*UAHealth, error)
	GetEmployee(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UAUser, error)
	GetPatients(ctx context.Context, in *UAEmpty, opts ...grpc.CallOption) (*Users, error)
	GetDepartments(ctx context.Context, in *UAEmpty, opts ...grpc.CallOption) (*Departments, error)
	GetBeds(ctx context.Context, in *UAEmpty, opts ...grpc.CallOption) (*Beds, error)
	GetAvailableBeds(ctx context.Context, in *BedsRequest, opts ...grpc.CallOption) (*Beds, error)
	GetHospitals(ctx context.Context, in *UAEmpty, opts ...grpc.CallOption) (*Hospitals, error)
}

type useradminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUseradminServiceClient(cc grpc.ClientConnInterface) UseradminServiceClient {
	return &useradminServiceClient{cc}
}

func (c *useradminServiceClient) GetHealth(ctx context.Context, in *UAEmpty, opts ...grpc.CallOption) (*UAHealth, error) {
	out := new(UAHealth)
	err := c.cc.Invoke(ctx, "/UseradminService/GetHealth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *useradminServiceClient) GetEmployee(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UAUser, error) {
	out := new(UAUser)
	err := c.cc.Invoke(ctx, "/UseradminService/GetEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *useradminServiceClient) GetPatients(ctx context.Context, in *UAEmpty, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := c.cc.Invoke(ctx, "/UseradminService/GetPatients", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *useradminServiceClient) GetDepartments(ctx context.Context, in *UAEmpty, opts ...grpc.CallOption) (*Departments, error) {
	out := new(Departments)
	err := c.cc.Invoke(ctx, "/UseradminService/GetDepartments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *useradminServiceClient) GetBeds(ctx context.Context, in *UAEmpty, opts ...grpc.CallOption) (*Beds, error) {
	out := new(Beds)
	err := c.cc.Invoke(ctx, "/UseradminService/GetBeds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *useradminServiceClient) GetAvailableBeds(ctx context.Context, in *BedsRequest, opts ...grpc.CallOption) (*Beds, error) {
	out := new(Beds)
	err := c.cc.Invoke(ctx, "/UseradminService/GetAvailableBeds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *useradminServiceClient) GetHospitals(ctx context.Context, in *UAEmpty, opts ...grpc.CallOption) (*Hospitals, error) {
	out := new(Hospitals)
	err := c.cc.Invoke(ctx, "/UseradminService/GetHospitals", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UseradminServiceServer is the server API for UseradminService service.
// All implementations must embed UnimplementedUseradminServiceServer
// for forward compatibility
type UseradminServiceServer interface {
	GetHealth(context.Context, *UAEmpty) (*UAHealth, error)
	GetEmployee(context.Context, *UserRequest) (*UAUser, error)
	GetPatients(context.Context, *UAEmpty) (*Users, error)
	GetDepartments(context.Context, *UAEmpty) (*Departments, error)
	GetBeds(context.Context, *UAEmpty) (*Beds, error)
	GetAvailableBeds(context.Context, *BedsRequest) (*Beds, error)
	GetHospitals(context.Context, *UAEmpty) (*Hospitals, error)
	mustEmbedUnimplementedUseradminServiceServer()
}

// UnimplementedUseradminServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUseradminServiceServer struct {
}

func (UnimplementedUseradminServiceServer) GetHealth(context.Context, *UAEmpty) (*UAHealth, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHealth not implemented")
}
func (UnimplementedUseradminServiceServer) GetEmployee(context.Context, *UserRequest) (*UAUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmployee not implemented")
}
func (UnimplementedUseradminServiceServer) GetPatients(context.Context, *UAEmpty) (*Users, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPatients not implemented")
}
func (UnimplementedUseradminServiceServer) GetDepartments(context.Context, *UAEmpty) (*Departments, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDepartments not implemented")
}
func (UnimplementedUseradminServiceServer) GetBeds(context.Context, *UAEmpty) (*Beds, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBeds not implemented")
}
func (UnimplementedUseradminServiceServer) GetAvailableBeds(context.Context, *BedsRequest) (*Beds, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvailableBeds not implemented")
}
func (UnimplementedUseradminServiceServer) GetHospitals(context.Context, *UAEmpty) (*Hospitals, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHospitals not implemented")
}
func (UnimplementedUseradminServiceServer) mustEmbedUnimplementedUseradminServiceServer() {}

// UnsafeUseradminServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UseradminServiceServer will
// result in compilation errors.
type UnsafeUseradminServiceServer interface {
	mustEmbedUnimplementedUseradminServiceServer()
}

func RegisterUseradminServiceServer(s grpc.ServiceRegistrar, srv UseradminServiceServer) {
	s.RegisterService(&UseradminService_ServiceDesc, srv)
}

func _UseradminService_GetHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UAEmpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UseradminServiceServer).GetHealth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UseradminService/GetHealth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UseradminServiceServer).GetHealth(ctx, req.(*UAEmpty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UseradminService_GetEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UseradminServiceServer).GetEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UseradminService/GetEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UseradminServiceServer).GetEmployee(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UseradminService_GetPatients_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UAEmpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UseradminServiceServer).GetPatients(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UseradminService/GetPatients",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UseradminServiceServer).GetPatients(ctx, req.(*UAEmpty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UseradminService_GetDepartments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UAEmpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UseradminServiceServer).GetDepartments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UseradminService/GetDepartments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UseradminServiceServer).GetDepartments(ctx, req.(*UAEmpty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UseradminService_GetBeds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UAEmpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UseradminServiceServer).GetBeds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UseradminService/GetBeds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UseradminServiceServer).GetBeds(ctx, req.(*UAEmpty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UseradminService_GetAvailableBeds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BedsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UseradminServiceServer).GetAvailableBeds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UseradminService/GetAvailableBeds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UseradminServiceServer).GetAvailableBeds(ctx, req.(*BedsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UseradminService_GetHospitals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UAEmpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UseradminServiceServer).GetHospitals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UseradminService/GetHospitals",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UseradminServiceServer).GetHospitals(ctx, req.(*UAEmpty))
	}
	return interceptor(ctx, in, info, handler)
}

// UseradminService_ServiceDesc is the grpc.ServiceDesc for UseradminService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UseradminService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "UseradminService",
	HandlerType: (*UseradminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHealth",
			Handler:    _UseradminService_GetHealth_Handler,
		},
		{
			MethodName: "GetEmployee",
			Handler:    _UseradminService_GetEmployee_Handler,
		},
		{
			MethodName: "GetPatients",
			Handler:    _UseradminService_GetPatients_Handler,
		},
		{
			MethodName: "GetDepartments",
			Handler:    _UseradminService_GetDepartments_Handler,
		},
		{
			MethodName: "GetBeds",
			Handler:    _UseradminService_GetBeds_Handler,
		},
		{
			MethodName: "GetAvailableBeds",
			Handler:    _UseradminService_GetAvailableBeds_Handler,
		},
		{
			MethodName: "GetHospitals",
			Handler:    _UseradminService_GetHospitals_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "useradmin.proto",
}
