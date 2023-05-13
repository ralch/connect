package runtime

import (
	"github.com/ralch/connect/internal/proto/connect/runtime/v1/runtimev1connect"
	"github.com/ralch/connect/internal/proto/connect/runtime/v1/runtimev1connect/middleware"
)

var (
	// WithTopic returns a option that sets the connect.runtime.v1.EventService service broker topic.
	WithTopic = runtimev1connect.WithTopic
	// WithRecover recovers the handler from any panic.
	WithRecover = middleware.WithRecover
)

var (
	// WithTracer set up the Open Telemetry Tracer.
	WithTracer = middleware.WithTracer
	// WithLogger set up the logger.
	WithLogger = middleware.WithLogger
	// WithValidator set up the request validator
	WithValidator = middleware.WithValidator
)

// WithUnaryHandler represents an unary handler
var WithUnaryHandler = middleware.WithUnaryHandler
