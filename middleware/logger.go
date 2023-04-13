package middleware

import (
	"context"
	"net/http"

	"github.com/bufbuild/connect-go"
	"github.com/gofrs/uuid"
	"github.com/ralch/stack"
)

// WithLogger set up the logger.
func WithLogger() *UnaryHandler {
	interFn := func(next connect.UnaryFunc) connect.UnaryFunc {
		// prepare the callback
		fn := func(ctx context.Context, request connect.AnyRequest) (connect.AnyResponse, error) {
			logger := stack.FromContext(ctx)

			var (
				id        = uuid.Must(uuid.NewV4()).String()
				procedure = request.Spec().Procedure
			)

			// log the start
			logger.InfoCtx(ctx, "execution started",
				stack.OperationStart(id, procedure),
			)

			// prepare the context
			ctx = stack.WithContext(ctx, logger.With(
				stack.OperationContinue(id, procedure),
			))

			// execute the method
			response, err := next(ctx, request)
			if err != nil {
				// log the end
				logger.ErrorCtx(ctx, "execution finished",
					stack.OperationEnd(id, procedure),
					stack.Error(err),
				)
			} else {
				// log the end
				logger.InfoCtx(ctx, "execution finished",
					stack.OperationEnd(id, procedure),
				)
			}

			return response, err
		}
		// done!
		return fn
	}

	handleFn := func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			w = &ResponseWriter{ResponseWriter: w}

			ctx := r.Context()

			logger := stack.FromContext(ctx)
			// log the start
			logger = logger.With(stack.Request(r))
			logger.InfoCtx(ctx, "request received")

			// prepare the context
			ctx = stack.WithContext(ctx, logger)
			r = r.WithContext(ctx)

			// execute the handler
			next.ServeHTTP(w, r)

			// log the end
			logger = logger.With(stack.ResponseWriter(w))
			logger.InfoCtx(ctx, "request completed")
		}

		return http.HandlerFunc(fn)
	}

	return &UnaryHandler{
		UnaryHandler:     UnaryHandlerFunc(handleFn),
		UnaryInterceptor: UnaryInterceptorFunc(interFn),
	}
}

var _ http.ResponseWriter = &ResponseWriter{}

type ResponseWriter struct {
	StatusCode     int32
	ContentLength  int64
	ResponseWriter http.ResponseWriter
}

// Header implements http.ResponseWriter
func (r *ResponseWriter) Header() http.Header {
	return r.ResponseWriter.Header()
}

// Write implements http.ResponseWriter
func (r *ResponseWriter) Write(data []byte) (int, error) {
	n, err := r.ResponseWriter.Write(data)
	r.ContentLength = r.ContentLength + int64(n)
	return n, err
}

// WriteHeader implements http.ResponseWriter
func (r *ResponseWriter) WriteHeader(code int) {
	r.StatusCode = int32(code)
	r.ResponseWriter.WriteHeader(code)
}

func (r *ResponseWriter) GetStatusCode() int32 {
	return r.StatusCode
}

func (r *ResponseWriter) GetContentLength() int64 {
	return r.ContentLength
}
