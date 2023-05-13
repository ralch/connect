package runtimev1connect

import (
	"github.com/bufbuild/connect-go"
	"github.com/bufbuild/connect-grpcreflect-go"
	"github.com/go-chi/chi/v5"
	"github.com/ralch/connect/internal/proto/connect/runtime/v1/runtimev1connect/middleware"
)

// ServerReflectionController represents a controller for grpc.reflection.v1.ServerReflection handler and grpc.reflection.v1alpha.ServerReflection handler.
type ServerReflectionController struct {
	// Services contains a collection of service names
	Services []string
}

// Mount mounts the controller to a given router.
func (x *ServerReflectionController) Mount(r chi.Router) {
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

	// create the reflector
	reflector := grpcreflect.NewStaticReflector(x.Services...)

	r.Group(func(r chi.Router) {
		// mount the middleware
		r.Use(middleware.WithUnaryHandler(interceptors))

		// create the handler v1.0
		path, handler := grpcreflect.NewHandlerV1(reflector, options...)
		// mount the handler
		r.Mount(path, handler)
	})

	r.Group(func(r chi.Router) {
		// mount the middleware
		r.Use(middleware.WithUnaryHandler(interceptors))

		// create the handler v1.0-alpha
		path, handler := grpcreflect.NewHandlerV1Alpha(reflector, options...)
		// mount the handler
		r.Mount(path, handler)
	})
}
