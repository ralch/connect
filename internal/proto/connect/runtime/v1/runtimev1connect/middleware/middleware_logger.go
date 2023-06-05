package middleware

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/gofrs/uuid"
	"github.com/ralch/slogr"
)

// WithLogger set up the logger.
func WithLogger() *UnaryInterceptor {
	handleFn := func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			start := time.Now()

			ww := &ResponseWriter{ResponseWriter: w}
			// prepare the logger
			logger := slogr.FromContext(ctx)
			logger = logger.With(slogr.Request(r))

			// prepare the context
			ctx = slogr.WithContext(ctx, logger)
			// prepare the request
			r = r.WithContext(ctx)

			// execute the handler
			next.ServeHTTP(ww, r)

			duration := time.Now().Sub(start)
			// log the request end
			logger = logger.With(slogr.ResponseWriter(ww, slogr.WithLatency(duration)))

			status := ww.GetStatusCode()
			switch {
			case status < 400:
				logger.InfoCtx(ctx, "")
			case status < 500:
				logger.WarnCtx(ctx, "")
			default:
				logger.ErrorCtx(ctx, "")
			}
		}

		return http.HandlerFunc(fn)
	}

	interFn := func(next connect.UnaryFunc) connect.UnaryFunc {
		// prepare the callback
		fn := func(ctx context.Context, request connect.AnyRequest) (connect.AnyResponse, error) {
			logger := slogr.FromContext(ctx)

			var (
				id   = uuid.Must(uuid.NewV4()).String()
				spec = request.Spec()
				peer = request.Peer()
			)

			uri := &url.URL{
				Scheme: peer.Protocol,
				Host:   peer.Addr,
				Path:   spec.Procedure,
			}

			message := fmt.Sprintf("[%v] %v", spec.StreamType, uri)

			// log the start
			logger.InfoCtx(ctx, message,
				slogr.OperationStart(id, spec.Procedure),
			)

			// prepare the context
			ctx = slogr.WithContext(ctx, logger.With(
				slogr.OperationContinue(id, spec.Procedure),
			))

			// execute the method
			response, err := next(ctx, request)
			if err == nil {
				// log the end
				logger.InfoCtx(ctx, message,
					slogr.OperationEnd(id, spec.Procedure),
				)
			} else {
				// log the end
				logger.ErrorCtx(ctx, message,
					slogr.OperationEnd(id, spec.Procedure),
					slogr.Error(err),
				)
			}

			return response, err
		}
		// done!
		return fn
	}

	return &UnaryInterceptor{
		UnaryHandler:     UnaryHandlerFunc(handleFn),
		UnaryInterceptor: UnaryInterceptorFunc(interFn),
	}
}

var _ http.ResponseWriter = &ResponseWriter{}

// ResponseWriter repersents a response writer.
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

// GetStatusCode returns the StatusCode.
func (r *ResponseWriter) GetStatusCode() int32 {
	return r.StatusCode
}

// GetContentLength returns the ContentLength.
func (r *ResponseWriter) GetContentLength() int64 {
	return r.ContentLength
}
