package runtimev1connect

import (
	"context"
	"fmt"

	"github.com/bufbuild/connect-go"
	runtimev1 "github.com/ralch/connect/internal/proto/connect/runtime/v1"
)

var _ HealthServiceHandler = &HealthServiceDictionary{}

// HealthServiceHandlerDictionary represents a map of connect.runtime.v1.HealthServiceHandler handler.
type HealthServiceDictionary map[string]HealthServiceHandler

// Check checks the health of a given service.
func (x *HealthServiceDictionary) Check(ctx context.Context, r *connect.Request[runtimev1.HealthCheckRequest]) (*connect.Response[runtimev1.HealthCheckResponse], error) {
	if h, ok := (*x)[r.Msg.Service]; ok {
		return h.Check(ctx, r)
	}

	return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("not found a registered instance of connect.runtime.v1.HealthServiceHandler for %v", r.Msg.Service))
}

// Watch watches the health of a given service.
func (x *HealthServiceDictionary) Watch(ctx context.Context, r *connect.Request[runtimev1.HealthCheckRequest], s *connect.ServerStream[runtimev1.HealthCheckResponse]) error {
	if h, ok := (*x)[r.Msg.Service]; ok {
		return h.Watch(ctx, r, s)
	}

	return connect.NewError(connect.CodeNotFound, fmt.Errorf("not found a registered instance of connect.runtime.v1.HealthServiceHandler for %v", r.Msg.Service))
}
