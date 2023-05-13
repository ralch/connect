package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bufbuild/connect-go"
	"github.com/ralch/slogr"
	"golang.org/x/exp/slog"
)

// WithRecover recovers the handler from any panic.
func WithRecover() connect.HandlerOption {
	err := fmt.Errorf("the system is not in a state required for the operation's execution")

	return connect.WithRecover(
		func(ctx context.Context, _ connect.Spec, _ http.Header, r any) error {
			key := slog.Group("error",
				slog.Any("panic", r),
			)
			// log the stack trace
			logger := slogr.FromContext(ctx)
			logger.With(key).Error(err.Error())
			// return the error
			return connect.NewError(connect.CodeFailedPrecondition, err)
		},
	)
}
