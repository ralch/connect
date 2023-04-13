package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bufbuild/connect-go"
)

// WithRecoverer recovers the handler from any panic.
func WithRecoverer() connect.HandlerOption {
	return connect.WithRecover(
		func(_ context.Context, _ connect.Spec, _ http.Header, r any) error {
			return connect.NewError(connect.CodeFailedPrecondition, fmt.Errorf("panic: %v", r))
		},
	)
}
