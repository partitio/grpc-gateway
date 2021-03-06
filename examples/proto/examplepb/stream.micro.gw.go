// Code generated by protoc-gen-micro-gateway. DO NOT EDIT.
// source: examples/proto/examplepb/stream.proto

/*
Package examplepb is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package examplepb

import (
	"io"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/partitio/micro-gateway/examples/proto/sub"
	"github.com/partitio/micro-gateway/runtime"
	"github.com/partitio/micro-gateway/utilities"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray

func request_StreamService_BulkCreate_0(ctx context.Context, marshaler runtime.Marshaler, client StreamService, req *http.Request, pathParams map[string]string) (proto.Message, error) {
	stream, err := client.BulkCreate(ctx)
	if err != nil {
		// grpclog.Infof("Failed to start streaming: %v", err)
		return nil, err
	}
	defer stream.Close()
	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	dec := marshaler.NewDecoder(newReader())
	for {
		var protoReq ABitOfEverything
		err = dec.Decode(&protoReq)
		if err == io.EOF {
			break
		}
		if err != nil {
			// grpclog.Infof("Failed to decode request: %v", err)
			return nil, status.Errorf(codes.InvalidArgument, "%v", err)
		}
		if err = stream.Send(&protoReq); err != nil {
			// grpclog.Infof("Failed to send request: %v", err)
			return nil, err
		}
	}

	if err := stream.Close(); err != nil {
		// grpclog.Infof("Failed to terminate client stream: %v", err)
		return nil, err
	}

	var msg *ABitOfEverything
	return msg, stream.RecvMsg(msg)

}

func request_StreamService_List_0(ctx context.Context, marshaler runtime.Marshaler, client StreamService, req *http.Request, pathParams map[string]string) (StreamService_ListService, error) {
	var protoReq empty.Empty

	stream, err := client.List(ctx, &protoReq)
	if err != nil {
		return nil, err
	}
	return stream, nil

}

func request_StreamService_BulkEcho_0(ctx context.Context, marshaler runtime.Marshaler, client StreamService, req *http.Request, pathParams map[string]string) (StreamService_BulkEchoService, error) {
	stream, err := client.BulkEcho(ctx)
	if err != nil {
		// grpclog.Infof("Failed to start streaming: %v", err)
		return nil, err
	}
	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, berr
	}
	dec := marshaler.NewDecoder(newReader())
	handleSend := func() error {
		var protoReq sub.StringMessage
		err := dec.Decode(&protoReq)
		if err == io.EOF {
			return err
		}
		if err != nil {
			// grpclog.Infof("Failed to decode request: %v", err)
			return err
		}
		if err := stream.Send(&protoReq); err != nil {
			// grpclog.Infof("Failed to send request: %v", err)
			return err
		}
		return nil
	}
	if err := handleSend(); err != nil {
		if cerr := stream.Close(); cerr != nil {
			// grpclog.Infof("Failed to terminate client stream: %v", cerr)
		}
		if err == io.EOF {
			return stream, nil
		}
		return nil, err
	}
	go func() {
		for {
			if err := handleSend(); err != nil {
				break
			}
		}
		if err := stream.Close(); err != nil {
			// grpclog.Infof("Failed to terminate client stream: %v", err)
		}
	}()
	return stream, nil
}

// RegisterStreamServiceHandler registers the http handlers for service StreamService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
// func RegisterStreamServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
//  	return RegisterStreamServiceHandlerClient(ctx, mux, NewStreamServiceClient(conn))
//}

// RegisterStreamServiceHandlerService registers the http handlers for service StreamService
// to "mux". The handlers forward requests to the micro service endpoint over the given implementation of "StreamService".
func RegisterStreamServiceHandlerService(ctx context.Context, mux *runtime.ServeMux, client StreamService) error {

	mux.Handle("POST", pattern_StreamService_BulkCreate_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, err := request_StreamService_BulkCreate_0(rctx, inboundMarshaler, client, req, pathParams)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_StreamService_BulkCreate_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_StreamService_List_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, err := request_StreamService_List_0(rctx, inboundMarshaler, client, req, pathParams)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_StreamService_List_0(ctx, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_StreamService_BulkEcho_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, err := request_StreamService_BulkEcho_0(rctx, inboundMarshaler, client, req, pathParams)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_StreamService_BulkEcho_0(ctx, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_StreamService_BulkCreate_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"v1", "example", "a_bit_of_everything", "bulk"}, ""))

	pattern_StreamService_List_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "example", "a_bit_of_everything"}, ""))

	pattern_StreamService_BulkEcho_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"v1", "example", "a_bit_of_everything", "echo"}, ""))
)

var (
	forward_StreamService_BulkCreate_0 = runtime.ForwardResponseMessage

	forward_StreamService_List_0 = runtime.ForwardResponseStream

	forward_StreamService_BulkEcho_0 = runtime.ForwardResponseStream
)
