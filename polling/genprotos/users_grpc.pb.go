// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: protos/users.proto

package genprotos

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
	UserService_Register_FullMethodName       = "/polling.UserService/Register"
	UserService_ConfirmUser_FullMethodName    = "/polling.UserService/ConfirmUser"
	UserService_Profile_FullMethodName        = "/polling.UserService/Profile"
	UserService_UpdatePassword_FullMethodName = "/polling.UserService/UpdatePassword"
	UserService_IsEmailExists_FullMethodName  = "/polling.UserService/IsEmailExists"
	UserService_GetByID_FullMethodName        = "/polling.UserService/GetByID"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*Void, error)
	ConfirmUser(ctx context.Context, in *ConfirmUserReq, opts ...grpc.CallOption) (*Void, error)
	Profile(ctx context.Context, in *GetProfileReq, opts ...grpc.CallOption) (*GetProfileResp, error)
	UpdatePassword(ctx context.Context, in *UpdatePasswordReq, opts ...grpc.CallOption) (*Void, error)
	IsEmailExists(ctx context.Context, in *IsEmailExistsReq, opts ...grpc.CallOption) (*IsEmailExistsResp, error)
	GetByID(ctx context.Context, in *GetProfileByIdReq, opts ...grpc.CallOption) (*GetProfileByIdResp, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, UserService_Register_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ConfirmUser(ctx context.Context, in *ConfirmUserReq, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, UserService_ConfirmUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Profile(ctx context.Context, in *GetProfileReq, opts ...grpc.CallOption) (*GetProfileResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetProfileResp)
	err := c.cc.Invoke(ctx, UserService_Profile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdatePassword(ctx context.Context, in *UpdatePasswordReq, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, UserService_UpdatePassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) IsEmailExists(ctx context.Context, in *IsEmailExistsReq, opts ...grpc.CallOption) (*IsEmailExistsResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IsEmailExistsResp)
	err := c.cc.Invoke(ctx, UserService_IsEmailExists_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetByID(ctx context.Context, in *GetProfileByIdReq, opts ...grpc.CallOption) (*GetProfileByIdResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetProfileByIdResp)
	err := c.cc.Invoke(ctx, UserService_GetByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	Register(context.Context, *RegisterReq) (*Void, error)
	ConfirmUser(context.Context, *ConfirmUserReq) (*Void, error)
	Profile(context.Context, *GetProfileReq) (*GetProfileResp, error)
	UpdatePassword(context.Context, *UpdatePasswordReq) (*Void, error)
	IsEmailExists(context.Context, *IsEmailExistsReq) (*IsEmailExistsResp, error)
	GetByID(context.Context, *GetProfileByIdReq) (*GetProfileByIdResp, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) Register(context.Context, *RegisterReq) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserServiceServer) ConfirmUser(context.Context, *ConfirmUserReq) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmUser not implemented")
}
func (UnimplementedUserServiceServer) Profile(context.Context, *GetProfileReq) (*GetProfileResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Profile not implemented")
}
func (UnimplementedUserServiceServer) UpdatePassword(context.Context, *UpdatePasswordReq) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePassword not implemented")
}
func (UnimplementedUserServiceServer) IsEmailExists(context.Context, *IsEmailExistsReq) (*IsEmailExistsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsEmailExists not implemented")
}
func (UnimplementedUserServiceServer) GetByID(context.Context, *GetProfileByIdReq) (*GetProfileByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ConfirmUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ConfirmUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_ConfirmUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ConfirmUser(ctx, req.(*ConfirmUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Profile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Profile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Profile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Profile(ctx, req.(*GetProfileReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdatePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdatePassword(ctx, req.(*UpdatePasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_IsEmailExists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsEmailExistsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).IsEmailExists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_IsEmailExists_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).IsEmailExists(ctx, req.(*IsEmailExistsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetByID(ctx, req.(*GetProfileByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "polling.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _UserService_Register_Handler,
		},
		{
			MethodName: "ConfirmUser",
			Handler:    _UserService_ConfirmUser_Handler,
		},
		{
			MethodName: "Profile",
			Handler:    _UserService_Profile_Handler,
		},
		{
			MethodName: "UpdatePassword",
			Handler:    _UserService_UpdatePassword_Handler,
		},
		{
			MethodName: "IsEmailExists",
			Handler:    _UserService_IsEmailExists_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _UserService_GetByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/users.proto",
}
