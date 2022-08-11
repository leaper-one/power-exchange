// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: mixin_user.proto

package mixinuser

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

// MixinUserClient is the client API for MixinUser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MixinUserClient interface {
	// 创建 MixinUser
	CreateMixinUser(ctx context.Context, in *CreateMixinUserRequest, opts ...grpc.CallOption) (*CreateMixinUserResponse, error)
}

type mixinUserClient struct {
	cc grpc.ClientConnInterface
}

func NewMixinUserClient(cc grpc.ClientConnInterface) MixinUserClient {
	return &mixinUserClient{cc}
}

func (c *mixinUserClient) CreateMixinUser(ctx context.Context, in *CreateMixinUserRequest, opts ...grpc.CallOption) (*CreateMixinUserResponse, error) {
	out := new(CreateMixinUserResponse)
	err := c.cc.Invoke(ctx, "/mixin_user.MixinUser/CreateMixinUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MixinUserServer is the server API for MixinUser service.
// All implementations must embed UnimplementedMixinUserServer
// for forward compatibility
type MixinUserServer interface {
	// 创建 MixinUser
	CreateMixinUser(context.Context, *CreateMixinUserRequest) (*CreateMixinUserResponse, error)
	mustEmbedUnimplementedMixinUserServer()
}

// UnimplementedMixinUserServer must be embedded to have forward compatible implementations.
type UnimplementedMixinUserServer struct {
}

func (UnimplementedMixinUserServer) CreateMixinUser(context.Context, *CreateMixinUserRequest) (*CreateMixinUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMixinUser not implemented")
}
func (UnimplementedMixinUserServer) mustEmbedUnimplementedMixinUserServer() {}

// UnsafeMixinUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MixinUserServer will
// result in compilation errors.
type UnsafeMixinUserServer interface {
	mustEmbedUnimplementedMixinUserServer()
}

func RegisterMixinUserServer(s grpc.ServiceRegistrar, srv MixinUserServer) {
	s.RegisterService(&MixinUser_ServiceDesc, srv)
}

func _MixinUser_CreateMixinUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMixinUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MixinUserServer).CreateMixinUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mixin_user.MixinUser/CreateMixinUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MixinUserServer).CreateMixinUser(ctx, req.(*CreateMixinUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MixinUser_ServiceDesc is the grpc.ServiceDesc for MixinUser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MixinUser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mixin_user.MixinUser",
	HandlerType: (*MixinUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMixinUser",
			Handler:    _MixinUser_CreateMixinUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mixin_user.proto",
}