// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package booking

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

// BookingServiceClient is the client API for BookingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookingServiceClient interface {
	GetHealth(ctx context.Context, in *BEmpty, opts ...grpc.CallOption) (*BHealth, error)
	CreateBooking(ctx context.Context, in *Booking, opts ...grpc.CallOption) (*Booking, error)
	GetBooking(ctx context.Context, in *BRequest, opts ...grpc.CallOption) (*Booking, error)
	UpdateBooking(ctx context.Context, in *Booking, opts ...grpc.CallOption) (*Booking, error)
	DeleteBooking(ctx context.Context, in *BRequest, opts ...grpc.CallOption) (*BStatus, error)
	GetBookingsByPatient(ctx context.Context, in *BRequest, opts ...grpc.CallOption) (*Bookings, error)
	GetBookingsByCreatingEmployee(ctx context.Context, in *BRequest, opts ...grpc.CallOption) (*Bookings, error)
	GetBookingsByInTimeFrame(ctx context.Context, in *BTimeFrameRequest, opts ...grpc.CallOption) (*Bookings, error)
}

type bookingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookingServiceClient(cc grpc.ClientConnInterface) BookingServiceClient {
	return &bookingServiceClient{cc}
}

func (c *bookingServiceClient) GetHealth(ctx context.Context, in *BEmpty, opts ...grpc.CallOption) (*BHealth, error) {
	out := new(BHealth)
	err := c.cc.Invoke(ctx, "/BookingService/GetHealth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) CreateBooking(ctx context.Context, in *Booking, opts ...grpc.CallOption) (*Booking, error) {
	out := new(Booking)
	err := c.cc.Invoke(ctx, "/BookingService/CreateBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) GetBooking(ctx context.Context, in *BRequest, opts ...grpc.CallOption) (*Booking, error) {
	out := new(Booking)
	err := c.cc.Invoke(ctx, "/BookingService/GetBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) UpdateBooking(ctx context.Context, in *Booking, opts ...grpc.CallOption) (*Booking, error) {
	out := new(Booking)
	err := c.cc.Invoke(ctx, "/BookingService/UpdateBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) DeleteBooking(ctx context.Context, in *BRequest, opts ...grpc.CallOption) (*BStatus, error) {
	out := new(BStatus)
	err := c.cc.Invoke(ctx, "/BookingService/DeleteBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) GetBookingsByPatient(ctx context.Context, in *BRequest, opts ...grpc.CallOption) (*Bookings, error) {
	out := new(Bookings)
	err := c.cc.Invoke(ctx, "/BookingService/GetBookingsByPatient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) GetBookingsByCreatingEmployee(ctx context.Context, in *BRequest, opts ...grpc.CallOption) (*Bookings, error) {
	out := new(Bookings)
	err := c.cc.Invoke(ctx, "/BookingService/GetBookingsByCreatingEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) GetBookingsByInTimeFrame(ctx context.Context, in *BTimeFrameRequest, opts ...grpc.CallOption) (*Bookings, error) {
	out := new(Bookings)
	err := c.cc.Invoke(ctx, "/BookingService/GetBookingsByInTimeFrame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookingServiceServer is the server API for BookingService service.
// All implementations must embed UnimplementedBookingServiceServer
// for forward compatibility
type BookingServiceServer interface {
	GetHealth(context.Context, *BEmpty) (*BHealth, error)
	CreateBooking(context.Context, *Booking) (*Booking, error)
	GetBooking(context.Context, *BRequest) (*Booking, error)
	UpdateBooking(context.Context, *Booking) (*Booking, error)
	DeleteBooking(context.Context, *BRequest) (*BStatus, error)
	GetBookingsByPatient(context.Context, *BRequest) (*Bookings, error)
	GetBookingsByCreatingEmployee(context.Context, *BRequest) (*Bookings, error)
	GetBookingsByInTimeFrame(context.Context, *BTimeFrameRequest) (*Bookings, error)
	mustEmbedUnimplementedBookingServiceServer()
}

// UnimplementedBookingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBookingServiceServer struct {
}

func (UnimplementedBookingServiceServer) GetHealth(context.Context, *BEmpty) (*BHealth, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHealth not implemented")
}
func (UnimplementedBookingServiceServer) CreateBooking(context.Context, *Booking) (*Booking, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBooking not implemented")
}
func (UnimplementedBookingServiceServer) GetBooking(context.Context, *BRequest) (*Booking, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBooking not implemented")
}
func (UnimplementedBookingServiceServer) UpdateBooking(context.Context, *Booking) (*Booking, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBooking not implemented")
}
func (UnimplementedBookingServiceServer) DeleteBooking(context.Context, *BRequest) (*BStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBooking not implemented")
}
func (UnimplementedBookingServiceServer) GetBookingsByPatient(context.Context, *BRequest) (*Bookings, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookingsByPatient not implemented")
}
func (UnimplementedBookingServiceServer) GetBookingsByCreatingEmployee(context.Context, *BRequest) (*Bookings, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookingsByCreatingEmployee not implemented")
}
func (UnimplementedBookingServiceServer) GetBookingsByInTimeFrame(context.Context, *BTimeFrameRequest) (*Bookings, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookingsByInTimeFrame not implemented")
}
func (UnimplementedBookingServiceServer) mustEmbedUnimplementedBookingServiceServer() {}

// UnsafeBookingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookingServiceServer will
// result in compilation errors.
type UnsafeBookingServiceServer interface {
	mustEmbedUnimplementedBookingServiceServer()
}

func RegisterBookingServiceServer(s grpc.ServiceRegistrar, srv BookingServiceServer) {
	s.RegisterService(&BookingService_ServiceDesc, srv)
}

func _BookingService_GetHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BEmpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).GetHealth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BookingService/GetHealth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).GetHealth(ctx, req.(*BEmpty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_CreateBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Booking)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).CreateBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BookingService/CreateBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).CreateBooking(ctx, req.(*Booking))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_GetBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).GetBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BookingService/GetBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).GetBooking(ctx, req.(*BRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_UpdateBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Booking)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).UpdateBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BookingService/UpdateBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).UpdateBooking(ctx, req.(*Booking))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_DeleteBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).DeleteBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BookingService/DeleteBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).DeleteBooking(ctx, req.(*BRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_GetBookingsByPatient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).GetBookingsByPatient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BookingService/GetBookingsByPatient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).GetBookingsByPatient(ctx, req.(*BRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_GetBookingsByCreatingEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).GetBookingsByCreatingEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BookingService/GetBookingsByCreatingEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).GetBookingsByCreatingEmployee(ctx, req.(*BRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_GetBookingsByInTimeFrame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BTimeFrameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).GetBookingsByInTimeFrame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BookingService/GetBookingsByInTimeFrame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).GetBookingsByInTimeFrame(ctx, req.(*BTimeFrameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BookingService_ServiceDesc is the grpc.ServiceDesc for BookingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BookingService",
	HandlerType: (*BookingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHealth",
			Handler:    _BookingService_GetHealth_Handler,
		},
		{
			MethodName: "CreateBooking",
			Handler:    _BookingService_CreateBooking_Handler,
		},
		{
			MethodName: "GetBooking",
			Handler:    _BookingService_GetBooking_Handler,
		},
		{
			MethodName: "UpdateBooking",
			Handler:    _BookingService_UpdateBooking_Handler,
		},
		{
			MethodName: "DeleteBooking",
			Handler:    _BookingService_DeleteBooking_Handler,
		},
		{
			MethodName: "GetBookingsByPatient",
			Handler:    _BookingService_GetBookingsByPatient_Handler,
		},
		{
			MethodName: "GetBookingsByCreatingEmployee",
			Handler:    _BookingService_GetBookingsByCreatingEmployee_Handler,
		},
		{
			MethodName: "GetBookingsByInTimeFrame",
			Handler:    _BookingService_GetBookingsByInTimeFrame_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "booking.proto",
}