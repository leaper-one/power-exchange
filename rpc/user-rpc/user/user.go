// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package user

import (
	"context"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	User interface {
		//  创建用户
		CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
		//  查询用户
		QueryUser(ctx context.Context, in *QueryUserRequest, opts ...grpc.CallOption) (*QueryUserResponse, error)
		//  更新用户
		UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
		//  删除用户
		DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

//  创建用户
func (m *defaultUser) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	client := NewUserClient(m.cli.Conn())
	return client.CreateUser(ctx, in, opts...)
}

//  查询用户
func (m *defaultUser) QueryUser(ctx context.Context, in *QueryUserRequest, opts ...grpc.CallOption) (*QueryUserResponse, error) {
	client := NewUserClient(m.cli.Conn())
	return client.QueryUser(ctx, in, opts...)
}

//  更新用户
func (m *defaultUser) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	client := NewUserClient(m.cli.Conn())
	return client.UpdateUser(ctx, in, opts...)
}

//  删除用户
func (m *defaultUser) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	client := NewUserClient(m.cli.Conn())
	return client.DeleteUser(ctx, in, opts...)
}
