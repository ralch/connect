package runtime

import (
	runtimev1 "github.com/ralch/connect/internal/proto/connect/runtime/v1"
	runtimev1connect "github.com/ralch/connect/internal/proto/connect/runtime/v1/runtimev1connect"
)

type (
	// Event represents an event
	Event = runtimev1.Event
	// PushEventRequest represents a request for connect.runtime.v1.EventService.PushEvent method.
	PushEventRequest = runtimev1.PushEventRequest
	// PushEventResponse represents a response for connect.runtime.v1.EventService.PushEvent method.
	PushEventResponse = runtimev1.PushEventResponse
	// EventServiceClient is a client for the connect.runtime.v1.EventService service.
	EventServiceClient = runtimev1connect.EventServiceClient
	// EventServiceClientDiscard is a client on which all operations succeed without doing anything.
	EventServiceClientDiscard = runtimev1connect.EventServiceClientDiscard
	// EventServiceHandler is an implementation of the connect.runtime.v1.EventService service.
	EventServiceHandler = runtimev1connect.EventServiceHandler
	// EventServiceController represents a controller for connect.runtime.v1.EventServiceHandler handler.
	EventServiceController = runtimev1connect.EventServiceController
	// EventServiceClientBroker is a client broker for the connect.runtime.v1.EventServiceHandler handler.
	EventServiceClientBroker = runtimev1connect.EventServiceClientBroker
	// EventServiceHandlerBroker represents a message broker for connect.runtime.v1.EventService service.
	EventServiceHandlerBroker = runtimev1connect.EventServiceHandlerBroker
)

var (
	// NewEvent returns a new instance of connect.runtime.v1.Event message.
	NewEvent = runtimev1.NewEvent
	// NewEventServiceClient constructs a client for the connect.runtime.v1.EventService service.
	NewEventServiceClient = runtimev1connect.NewEventServiceClient
	// NewEventServiceClientBroker constructs a client broker for the connect.runtime.v1.EventService service.
	NewEventServiceClientBroker = runtimev1connect.NewEventServiceClientBroker
)

var (
	// ErrMissingTopic is returned by NewEventServiceClientBroker when the topic argument is not provided.
	ErrMissingTopic = runtimev1connect.ErrMissingTopic
	// ErrMissingProject is returned by NewEventServiceClientBroker when the project argument is not provided.
	ErrMissingProject = runtimev1connect.ErrMissingProject
)

type (
	// A message data and its attributes.
	Message = runtimev1.Message
	// PushMessageRequest represents a request for connect.runtime.v1.MessageService.PushMessage method.
	PushMessageRequest = runtimev1.PushEventRequest
	// PushMessageResponse represents a response for connect.runtime.v1.MessageService.PushMessage method.
	PushMessageResponse = runtimev1.PushMessageResponse
	// MessageServiceClient is a client for the connect.runtime.v1.MessageService service.
	MessageServiceClient = runtimev1connect.MessageServiceClient
	// MessageServiceClient is a client for the connect.runtime.v1.MessageService service.
	MessageServiceHandler = runtimev1connect.MessageServiceHandler
	// MessageServiceController represents a controller for connect.runtime.v1.MessageServiceHandler handler.
	MessageServiceController = runtimev1connect.MessageServiceController
)

// NewMessageServiceClient constructs a client for the connect.runtime.v1.MessageService service.
var NewMessageServiceClient = runtimev1connect.NewMessageServiceClient

type (
	// HealthCheckRequest represents a request for connect.runtime.v1.HealthService.Check method.
	HealthCheckRequest = runtimev1.HealthCheckRequest
	// HealthCheckResponse represents a response for connect.runtime.v1.HealthService.Check method.
	HealthCheckResponse = runtimev1.HealthCheckResponse
	// HealthServiceClient is a client for the connect.runtime.v1.HealthService service.
	HealthServiceClient = runtimev1connect.HealthServiceClient
	// HealthServiceHandler is an implementation of the connect.runtime.v1.HealthService service.
	HealthServiceHandler = runtimev1connect.HealthServiceHandler
	// HealthServiceController represents a controller for connect.runtime.v1.HealthServiceHandler handler.
	HealthServiceController = runtimev1connect.HealthServiceController
	// HealthServiceDictionary represents a map of connect.runtime.v1.HealthServiceHandler handler.
	HealthServiceDictionary = runtimev1connect.HealthServiceDictionary
)

// NewHealthServiceClient constructs a client for the connect.runtime.v1.HealthService service.
var NewHealthServiceClient = runtimev1connect.NewHealthServiceClient

type (
	// ServerReflectionController represents a controller for grpc.reflection.v1.ServerReflection handler and grpc.reflection.v1alpha.ServerReflection handler.
	ServerReflectionController = runtimev1connect.ServerReflectionController
)
