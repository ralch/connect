package service

import (
	"context"
	"fmt"

	grpchealth "github.com/bufbuild/connect-grpchealth-go"
)

// HealthStatus describes the health of a service.
type HealthStatus = grpchealth.Status

var (
	// HealthStatusUnknown indicates that the service's health state is indeterminate.
	HealthStatusUnknown = grpchealth.StatusUnknown

	// HealthStatusServing indicates that the service is ready to accept requests.
	HealthStatusServing = grpchealth.StatusServing

	// HealthStatusNotServing indicates that the process is healthy but the service is
	// not accepting requests. For example, StatusNotServing is often appropriate
	// when your primary database is down or unreachable.
	HealthStatusNotServing = grpchealth.StatusNotServing
)

//counterfeiter:generate -o ./fake/health_checker.go . HealthChecker

// A Checker reports the health of a service. It must be safe to call
// concurrently.
type HealthChecker = grpchealth.Checker

type (
	// CheckRequest is a request for the health of a service. When using protobuf,
	// Service will be a fully-qualified service name (for example,
	// "acme.ping.v1.PingService"). If the Service is an empty string, the caller
	// is asking for the health status of whole process.
	HealthCheckRequest = grpchealth.CheckRequest

	// CheckResponse reports the health of a service (or of the whole process). The
	// only valid Status values are StatusUnknown, StatusServing, and
	// StatusNotServing. When asked to report on the status of an unknown service,
	// Checkers should return a connect.CodeNotFound error.
	//
	// Often, systems monitoring health respond to errors by restarting the
	// process. They often respond to StatusNotServing by removing the process from
	// a load balancer pool.
	HealthCheckResponse = grpchealth.CheckResponse
)

var _ HealthChecker = &CompositeHealthChecker{}

// CompositeHealthChecker composes a checker for each service
type CompositeHealthChecker struct {
	registry map[string]HealthChecker
}

// NewCompositeHealthChecker creates a new CompositeHealthChecker
func NewCompositeHealthChecker() *CompositeHealthChecker {
	return &CompositeHealthChecker{
		registry: make(map[string]HealthChecker),
	}
}

// Check executes a service checker
func (x *CompositeHealthChecker) Check(ctx context.Context, r *HealthCheckRequest) (*HealthCheckResponse, error) {
	if checker, ok := x.registry[r.Service]; ok {
		return checker.Check(ctx, r)
	}

	return nil, fmt.Errorf("health checker not found for service %v", r.Service)
}

// Register registers a service checker
func (x *CompositeHealthChecker) Register(service string, checker HealthChecker) {
	x.registry[service] = checker
}
