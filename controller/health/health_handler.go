package health

import (
	"net/http"

	"github.com/bufbuild/connect-grpchealth-go"
	"github.com/go-chi/chi/v5"
)

// ServiceHandler wraps the https://github.com/bufbuild/connect-grpchealth-go handler
type ServiceHandler struct {
	path    string
	handler http.Handler
}

// NewServiceHandler creates a new service handler
func NewServiceHandler(checker Checker) *ServiceHandler {
	path, handler := grpchealth.NewHandler(checker)

	return &ServiceHandler{
		path:    path,
		handler: handler,
	}
}

// Mount mounds the handler
func (x *ServiceHandler) Mount(r chi.Router) {
	r.Mount(x.path, x.handler)
}
