// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.0--rc3
// source: blog/blog.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	User_UserLogin_FullMethodName         = "/proto.User/UserLogin"
	User_UserUsePhoneLogin_FullMethodName = "/proto.User/UserUsePhoneLogin"
	User_UserUseEmailLogin_FullMethodName = "/proto.User/UserUseEmailLogin"
	User_SetUserName_FullMethodName       = "/proto.User/SetUserName"
	User_UserRegister_FullMethodName      = "/proto.User/UserRegister"
	User_UserNameExist_FullMethodName     = "/proto.User/UserNameExist"
	User_BindEmail_FullMethodName         = "/proto.User/BindEmail"
	User_SendEmailCode_FullMethodName     = "/proto.User/SendEmailCode"
	User_SendPhoneCode_FullMethodName     = "/proto.User/SendPhoneCode"
)

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error)
	UserUsePhoneLogin(ctx context.Context, in *UserUsePhoneLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error)
	UserUseEmailLogin(ctx context.Context, in *UserUseEmailLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error)
	SetUserName(ctx context.Context, in *SetUserNameRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	UserRegister(ctx context.Context, in *UserRegisterRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	UserNameExist(ctx context.Context, in *UserNameExistRequest, opts ...grpc.CallOption) (*UserNameExistResponse, error)
	BindEmail(ctx context.Context, in *BindEmailRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	SendEmailCode(ctx context.Context, in *SendCodeRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	SendPhoneCode(ctx context.Context, in *SendPhoneCodeRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserLoginResponse)
	err := c.cc.Invoke(ctx, User_UserLogin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserUsePhoneLogin(ctx context.Context, in *UserUsePhoneLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserLoginResponse)
	err := c.cc.Invoke(ctx, User_UserUsePhoneLogin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserUseEmailLogin(ctx context.Context, in *UserUseEmailLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserLoginResponse)
	err := c.cc.Invoke(ctx, User_UserUseEmailLogin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) SetUserName(ctx context.Context, in *SetUserNameRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, User_SetUserName_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserRegister(ctx context.Context, in *UserRegisterRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, User_UserRegister_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserNameExist(ctx context.Context, in *UserNameExistRequest, opts ...grpc.CallOption) (*UserNameExistResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserNameExistResponse)
	err := c.cc.Invoke(ctx, User_UserNameExist_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) BindEmail(ctx context.Context, in *BindEmailRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, User_BindEmail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) SendEmailCode(ctx context.Context, in *SendCodeRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, User_SendEmailCode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) SendPhoneCode(ctx context.Context, in *SendPhoneCodeRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, User_SendPhoneCode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	UserLogin(context.Context, *UserLoginRequest) (*UserLoginResponse, error)
	UserUsePhoneLogin(context.Context, *UserUsePhoneLoginRequest) (*UserLoginResponse, error)
	UserUseEmailLogin(context.Context, *UserUseEmailLoginRequest) (*UserLoginResponse, error)
	SetUserName(context.Context, *SetUserNameRequest) (*EmptyResponse, error)
	UserRegister(context.Context, *UserRegisterRequest) (*EmptyResponse, error)
	UserNameExist(context.Context, *UserNameExistRequest) (*UserNameExistResponse, error)
	BindEmail(context.Context, *BindEmailRequest) (*EmptyResponse, error)
	SendEmailCode(context.Context, *SendCodeRequest) (*EmptyResponse, error)
	SendPhoneCode(context.Context, *SendPhoneCodeRequest) (*EmptyResponse, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) UserLogin(context.Context, *UserLoginRequest) (*UserLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (UnimplementedUserServer) UserUsePhoneLogin(context.Context, *UserUsePhoneLoginRequest) (*UserLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserUsePhoneLogin not implemented")
}
func (UnimplementedUserServer) UserUseEmailLogin(context.Context, *UserUseEmailLoginRequest) (*UserLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserUseEmailLogin not implemented")
}
func (UnimplementedUserServer) SetUserName(context.Context, *SetUserNameRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUserName not implemented")
}
func (UnimplementedUserServer) UserRegister(context.Context, *UserRegisterRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRegister not implemented")
}
func (UnimplementedUserServer) UserNameExist(context.Context, *UserNameExistRequest) (*UserNameExistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserNameExist not implemented")
}
func (UnimplementedUserServer) BindEmail(context.Context, *BindEmailRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BindEmail not implemented")
}
func (UnimplementedUserServer) SendEmailCode(context.Context, *SendCodeRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEmailCode not implemented")
}
func (UnimplementedUserServer) SendPhoneCode(context.Context, *SendPhoneCodeRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendPhoneCode not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserLogin(ctx, req.(*UserLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserUsePhoneLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUsePhoneLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserUsePhoneLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserUsePhoneLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserUsePhoneLogin(ctx, req.(*UserUsePhoneLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserUseEmailLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUseEmailLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserUseEmailLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserUseEmailLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserUseEmailLogin(ctx, req.(*UserUseEmailLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_SetUserName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetUserNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SetUserName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_SetUserName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SetUserName(ctx, req.(*SetUserNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserRegister_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserRegister(ctx, req.(*UserRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserNameExist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserNameExistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserNameExist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserNameExist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserNameExist(ctx, req.(*UserNameExistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_BindEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BindEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).BindEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_BindEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).BindEmail(ctx, req.(*BindEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_SendEmailCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SendEmailCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_SendEmailCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SendEmailCode(ctx, req.(*SendCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_SendPhoneCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendPhoneCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SendPhoneCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_SendPhoneCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SendPhoneCode(ctx, req.(*SendPhoneCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserLogin",
			Handler:    _User_UserLogin_Handler,
		},
		{
			MethodName: "UserUsePhoneLogin",
			Handler:    _User_UserUsePhoneLogin_Handler,
		},
		{
			MethodName: "UserUseEmailLogin",
			Handler:    _User_UserUseEmailLogin_Handler,
		},
		{
			MethodName: "SetUserName",
			Handler:    _User_SetUserName_Handler,
		},
		{
			MethodName: "UserRegister",
			Handler:    _User_UserRegister_Handler,
		},
		{
			MethodName: "UserNameExist",
			Handler:    _User_UserNameExist_Handler,
		},
		{
			MethodName: "BindEmail",
			Handler:    _User_BindEmail_Handler,
		},
		{
			MethodName: "SendEmailCode",
			Handler:    _User_SendEmailCode_Handler,
		},
		{
			MethodName: "SendPhoneCode",
			Handler:    _User_SendPhoneCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blog/blog.proto",
}
