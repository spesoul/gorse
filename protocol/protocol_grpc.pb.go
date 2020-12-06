// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protocol

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// WorkerClient is the client API for Worker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkerClient interface {
	BroadcastModel(ctx context.Context, in *Model, opts ...grpc.CallOption) (*Response, error)
	BroadcastUserPartition(ctx context.Context, in *Partition, opts ...grpc.CallOption) (*Response, error)
	BroadcastItemPartition(ctx context.Context, in *Partition, opts ...grpc.CallOption) (*Response, error)
	BroadcastLabelPartition(ctx context.Context, in *Partition, opts ...grpc.CallOption) (*Response, error)
}

type workerClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkerClient(cc grpc.ClientConnInterface) WorkerClient {
	return &workerClient{cc}
}

func (c *workerClient) BroadcastModel(ctx context.Context, in *Model, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/protocol.Worker/BroadcastModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) BroadcastUserPartition(ctx context.Context, in *Partition, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/protocol.Worker/BroadcastUserPartition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) BroadcastItemPartition(ctx context.Context, in *Partition, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/protocol.Worker/BroadcastItemPartition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) BroadcastLabelPartition(ctx context.Context, in *Partition, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/protocol.Worker/BroadcastLabelPartition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkerServer is the server API for Worker service.
// All implementations must embed UnimplementedWorkerServer
// for forward compatibility
type WorkerServer interface {
	BroadcastModel(context.Context, *Model) (*Response, error)
	BroadcastUserPartition(context.Context, *Partition) (*Response, error)
	BroadcastItemPartition(context.Context, *Partition) (*Response, error)
	BroadcastLabelPartition(context.Context, *Partition) (*Response, error)
	mustEmbedUnimplementedWorkerServer()
}

// UnimplementedWorkerServer must be embedded to have forward compatible implementations.
type UnimplementedWorkerServer struct {
}

func (UnimplementedWorkerServer) BroadcastModel(context.Context, *Model) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BroadcastModel not implemented")
}
func (UnimplementedWorkerServer) BroadcastUserPartition(context.Context, *Partition) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BroadcastUserPartition not implemented")
}
func (UnimplementedWorkerServer) BroadcastItemPartition(context.Context, *Partition) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BroadcastItemPartition not implemented")
}
func (UnimplementedWorkerServer) BroadcastLabelPartition(context.Context, *Partition) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BroadcastLabelPartition not implemented")
}
func (UnimplementedWorkerServer) mustEmbedUnimplementedWorkerServer() {}

// UnsafeWorkerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkerServer will
// result in compilation errors.
type UnsafeWorkerServer interface {
	mustEmbedUnimplementedWorkerServer()
}

func RegisterWorkerServer(s grpc.ServiceRegistrar, srv WorkerServer) {
	s.RegisterService(&_Worker_serviceDesc, srv)
}

func _Worker_BroadcastModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Model)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).BroadcastModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Worker/BroadcastModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).BroadcastModel(ctx, req.(*Model))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_BroadcastUserPartition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Partition)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).BroadcastUserPartition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Worker/BroadcastUserPartition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).BroadcastUserPartition(ctx, req.(*Partition))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_BroadcastItemPartition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Partition)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).BroadcastItemPartition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Worker/BroadcastItemPartition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).BroadcastItemPartition(ctx, req.(*Partition))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_BroadcastLabelPartition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Partition)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).BroadcastLabelPartition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Worker/BroadcastLabelPartition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).BroadcastLabelPartition(ctx, req.(*Partition))
	}
	return interceptor(ctx, in, info, handler)
}

var _Worker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.Worker",
	HandlerType: (*WorkerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BroadcastModel",
			Handler:    _Worker_BroadcastModel_Handler,
		},
		{
			MethodName: "BroadcastUserPartition",
			Handler:    _Worker_BroadcastUserPartition_Handler,
		},
		{
			MethodName: "BroadcastItemPartition",
			Handler:    _Worker_BroadcastItemPartition_Handler,
		},
		{
			MethodName: "BroadcastLabelPartition",
			Handler:    _Worker_BroadcastLabelPartition_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protocol.proto",
}
