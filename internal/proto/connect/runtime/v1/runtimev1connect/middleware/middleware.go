package middleware

import (
	"net/http"

	"github.com/bufbuild/connect-go"
)

// UnaryHandlerFunc represent the handler.
type UnaryHandlerFunc = func(http.Handler) http.Handler

// UnaryInterceptorFunc is a simple Interceptor implementation that only
// wraps unary RPCs. It has no effect on streaming RPCs.
type UnaryInterceptorFunc = connect.UnaryInterceptorFunc

// UnaryInterceptor represents an unary handler.
type UnaryInterceptor struct {
	UnaryHandler     UnaryHandlerFunc
	UnaryInterceptor UnaryInterceptorFunc
}

// WrapStreamingClient implements connect.Interceptor
func (x *UnaryInterceptor) WrapStreamingClient(fn connect.StreamingClientFunc) connect.StreamingClientFunc {
	return x.UnaryInterceptor.WrapStreamingClient(fn)
}

// WrapStreamingHandler implements connect.Interceptor
func (x *UnaryInterceptor) WrapStreamingHandler(fn connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return x.UnaryInterceptor.WrapStreamingHandler(fn)
}

// WrapUnary implements connect.Interceptor
func (x *UnaryInterceptor) WrapUnary(fn connect.UnaryFunc) connect.UnaryFunc {
	return x.UnaryInterceptor.WrapUnary(fn)
}

// WrapUnaryHandler implements UnaryHandlerFunc.
func (x *UnaryInterceptor) WrapUnaryHandler(fn http.Handler) http.Handler {
	return x.UnaryHandler(fn)
}

// WithUnaryHandler represents an unary handler
func WithUnaryHandler(collection []connect.Interceptor) UnaryHandlerFunc {
	fn := func(handler http.Handler) http.Handler {
		for _, item := range collection {
			if interceptor, ok := item.(*UnaryInterceptor); ok {
				// use the handler
				handler = interceptor.WrapUnaryHandler(handler)
			}
		}

		return handler
	}

	return fn
}
