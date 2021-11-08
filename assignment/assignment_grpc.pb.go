// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package assignment

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

// AssignmentClient is the client API for Assignment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AssignmentClient interface {
	InsertAssignment(ctx context.Context, in *InsertAssignmentRequest, opts ...grpc.CallOption) (*InsertAssignmentResponse, error)
}

type assignmentClient struct {
	cc grpc.ClientConnInterface
}

func NewAssignmentClient(cc grpc.ClientConnInterface) AssignmentClient {
	return &assignmentClient{cc}
}

func (c *assignmentClient) InsertAssignment(ctx context.Context, in *InsertAssignmentRequest, opts ...grpc.CallOption) (*InsertAssignmentResponse, error) {
	out := new(InsertAssignmentResponse)
	err := c.cc.Invoke(ctx, "/assignment.Assignment/InsertAssignment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AssignmentServer is the server API for Assignment service.
// All implementations should embed UnimplementedAssignmentServer
// for forward compatibility
type AssignmentServer interface {
	InsertAssignment(context.Context, *InsertAssignmentRequest) (*InsertAssignmentResponse, error)
}

// UnimplementedAssignmentServer should be embedded to have forward compatible implementations.
type UnimplementedAssignmentServer struct {
}

func (UnimplementedAssignmentServer) InsertAssignment(context.Context, *InsertAssignmentRequest) (*InsertAssignmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertAssignment not implemented")
}

// UnsafeAssignmentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AssignmentServer will
// result in compilation errors.
type UnsafeAssignmentServer interface {
	mustEmbedUnimplementedAssignmentServer()
}

func RegisterAssignmentServer(s grpc.ServiceRegistrar, srv AssignmentServer) {
	s.RegisterService(&Assignment_ServiceDesc, srv)
}

func _Assignment_InsertAssignment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertAssignmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssignmentServer).InsertAssignment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/assignment.Assignment/InsertAssignment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssignmentServer).InsertAssignment(ctx, req.(*InsertAssignmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Assignment_ServiceDesc is the grpc.ServiceDesc for Assignment service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Assignment_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "assignment.Assignment",
	HandlerType: (*AssignmentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertAssignment",
			Handler:    _Assignment_InsertAssignment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "assignment/assignment.proto",
}
