// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: connect/runtime/v1/event.proto

package runtimev1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/ralch/connect/internal/proto/connect/runtime/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// EventServiceName is the fully-qualified name of the EventService service.
	EventServiceName = "connect.runtime.v1.EventService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// EventServicePushEventProcedure is the fully-qualified name of the EventService's PushEvent RPC.
	EventServicePushEventProcedure = "/connect.runtime.v1.EventService/PushEvent"
)

// EventServiceClient is a client for the connect.runtime.v1.EventService service.
type EventServiceClient interface {
	// PushEvent pushes a given event to connect.runtime.v1.EventService service.
	PushEvent(context.Context, *connect_go.Request[v1.PushEventRequest]) (*connect_go.Response[v1.PushEventResponse], error)
}

// NewEventServiceClient constructs a client for the connect.runtime.v1.EventService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewEventServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) EventServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &eventServiceClient{
		pushEvent: connect_go.NewClient[v1.PushEventRequest, v1.PushEventResponse](
			httpClient,
			baseURL+EventServicePushEventProcedure,
			opts...,
		),
	}
}

// eventServiceClient implements EventServiceClient.
type eventServiceClient struct {
	pushEvent *connect_go.Client[v1.PushEventRequest, v1.PushEventResponse]
}

// PushEvent calls connect.runtime.v1.EventService.PushEvent.
func (c *eventServiceClient) PushEvent(ctx context.Context, req *connect_go.Request[v1.PushEventRequest]) (*connect_go.Response[v1.PushEventResponse], error) {
	return c.pushEvent.CallUnary(ctx, req)
}

// EventServiceHandler is an implementation of the connect.runtime.v1.EventService service.
type EventServiceHandler interface {
	// PushEvent pushes a given event to connect.runtime.v1.EventService service.
	PushEvent(context.Context, *connect_go.Request[v1.PushEventRequest]) (*connect_go.Response[v1.PushEventResponse], error)
}

// NewEventServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewEventServiceHandler(svc EventServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(EventServicePushEventProcedure, connect_go.NewUnaryHandler(
		EventServicePushEventProcedure,
		svc.PushEvent,
		opts...,
	))
	return "/connect.runtime.v1.EventService/", mux
}

// UnimplementedEventServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedEventServiceHandler struct{}

func (UnimplementedEventServiceHandler) PushEvent(context.Context, *connect_go.Request[v1.PushEventRequest]) (*connect_go.Response[v1.PushEventResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("connect.runtime.v1.EventService.PushEvent is not implemented"))
}