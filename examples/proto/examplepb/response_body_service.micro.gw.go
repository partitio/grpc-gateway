// Code generated by protoc-gen-micro-gateway. DO NOT EDIT.
// source: examples/proto/examplepb/response_body_service.proto

/*
Package examplepb is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package examplepb

import (
	"io"
	"net/http"

	"github.com/golang/protobuf/proto"
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

func request_ResponseBodyService_GetResponseBody_0(ctx context.Context, marshaler runtime.Marshaler, client ResponseBodyService, req *http.Request, pathParams map[string]string) (proto.Message, error) {
	var protoReq ResponseBodyIn

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["data"]
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "missing parameter %s", "data")
	}

	protoReq.Data, err = runtime.String(val)

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "data", err)
	}

	msg, err := client.GetResponseBody(ctx, &protoReq)
	return msg, err

}

// RegisterResponseBodyServiceHandler registers the http handlers for service ResponseBodyService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
// func RegisterResponseBodyServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
//  	return RegisterResponseBodyServiceHandlerClient(ctx, mux, NewResponseBodyServiceClient(conn))
//}

// RegisterResponseBodyServiceHandlerService registers the http handlers for service ResponseBodyService
// to "mux". The handlers forward requests to the micro service endpoint over the given implementation of "ResponseBodyService".
func RegisterResponseBodyServiceHandlerService(ctx context.Context, mux *runtime.ServeMux, client ResponseBodyService) error {

	mux.Handle("GET", pattern_ResponseBodyService_GetResponseBody_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, err := request_ResponseBodyService_GetResponseBody_0(rctx, inboundMarshaler, client, req, pathParams)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ResponseBodyService_GetResponseBody_0(ctx, mux, outboundMarshaler, w, req, response_ResponseBodyService_GetResponseBody_0{resp}, mux.GetForwardResponseOptions()...)

	})

	return nil
}

type response_ResponseBodyService_GetResponseBody_0 struct {
	proto.Message
}

func (m response_ResponseBodyService_GetResponseBody_0) XXX_ResponseBody() interface{} {
	response := m.Message.(*ResponseBodyOut)
	return response.Response
}

var (
	pattern_ResponseBodyService_GetResponseBody_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 1, 0, 4, 1, 5, 1}, []string{"responsebody", "data"}, ""))
)

var (
	forward_ResponseBodyService_GetResponseBody_0 = runtime.ForwardResponseMessage
)