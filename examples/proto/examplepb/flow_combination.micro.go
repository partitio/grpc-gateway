// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: examples/proto/examplepb/flow_combination.proto

package examplepb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for FlowCombination service

type FlowCombinationService interface {
	RpcEmptyRpc(ctx context.Context, in *EmptyProto, opts ...client.CallOption) (*EmptyProto, error)
	RpcEmptyStream(ctx context.Context, in *EmptyProto, opts ...client.CallOption) (FlowCombination_RpcEmptyStreamService, error)
	StreamEmptyRpc(ctx context.Context, opts ...client.CallOption) (FlowCombination_StreamEmptyRpcService, error)
	StreamEmptyStream(ctx context.Context, opts ...client.CallOption) (FlowCombination_StreamEmptyStreamService, error)
	RpcBodyRpc(ctx context.Context, in *NonEmptyProto, opts ...client.CallOption) (*EmptyProto, error)
	RpcPathSingleNestedRpc(ctx context.Context, in *SingleNestedProto, opts ...client.CallOption) (*EmptyProto, error)
	RpcPathNestedRpc(ctx context.Context, in *NestedProto, opts ...client.CallOption) (*EmptyProto, error)
	RpcBodyStream(ctx context.Context, in *NonEmptyProto, opts ...client.CallOption) (FlowCombination_RpcBodyStreamService, error)
	RpcPathSingleNestedStream(ctx context.Context, in *SingleNestedProto, opts ...client.CallOption) (FlowCombination_RpcPathSingleNestedStreamService, error)
	RpcPathNestedStream(ctx context.Context, in *NestedProto, opts ...client.CallOption) (FlowCombination_RpcPathNestedStreamService, error)
}

type flowCombinationService struct {
	c    client.Client
	name string
}

func NewFlowCombinationService(name string, c client.Client) FlowCombinationService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "grpc.gateway.examples.examplepb"
	}
	return &flowCombinationService{
		c:    c,
		name: name,
	}
}

func (c *flowCombinationService) RpcEmptyRpc(ctx context.Context, in *EmptyProto, opts ...client.CallOption) (*EmptyProto, error) {
	req := c.c.NewRequest(c.name, "FlowCombination.RpcEmptyRpc", in)
	out := new(EmptyProto)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flowCombinationService) RpcEmptyStream(ctx context.Context, in *EmptyProto, opts ...client.CallOption) (FlowCombination_RpcEmptyStreamService, error) {
	req := c.c.NewRequest(c.name, "FlowCombination.RpcEmptyStream", &EmptyProto{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &flowCombinationServiceRpcEmptyStream{stream}, nil
}

type FlowCombination_RpcEmptyStreamService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*EmptyProto, error)
}

type flowCombinationServiceRpcEmptyStream struct {
	stream client.Stream
}

func (x *flowCombinationServiceRpcEmptyStream) Close() error {
	return x.stream.Close()
}

func (x *flowCombinationServiceRpcEmptyStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *flowCombinationServiceRpcEmptyStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *flowCombinationServiceRpcEmptyStream) Recv() (*EmptyProto, error) {
	m := new(EmptyProto)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *flowCombinationService) StreamEmptyRpc(ctx context.Context, opts ...client.CallOption) (FlowCombination_StreamEmptyRpcService, error) {
	req := c.c.NewRequest(c.name, "FlowCombination.StreamEmptyRpc", &EmptyProto{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &flowCombinationServiceStreamEmptyRpc{stream}, nil
}

type FlowCombination_StreamEmptyRpcService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*EmptyProto) error
}

type flowCombinationServiceStreamEmptyRpc struct {
	stream client.Stream
}

func (x *flowCombinationServiceStreamEmptyRpc) Close() error {
	return x.stream.Close()
}

func (x *flowCombinationServiceStreamEmptyRpc) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *flowCombinationServiceStreamEmptyRpc) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *flowCombinationServiceStreamEmptyRpc) Send(m *EmptyProto) error {
	return x.stream.Send(m)
}

