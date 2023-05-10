package middleware

import (
	"github.com/bufbuild/connect-go"
	otelconnect "github.com/bufbuild/connect-opentelemetry-go"
)

// WithTracer set up the Open Telemetry Tracer.
func WithTracer() connect.Interceptor {
	return otelconnect.NewInterceptor()
}
