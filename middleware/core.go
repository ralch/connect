package middleware

import (
	"net/http"

	"github.com/bufbuild/connect-go"
)

// UnaryHandlerFunc represent the handler.
type UnaryHandlerFunc func(http.Handler) http.Handler

// UnaryInterceptorFunc is a simple Interceptor implementation that only
// wraps unary RPCs. It has no effect on streaming RPCs.
type UnaryInterceptorFunc = connect.UnaryInterceptorFunc

// UnaryHandler represents an unary handler.
type UnaryHandler struct {
	UnaryHandler     UnaryHandlerFunc
	UnaryInterceptor UnaryInterceptorFunc
}

// WrapHTTPHandler implements UnaryInterceptorFunc
func (x *UnaryHandler) WrapHTTPHandler(fn http.Handler) http.Handler {
	return x.UnaryHandler(fn)
}

// WrapStreamingClient implements connect.Interceptor
func (x *UnaryHandler) WrapStreamingClient(fn connect.StreamingClientFunc) connect.StreamingClientFunc {
	return x.UnaryInterceptor.WrapStreamingClient(fn)
}

// WrapStreamingHandler implements connect.Interceptor
func (x *UnaryHandler) WrapStreamingHandler(fn connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return x.UnaryInterceptor.WrapStreamingHandler(fn)
}

// WrapUnary implements connect.Interceptor
func (x *UnaryHandler) WrapUnary(fn connect.UnaryFunc) connect.UnaryFunc {
	return x.UnaryInterceptor.WrapUnary(fn)
}
