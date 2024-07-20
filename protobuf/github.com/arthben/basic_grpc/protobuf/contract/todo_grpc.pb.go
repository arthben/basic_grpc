// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: todo.proto

package contract

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ThingsTodoClient is the client API for ThingsTodo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ThingsTodoClient interface {
	TaskList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*TodoItems, error)
	AddTask(ctx context.Context, in *AddTodoReq, opts ...grpc.CallOption) (*TodoItem, error)
}

type thingsTodoClient struct {
	cc grpc.ClientConnInterface
}

func NewThingsTodoClient(cc grpc.ClientConnInterface) ThingsTodoClient {
	return &thingsTodoClient{cc}
}

func (c *thingsTodoClient) TaskList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*TodoItems, error) {
	out := new(TodoItems)
	err := c.cc.Invoke(ctx, "/contract.ThingsTodo/TaskList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thingsTodoClient) AddTask(ctx context.Context, in *AddTodoReq, opts ...grpc.CallOption) (*TodoItem, error) {
	out := new(TodoItem)
	err := c.cc.Invoke(ctx, "/contract.ThingsTodo/AddTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ThingsTodoServer is the server API for ThingsTodo service.
// All implementations must embed UnimplementedThingsTodoServer
// for forward compatibility
type ThingsTodoServer interface {
	TaskList(context.Context, *emptypb.Empty) (*TodoItems, error)
	AddTask(context.Context, *AddTodoReq) (*TodoItem, error)
	mustEmbedUnimplementedThingsTodoServer()
}

// UnimplementedThingsTodoServer must be embedded to have forward compatible implementations.
type UnimplementedThingsTodoServer struct {
}

func (UnimplementedThingsTodoServer) TaskList(context.Context, *emptypb.Empty) (*TodoItems, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskList not implemented")
}
func (UnimplementedThingsTodoServer) AddTask(context.Context, *AddTodoReq) (*TodoItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTask not implemented")
}
func (UnimplementedThingsTodoServer) mustEmbedUnimplementedThingsTodoServer() {}

// UnsafeThingsTodoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ThingsTodoServer will
// result in compilation errors.
type UnsafeThingsTodoServer interface {
	mustEmbedUnimplementedThingsTodoServer()
}

func RegisterThingsTodoServer(s grpc.ServiceRegistrar, srv ThingsTodoServer) {
	s.RegisterService(&ThingsTodo_ServiceDesc, srv)
}

func _ThingsTodo_TaskList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThingsTodoServer).TaskList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contract.ThingsTodo/TaskList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThingsTodoServer).TaskList(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThingsTodo_AddTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTodoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThingsTodoServer).AddTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contract.ThingsTodo/AddTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThingsTodoServer).AddTask(ctx, req.(*AddTodoReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ThingsTodo_ServiceDesc is the grpc.ServiceDesc for ThingsTodo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ThingsTodo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "contract.ThingsTodo",
	HandlerType: (*ThingsTodoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TaskList",
			Handler:    _ThingsTodo_TaskList_Handler,
		},
		{
			MethodName: "AddTask",
			Handler:    _ThingsTodo_AddTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todo.proto",
}