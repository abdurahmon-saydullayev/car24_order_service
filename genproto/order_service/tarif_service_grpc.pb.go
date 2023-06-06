// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package order_service

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TarifServiceClient is the client API for TarifService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TarifServiceClient interface {
	Create(ctx context.Context, in *CreateTarif, opts ...grpc.CallOption) (*Tarif, error)
	GetByID(ctx context.Context, in *TarifPK, opts ...grpc.CallOption) (*Tarif, error)
	Delete(ctx context.Context, in *TarifPK, opts ...grpc.CallOption) (*empty.Empty, error)
}

type tarifServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTarifServiceClient(cc grpc.ClientConnInterface) TarifServiceClient {
	return &tarifServiceClient{cc}
}

func (c *tarifServiceClient) Create(ctx context.Context, in *CreateTarif, opts ...grpc.CallOption) (*Tarif, error) {
	out := new(Tarif)
	err := c.cc.Invoke(ctx, "/order_service.TarifService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tarifServiceClient) GetByID(ctx context.Context, in *TarifPK, opts ...grpc.CallOption) (*Tarif, error) {
	out := new(Tarif)
	err := c.cc.Invoke(ctx, "/order_service.TarifService/GetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tarifServiceClient) Delete(ctx context.Context, in *TarifPK, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/order_service.TarifService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TarifServiceServer is the server API for TarifService service.
// All implementations must embed UnimplementedTarifServiceServer
// for forward compatibility
type TarifServiceServer interface {
	Create(context.Context, *CreateTarif) (*Tarif, error)
	GetByID(context.Context, *TarifPK) (*Tarif, error)
	Delete(context.Context, *TarifPK) (*empty.Empty, error)
	mustEmbedUnimplementedTarifServiceServer()
}

// UnimplementedTarifServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTarifServiceServer struct {
}

func (UnimplementedTarifServiceServer) Create(context.Context, *CreateTarif) (*Tarif, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTarifServiceServer) GetByID(context.Context, *TarifPK) (*Tarif, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedTarifServiceServer) Delete(context.Context, *TarifPK) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedTarifServiceServer) mustEmbedUnimplementedTarifServiceServer() {}

// UnsafeTarifServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TarifServiceServer will
// result in compilation errors.
type UnsafeTarifServiceServer interface {
	mustEmbedUnimplementedTarifServiceServer()
}

func RegisterTarifServiceServer(s grpc.ServiceRegistrar, srv TarifServiceServer) {
	s.RegisterService(&TarifService_ServiceDesc, srv)
}

func _TarifService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTarif)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TarifServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order_service.TarifService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TarifServiceServer).Create(ctx, req.(*CreateTarif))
	}
	return interceptor(ctx, in, info, handler)
}

func _TarifService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TarifPK)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TarifServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order_service.TarifService/GetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TarifServiceServer).GetByID(ctx, req.(*TarifPK))
	}
	return interceptor(ctx, in, info, handler)
}

func _TarifService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TarifPK)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TarifServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order_service.TarifService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TarifServiceServer).Delete(ctx, req.(*TarifPK))
	}
	return interceptor(ctx, in, info, handler)
}

// TarifService_ServiceDesc is the grpc.ServiceDesc for TarifService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TarifService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order_service.TarifService",
	HandlerType: (*TarifServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _TarifService_Create_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _TarifService_GetByID_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TarifService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tarif_service.proto",
}
