// Copyright (c) 2024 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: assignment_function.proto

package assignmentFunction

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

// AssignmentFunctionClient is the client API for AssignmentFunction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AssignmentFunctionClient interface {
	Assign(ctx context.Context, in *AssignmentRequest, opts ...grpc.CallOption) (*AssignmentResponse, error)
}

type assignmentFunctionClient struct {
	cc grpc.ClientConnInterface
}

func NewAssignmentFunctionClient(cc grpc.ClientConnInterface) AssignmentFunctionClient {
	return &assignmentFunctionClient{cc}
}

func (c *assignmentFunctionClient) Assign(ctx context.Context, in *AssignmentRequest, opts ...grpc.CallOption) (*AssignmentResponse, error) {
	out := new(AssignmentResponse)
	err := c.cc.Invoke(ctx, "/accelbyte.challenge.assignmentFunction.AssignmentFunction/Assign", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AssignmentFunctionServer is the server API for AssignmentFunction service.
// All implementations must embed UnimplementedAssignmentFunctionServer
// for forward compatibility
type AssignmentFunctionServer interface {
	Assign(context.Context, *AssignmentRequest) (*AssignmentResponse, error)
	mustEmbedUnimplementedAssignmentFunctionServer()
}

// UnimplementedAssignmentFunctionServer must be embedded to have forward compatible implementations.
type UnimplementedAssignmentFunctionServer struct {
}

func (UnimplementedAssignmentFunctionServer) Assign(context.Context, *AssignmentRequest) (*AssignmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Assign not implemented")
}
func (UnimplementedAssignmentFunctionServer) mustEmbedUnimplementedAssignmentFunctionServer() {}

// UnsafeAssignmentFunctionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AssignmentFunctionServer will
// result in compilation errors.
type UnsafeAssignmentFunctionServer interface {
	mustEmbedUnimplementedAssignmentFunctionServer()
}

func RegisterAssignmentFunctionServer(s grpc.ServiceRegistrar, srv AssignmentFunctionServer) {
	s.RegisterService(&AssignmentFunction_ServiceDesc, srv)
}

func _AssignmentFunction_Assign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssignmentFunctionServer).Assign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accelbyte.challenge.assignmentFunction.AssignmentFunction/Assign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssignmentFunctionServer).Assign(ctx, req.(*AssignmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AssignmentFunction_ServiceDesc is the grpc.ServiceDesc for AssignmentFunction service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AssignmentFunction_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "accelbyte.challenge.assignmentFunction.AssignmentFunction",
	HandlerType: (*AssignmentFunctionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Assign",
			Handler:    _AssignmentFunction_Assign_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "assignment_function.proto",
}
