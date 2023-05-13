package middleware

import (
	connect "github.com/bufbuild/connect-go"
	connectotel "github.com/bufbuild/connect-opentelemetry-go"
)

// WithTracer set up the Open Telemetry Tracer.
func WithTracer() connect.Interceptor {
	return connectotel.NewInterceptor()
}
