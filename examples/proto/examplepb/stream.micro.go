// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: examples/proto/examplepb/stream.proto

package examplepb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf2 "github.com/golang/protobuf/ptypes/empty"
import grpc_gateway_examples_sub "github.com/partitio/micro-gateway/examples/proto/sub"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = grpc_gateway_examples_sub.StringMessage{}
var _ = google_protobuf2.Empty{}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for StreamService service

type StreamService interface {
	BulkCreate(ctx context.Context, opts ...client.CallOption) (StreamService_BulkCreateService, error)
	List(ctx context.Context, in *google_protobuf2.Empty, opts ...client.CallOption) (StreamService_ListService, error)
	BulkEcho(ctx context.Context, opts ...client.CallOption) (StreamService_BulkEchoService, error)
}

type streamService struct {
	c    client.Client
	name string
}

func NewStreamService(name string, c client.Client) StreamService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "grpc.gateway.examples.examplepb"
	}
	return &streamService{
		c:    c,
		name: name,
	}
}

func (c *streamService) BulkCreate(ctx context.Context, opts ...client.CallOption) (StreamService_BulkCreateService, error) {
	req := c.c.NewRequest(c.name, "StreamService.BulkCreate", &ABitOfEverything{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &streamServiceBulkCreate{stream}, nil
}

type StreamService_BulkCreateService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*ABitOfEverything) error
}

type streamServiceBulkCreate struct {
	stream client.Stream
}

func (x *streamServiceBulkCreate) Close() error {
	return x.stream.Close()
}

func (x *streamServiceBulkCreate) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *streamServiceBulkCreate) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *streamServiceBulkCreate) Send(m *ABitOfEverything) error {
	return x.stream.Send(m)
}

func (c *streamService) List(ctx context.Context, in *google_protobuf2.Empty, opts ...client.CallOption) (StreamService_ListService, error) {
	req := c.c.NewRequest(c.name, "StreamService.List", &google_protobuf2.Empty{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &streamServiceList{stream}, nil
}

type StreamService_ListService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*ABitOfEverything, error)
}

type streamServiceList struct {
	stream client.Stream
}

func (x *streamServiceList) Close() error {
	return x.stream.Close()
}

func (x *streamServiceList) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *streamServiceList) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *streamServiceList) Recv() (*ABitOfEverything, error) {
	m := new(ABitOfEverything)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamService) BulkEcho(ctx context.Context, opts ...client.CallOption) (StreamService_BulkEchoService, error) {
	req := c.c.NewRequest(c.name, "StreamService.BulkEcho", &grpc_gateway_examples_sub.StringMessage{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &streamServiceBulkEcho{stream}, nil
}

type StreamService_BulkEchoService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*grpc_gateway_examples_sub.StringMessage) error
	Recv() (*grpc_gateway_examples_sub.StringMessage, error)
}

type streamServiceBulkEcho struct {
	stream client.Stream
}

func (x *streamServiceBulkEcho) Close() error {
	return x.stream.Close()
}

func (x *streamServiceBulkEcho) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *streamServiceBulkEcho) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *streamServiceBulkEcho) Send(m *grpc_gateway_examples_sub.StringMessage) error {
	return x.stream.Send(m)
}

func (x *streamServiceBulkEcho) Recv() (*grpc_gateway_examples_sub.StringMessage, error) {
	m := new(grpc_gateway_examples_sub.StringMessage)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for StreamService service

type StreamServiceHandler interface {
	BulkCreate(context.Context, StreamService_BulkCreateStream) error
	List(context.Context, *google_protobuf2.Empty, StreamService_ListStream) error
	BulkEcho(context.Context, StreamService_BulkEchoStream) error
}

func RegisterStreamServiceHandler(s server.Server, hdlr StreamServiceHandler, opts ...server.HandlerOption) error {
	type streamService interface {
		BulkCreate(ctx context.Context, stream server.Stream) error
		List(ctx context.Context, stream server.Stream) error
		BulkEcho(ctx context.Context, stream server.Stream) error
	}
	type StreamService struct {
		streamService
	}
	h := &streamServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&StreamService{h}, opts...))
}

type streamServiceHandler struct {
	StreamServiceHandler
}

func (h *streamServiceHandler) BulkCreate(ctx context.Context, stream server.Stream) error {
	return h.StreamServiceHandler.BulkCreate(ctx, &streamServiceBulkCreateStream{stream})
}

type StreamService_BulkCreateStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*ABitOfEverything, error)
}

type streamServiceBulkCreateStream struct {
	stream server.Stream
}

func (x *streamServiceBulkCreateStream) Close() error {
	return x.stream.Close()
}

func (x *streamServiceBulkCreateStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *streamServiceBulkCreateStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *streamServiceBulkCreateStream) Recv() (*ABitOfEverything, error) {
	m := new(ABitOfEverything)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (h *streamServiceHandler) List(ctx context.Context, stream server.Stream) error {
	m := new(google_protobuf2.Empty)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.StreamServiceHandler.List(ctx, m, &streamServiceListStream{stream})
}

type StreamService_ListStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*ABitOfEverything) error
}

type streamServiceListStream struct {
	stream server.Stream
}

func (x *streamServiceListStream) Close() error {
	return x.stream.Close()
}

func (x *streamServiceListStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *streamServiceListStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *streamServiceListStream) Send(m *ABitOfEverything) error {
	return x.stream.Send(m)
}

func (h *streamServiceHandler) BulkEcho(ctx context.Context, stream server.Stream) error {
	return h.StreamServiceHandler.BulkEcho(ctx, &streamServiceBulkEchoStream{stream})
}

type StreamService_BulkEchoStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*grpc_gateway_examples_sub.StringMessage) error
	Recv() (*grpc_gateway_examples_sub.StringMessage, error)
}

type streamServiceBulkEchoStream struct {
	stream server.Stream
}

func (x *streamServiceBulkEchoStream) Close() error {
	return x.stream.Close()
}

func (x *streamServiceBulkEchoStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *streamServiceBulkEchoStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *streamServiceBulkEchoStream) Send(m *grpc_gateway_examples_sub.StringMessage) error {
	return x.stream.Send(m)
}

func (x *streamServiceBulkEchoStream) Recv() (*grpc_gateway_examples_sub.StringMessage, error) {
	m := new(grpc_gateway_examples_sub.StringMessage)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
