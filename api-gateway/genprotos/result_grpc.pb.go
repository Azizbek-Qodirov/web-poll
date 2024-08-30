// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: protos/result.proto

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
	ResultService_CreateResult_FullMethodName      = "/polling.ResultService/CreateResult"
	ResultService_SavePollAnswer_FullMethodName    = "/polling.ResultService/SavePollAnswer"
	ResultService_GetResultsInExcel_FullMethodName = "/polling.ResultService/GetResultsInExcel"
	ResultService_GetPollResults_FullMethodName    = "/polling.ResultService/GetPollResults"
)

// ResultServiceClient is the client API for ResultService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResultServiceClient interface {
	CreateResult(ctx context.Context, in *CreateResultReq, opts ...grpc.CallOption) (*CreateResultRes, error)
	SavePollAnswer(ctx context.Context, in *SavePollAnswerReq, opts ...grpc.CallOption) (*Void, error)
	GetResultsInExcel(ctx context.Context, in *Void, opts ...grpc.CallOption) (*ExcelResultsRes, error)
	GetPollResults(ctx context.Context, in *ByIDs, opts ...grpc.CallOption) (*ByIDResponse, error)
}

type resultServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewResultServiceClient(cc grpc.ClientConnInterface) ResultServiceClient {
	return &resultServiceClient{cc}
}

func (c *resultServiceClient) CreateResult(ctx context.Context, in *CreateResultReq, opts ...grpc.CallOption) (*CreateResultRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateResultRes)
	err := c.cc.Invoke(ctx, ResultService_CreateResult_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resultServiceClient) SavePollAnswer(ctx context.Context, in *SavePollAnswerReq, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, ResultService_SavePollAnswer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resultServiceClient) GetResultsInExcel(ctx context.Context, in *Void, opts ...grpc.CallOption) (*ExcelResultsRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExcelResultsRes)
	err := c.cc.Invoke(ctx, ResultService_GetResultsInExcel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resultServiceClient) GetPollResults(ctx context.Context, in *ByIDs, opts ...grpc.CallOption) (*ByIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ByIDResponse)
	err := c.cc.Invoke(ctx, ResultService_GetPollResults_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResultServiceServer is the server API for ResultService service.
// All implementations must embed UnimplementedResultServiceServer
// for forward compatibility
type ResultServiceServer interface {
	CreateResult(context.Context, *CreateResultReq) (*CreateResultRes, error)
	SavePollAnswer(context.Context, *SavePollAnswerReq) (*Void, error)
	GetResultsInExcel(context.Context, *Void) (*ExcelResultsRes, error)
	GetPollResults(context.Context, *ByIDs) (*ByIDResponse, error)
	mustEmbedUnimplementedResultServiceServer()
}

// UnimplementedResultServiceServer must be embedded to have forward compatible implementations.
type UnimplementedResultServiceServer struct {
}

func (UnimplementedResultServiceServer) CreateResult(context.Context, *CreateResultReq) (*CreateResultRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateResult not implemented")
}
func (UnimplementedResultServiceServer) SavePollAnswer(context.Context, *SavePollAnswerReq) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SavePollAnswer not implemented")
}
func (UnimplementedResultServiceServer) GetResultsInExcel(context.Context, *Void) (*ExcelResultsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResultsInExcel not implemented")
}
func (UnimplementedResultServiceServer) GetPollResults(context.Context, *ByIDs) (*ByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPollResults not implemented")
}
func (UnimplementedResultServiceServer) mustEmbedUnimplementedResultServiceServer() {}

// UnsafeResultServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResultServiceServer will
// result in compilation errors.
type UnsafeResultServiceServer interface {
	mustEmbedUnimplementedResultServiceServer()
}

func RegisterResultServiceServer(s grpc.ServiceRegistrar, srv ResultServiceServer) {
	s.RegisterService(&ResultService_ServiceDesc, srv)
}

func _ResultService_CreateResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateResultReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResultServiceServer).CreateResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResultService_CreateResult_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResultServiceServer).CreateResult(ctx, req.(*CreateResultReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResultService_SavePollAnswer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SavePollAnswerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResultServiceServer).SavePollAnswer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResultService_SavePollAnswer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResultServiceServer).SavePollAnswer(ctx, req.(*SavePollAnswerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResultService_GetResultsInExcel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResultServiceServer).GetResultsInExcel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResultService_GetResultsInExcel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResultServiceServer).GetResultsInExcel(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResultService_GetPollResults_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByIDs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResultServiceServer).GetPollResults(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResultService_GetPollResults_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResultServiceServer).GetPollResults(ctx, req.(*ByIDs))
	}
	return interceptor(ctx, in, info, handler)
}

// ResultService_ServiceDesc is the grpc.ServiceDesc for ResultService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResultService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "polling.ResultService",
	HandlerType: (*ResultServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateResult",
			Handler:    _ResultService_CreateResult_Handler,
		},
		{
			MethodName: "SavePollAnswer",
			Handler:    _ResultService_SavePollAnswer_Handler,
		},
		{
			MethodName: "GetResultsInExcel",
			Handler:    _ResultService_GetResultsInExcel_Handler,
		},
		{
			MethodName: "GetPollResults",
			Handler:    _ResultService_GetPollResults_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/result.proto",
}