func (c *flowCombinationService) StreamEmptyStream(ctx context.Context, opts ...client.CallOption) (FlowCombination_StreamEmptyStreamService, error) {
	req := c.c.NewRequest(c.name, "FlowCombination.StreamEmptyStream", &EmptyProto{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &flowCombinationServiceStreamEmptyStream{stream}, nil
}

type FlowCombination_StreamEmptyStreamService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*EmptyProto) error
	Recv() (*EmptyProto, error)
}

type flowCombinationServiceStreamEmptyStream struct {
	stream client.Stream
}

func (x *flowCombinationServiceStreamEmptyStream) Close() error {
	return x.stream.Close()
}

func (x *flowCombinationServiceStreamEmptyStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *flowCombinationServiceStreamEmptyStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *flowCombinationServiceStreamEmptyStream) Send(m *EmptyProto) error {
	return x.stream.Send(m)
}

func (x *flowCombinationServiceStreamEmptyStream) Recv() (*EmptyProto, error) {
	m := new(EmptyProto)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *flowCombinationService) RpcBodyRpc(ctx context.Context, in *NonEmptyProto, opts ...client.CallOption) (*EmptyProto, error) {
	req := c.c.NewRequest(c.name, "FlowCombination.RpcBodyRpc", in)
	out := new(EmptyProto)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flowCombinationService) RpcPathSingleNestedRpc(ctx context.Context, in *SingleNestedProto, opts ...client.CallOption) (*EmptyProto, error) {
	req := c.c.NewRequest(c.name, "FlowCombination.RpcPathSingleNestedRpc", in)
	out := new(EmptyProto)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flowCombinationService) RpcPathNestedRpc(ctx context.Context, in *NestedProto, opts ...client.CallOption) (*EmptyProto, error) {
	req := c.c.NewRequest(c.name, "FlowCombination.RpcPathNestedRpc", in)
	out := new(EmptyProto)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flowCombinationService) RpcBodyStream(ctx context.Context, in *NonEmptyProto, opts ...client.CallOption) (FlowCombination_RpcBodyStreamService, error) {
	req := c.c.NewRequest(c.name, "FlowCombination.RpcBodyStream", &NonEmptyProto{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &flowCombinationServiceRpcBodyStream{stream}, nil
}

type FlowCombination_RpcBodyStreamService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*EmptyProto, error)
}

type flowCombinationServiceRpcBodyStream struct {
	stream client.Stream
}

func (x *flowCombinationServiceRpcBodyStream) Close() error {
	return x.stream.Close()
}

func (x *flowCombinationServiceRpcBodyStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *flowCombinationServiceRpcBodyStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *flowCombinationServiceRpcBodyStream) Recv() (*EmptyProto, error) {
	m := new(EmptyProto)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *flowCombinationService) RpcPathSingleNestedStream(ctx context.Context, in *SingleNestedProto, opts ...client.CallOption) (FlowCombination_RpcPathSingleNestedStreamService, error) {
	req := c.c.NewRequest(c.name, "FlowCombination.RpcPathSingleNestedStream", &SingleNestedProto{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &flowCombinationServiceRpcPathSingleNestedStream{stream}, nil
}

type FlowCombination_RpcPathSingleNestedStreamService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*EmptyProto, error)
}

type flowCombinationServiceRpcPathSingleNestedStream struct {
	stream client.Stream
}

func (x *flowCombinationServiceRpcPathSingleNestedStream) Close() error {
	return x.stream.Close()
}

func (x *flowCombinationServiceRpcPathSingleNestedStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *flowCombinationServiceRpcPathSingleNestedStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *flowCombinationServiceRpcPathSingleNestedStream) Recv() (*EmptyProto, error) {
	m := new(EmptyProto)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *flowCombinationService) RpcPathNestedStream(ctx context.Context, in *NestedProto, opts ...client.CallOption) (FlowCombination_RpcPathNestedStreamService, error) {
	req := c.c.NewRequest(c.name, "FlowCombination.RpcPathNestedStream", &NestedProto{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &flowCombinationServiceRpcPathNestedStream{stream}, nil
}

type FlowCombination_RpcPathNestedStreamService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*EmptyProto, error)
}

type flowCombinationServiceRpcPathNestedStream struct {
	stream client.Stream
}

func (x *flowCombinationServiceRpcPathNestedStream) Close() error {
	return x.stream.Close()
}

func (x *flowCombinationServiceRpcPathNestedStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *flowCombinationServiceRpcPathNestedStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *flowCombinationServiceRpcPathNestedStream) Recv() (*EmptyProto, error) {
	m := new(EmptyProto)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for FlowCombination service

type FlowCombinationHandler interface {
	RpcEmptyRpc(context.Context, *EmptyProto, *EmptyProto) error
	RpcEmptyStream(context.Context, *EmptyProto, FlowCombination_RpcEmptyStreamStream) error
	StreamEmptyRpc(context.Context, FlowCombination_StreamEmptyRpcStream) error
	StreamEmptyStream(context.Context, FlowCombination_StreamEmptyStreamStream) error
	RpcBodyRpc(context.Context, *NonEmptyProto, *EmptyProto) error
	RpcPathSingleNestedRpc(context.Context, *SingleNestedProto, *EmptyProto) error
	RpcPathNestedRpc(context.Context, *NestedProto, *EmptyProto) error
	RpcBodyStream(context.Context, *NonEmptyProto, FlowCombination_RpcBodyStreamStream) error
	RpcPathSingleNestedStream(context.Context, *SingleNestedProto, FlowCombination_RpcPathSingleNestedStreamStream) error
	RpcPathNestedStream(context.Context, *NestedProto, FlowCombination_RpcPathNestedStreamStream) error
}

func RegisterFlowCombinationHandler(s server.Server, hdlr FlowCombinationHandler, opts ...server.HandlerOption) error {
	type flowCombination interface {
		RpcEmptyRpc(ctx context.Context, in *EmptyProto, out *EmptyProto) error
		RpcEmptyStream(ctx context.Context, stream server.Stream) error
		StreamEmptyRpc(ctx context.Context, stream server.Stream) error
		StreamEmptyStream(ctx context.Context, stream server.Stream) error
		RpcBodyRpc(ctx context.Context, in *NonEmptyProto, out *EmptyProto) error
		RpcPathSingleNestedRpc(ctx context.Context, in *SingleNestedProto, out *EmptyProto) error
		RpcPathNestedRpc(ctx context.Context, in *NestedProto, out *EmptyProto) error
		RpcBodyStream(ctx context.Context, stream server.Stream) error
		RpcPathSingleNestedStream(ctx context.Context, stream server.Stream) error
		RpcPathNestedStream(ctx context.Context, stream server.Stream) error
	}
	type FlowCombination struct {
		flowCombination
	}
	h := &flowCombinationHandler{hdlr}
	return s.Handle(s.NewHandler(&FlowCombination{h}, opts...))
}

type flowCombinationHandler struct {
	FlowCombinationHandler
}

func (h *flowCombinationHandler) RpcEmptyRpc(ctx context.Context, in *EmptyProto, out *EmptyProto) error {
	return h.FlowCombinationHandler.RpcEmptyRpc(ctx, in, out)
}

func (h *flowCombinationHandler) RpcEmptyStream(ctx context.Context, stream server.Stream) error {
	m := new(EmptyProto)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.FlowCombinationHandler.RpcEmptyStream(ctx, m, &flowCombinationRpcEmptyStreamStream{stream})
}

type FlowCombination_RpcEmptyStreamStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*EmptyProto) error
}

type flowCombinationRpcEmptyStreamStream struct {
	stream server.Stream
}

func (x *flowCombinationRpcEmptyStreamStream) Close() error {
	return x.stream.Close()
}

func (x *flowCombinationRpcEmptyStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *flowCombinationRpcEmptyStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *flowCombinationRpcEmptyStreamStream) Send(m *EmptyProto) error {
	return x.stream.Send(m)
}

func (h *flowCombinationHandler) StreamEmptyRpc(ctx context.Context, stream server.Stream) error {
	return h.FlowCombinationHandler.StreamEmptyRpc(ctx, &flowCombinationStreamEmptyRpcStream{stream})
}

type FlowCombination_StreamEmptyRpcStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*EmptyProto, error)
}

type flowCombinationStreamEmptyRpcStream struct {
	stream server.Stream
}

func (x *flowCombinationStreamEmptyRpcStream) Close() error {
	return x.stream.Close()
}

func (x *flowCombinationStreamEmptyRpcStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *flowCombinationStreamEmptyRpcStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *flowCombinationStreamEmptyRpcStream) Recv() (*EmptyProto, error) {
	m := new(EmptyProto)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (h *flowCombinationHandler) StreamEmptyStream(ctx context.Context, stream server.Stream) error {
	return h.FlowCombinationHandler.StreamEmptyStream(ctx, &flowCombinationStreamEmptyStreamStream{stream})
}

type FlowCombination_StreamEmptyStreamStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*EmptyProto) error
	Recv() (*EmptyProto, error)
}

type flowCombinationStreamEmptyStreamStream struct {
	stream server.Stream
}

func (x *flowCombinationStreamEmptyStreamStream) Close() error {
	return x.stream.Close()
}

func (x *flowCombinationStreamEmptyStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *flowCombinationStreamEmptyStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *flowCombinationStreamEmptyStreamStream) Send(m *EmptyProto) error {
	return x.stream.Send(m)
}

func (x *flowCombinationStreamEmptyStreamStream) Recv() (*EmptyProto, error) {
	m := new(EmptyProto)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (h *flowCombinationHandler) RpcBodyRpc(ctx context.Context, in *NonEmptyProto, out *EmptyProto) error {
	return h.FlowCombinationHandler.RpcBodyRpc(ctx, in, out)
}

func (h *flowCombinationHandler) RpcPathSingleNestedRpc(ctx context.Context, in *SingleNestedProto, out *EmptyProto) error {
	return h.FlowCombinationHandler.RpcPathSingleNestedRpc(ctx, in, out)
}

func (h *flowCombinationHandler) RpcPathNestedRpc(ctx context.Context, in *NestedProto, out *EmptyProto) error {
	return h.FlowCombinationHandler.RpcPathNestedRpc(ctx, in, out)
}

func (h *flowCombinationHandler) RpcBodyStream(ctx context.Context, stream server.Stream) error {
	m := new(NonEmptyProto)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.FlowCombinationHandler.RpcBodyStream(ctx, m, &flowCombinationRpcBodyStreamStream{stream})
}

type FlowCombination_RpcBodyStreamStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*EmptyProto) error
}

type flowCombinationRpcBodyStreamStream struct {
	stream server.Stream
}

func (x *flowCombinationRpcBodyStreamStream) Close() error {
	return x.stream.Close()
}

func (x *flowCombinationRpcBodyStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *flowCombinationRpcBodyStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *flowCombinationRpcBodyStreamStream) Send(m *EmptyProto) error {
	return x.stream.Send(m)
}

func (h *flowCombinationHandler) RpcPathSingleNestedStream(ctx context.Context, stream server.Stream) error {
	m := new(SingleNestedProto)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.FlowCombinationHandler.RpcPathSingleNestedStream(ctx, m, &flowCombinationRpcPathSingleNestedStreamStream{stream})
}

type FlowCombination_RpcPathSingleNestedStreamStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*EmptyProto) error
}

type flowCombinationRpcPathSingleNestedStreamStream struct {
	stream server.Stream
}

func (x *flowCombinationRpcPathSingleNestedStreamStream) Close() error {
	return x.stream.Close()
}

func (x *flowCombinationRpcPathSingleNestedStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *flowCombinationRpcPathSingleNestedStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *flowCombinationRpcPathSingleNestedStreamStream) Send(m *EmptyProto) error {
	return x.stream.Send(m)
}

func (h *flowCombinationHandler) RpcPathNestedStream(ctx context.Context, stream server.Stream) error {
	m := new(NestedProto)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.FlowCombinationHandler.RpcPathNestedStream(ctx, m, &flowCombinationRpcPathNestedStreamStream{stream})
}

type FlowCombination_RpcPathNestedStreamStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*EmptyProto) error
}

type flowCombinationRpcPathNestedStreamStream struct {
	stream server.Stream
}

func (x *flowCombinationRpcPathNestedStreamStream) Close() error {
	return x.stream.Close()
}

func (x *flowCombinationRpcPathNestedStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *flowCombinationRpcPathNestedStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *flowCombinationRpcPathNestedStreamStream) Send(m *EmptyProto) error {
	return x.stream.Send(m)
}
