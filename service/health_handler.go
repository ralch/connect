package service

import (
	"net/http"

	"github.com/bufbuild/connect-grpchealth-go"
	"github.com/go-chi/chi/v5"
)

// HealthHandler wraps the https://github.com/bufbuild/connect-grpchealth-go handler
type HealthHandler struct {
	path    string
	handler http.Handler
}

// NewServiceHealthHandler creates a new service handler
func NewHealthHandler(checker HealthChecker) *HealthHandler {
	path, handler := grpchealth.NewHandler(checker)

	return &HealthHandler{
		path:    path,
		handler: handler,
	}
}

// Mount mounds the handler
func (x *HealthHandler) Mount(r chi.Router) {
	r.Mount(x.path, x.handler)
}
