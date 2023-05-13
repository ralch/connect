package runtimev1connect

import (
	"github.com/bufbuild/connect-go"
	"github.com/go-chi/chi/v5"
	"github.com/ralch/connect/internal/proto/connect/runtime/v1/runtimev1connect/middleware"
)

// MessageServiceController represents a controller for connect.runtime.v1.MessageServiceHandler handler.
type MessageServiceController struct {
	// MessageServiceHandler contains an instance of connect.runtime.v1.MessageServiceHandler handler.
	MessageServiceHandler MessageServiceHandler
}

// Mount mounts the controller to a given router.
func (x *MessageServiceController) Mount(r chi.Router) {
	var interceptors []connect.Interceptor
	// prepare the interceptors
	interceptors = append(interceptors, middleware.WithTracer())
	interceptors = append(interceptors, middleware.WithLogger())
	interceptors = append(interceptors, middleware.WithValidator())

	var options []connect.HandlerOption
	// prepare the options
	options = append(options, middleware.WithRecover())
	// prepare the options for interceptor collection
	options = append(options, connect.WithInterceptors(interceptors...))

	r.Group(func(r chi.Router) {
		// mount the middleware
		r.Use(middleware.WithUnaryHandler(interceptors))
		// create the handler
		path, handler := NewMessageServiceHandler(x.MessageServiceHandler, options...)
		// mount the handler
		r.Mount(path, handler)
	})
}
