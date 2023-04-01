package health

import (
	"context"
	"fmt"

	"github.com/bufbuild/connect-grpchealth-go"
)

// Status describes the health of a service.
type Status = grpchealth.Status

var (
	// StatusUnknown indicates that the service's health state is indeterminate.
	StatusUnknown = grpchealth.StatusUnknown

	// StatusServing indicates that the service is ready to accept requests.
	StatusServing = grpchealth.StatusServing

	// StatusNotServing indicates that the process is healthy but the service is
	// not accepting requests. For example, StatusNotServing is often appropriate
	// when your primary database is down or unreachable.
	StatusNotServing = grpchealth.StatusNotServing
)

//counterfeiter:generate -o ./fake/fake_checker.go . Checker

// A Checker reports the health of a service. It must be safe to call
// concurrently.
type Checker = grpchealth.Checker

type (
	// CheckRequest is a request for the health of a service. When using protobuf,
	// Service will be a fully-qualified service name (for example,
	// "acme.ping.v1.PingService"). If the Service is an empty string, the caller
	// is asking for the health status of whole process.
	CheckRequest = grpchealth.CheckRequest

	// CheckResponse reports the health of a service (or of the whole process). The
	// only valid Status values are StatusUnknown, StatusServing, and
	// StatusNotServing. When asked to report on the status of an unknown service,
	// Checkers should return a connect.CodeNotFound error.
	//
	// Often, systems monitoring health respond to errors by restarting the
	// process. They often respond to StatusNotServing by removing the process from
	// a load balancer pool.
	CheckResponse = grpchealth.CheckResponse
)

var _ Checker = &ServiceChecker{}

// ServiceChecker composes a checker for each service
type ServiceChecker struct {
	registry map[string]Checker
}

// NewServiceChecker creates a new ServiceChecker
func NewServiceChecker() *ServiceChecker {
	return &ServiceChecker{
		registry: make(map[string]Checker),
	}
}

// Check executes a service checker
func (x *ServiceChecker) Check(ctx context.Context, r *CheckRequest) (*CheckResponse, error) {
	if checker, ok := x.registry[r.Service]; ok {
		return checker.Check(ctx, r)
	}

	return nil, fmt.Errorf("health checker not found for service %v", r.Service)
}

// Register registers a service checker
func (x *ServiceChecker) Register(service string, checker Checker) {
	x.registry[service] = checker
}
