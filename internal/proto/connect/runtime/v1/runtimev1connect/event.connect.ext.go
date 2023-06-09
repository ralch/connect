package runtimev1connect

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
	"github.com/bufbuild/connect-go"
	runtimev1 "github.com/ralch/connect/internal/proto/connect/runtime/v1"
	"github.com/ralch/slogr"
	"golang.org/x/exp/slog"
)

var (
	// ErrMissingTopic is returned by NewEventServiceClientBroker when the topic argument is not provided.
	ErrMissingTopic = fmt.Errorf("no topic")
	// ErrMissingProject is returned by NewEventServiceClientBroker when the project argument is not provided.
	ErrMissingProject = fmt.Errorf("no project")
)

var _ EventServiceClient = &EventServiceClientBroker{}

// EventServiceClientBroker is a client broker for the connect.runtime.v1.EventService service.
type EventServiceClientBroker struct {
	client *pubsub.Client
	topic  string
}

// NewEventServiceClientBroker constructs a client broker for the connect.runtime.v1.EventService service.
func NewEventServiceClientBroker(ctx context.Context, project string, opts ...Option) (EventServiceClient, error) {
	if project == "" {
		return nil, ErrMissingProject
	}

	// prepare the client
	client, err := pubsub.NewClient(ctx, project)
	if err != nil {
		return nil, err
	}

	// prepare the broker
	broker := &EventServiceClientBroker{
		client: client,
	}

	// apply the options to the broker
	for _, opt := range opts {
		opt.Apply(broker)
	}

	if broker.topic == "" {
		return nil, ErrMissingTopic
	}

	// done!
	return broker, nil
}

// PushEvent pushes a given event to connect.runtime.v1.EventService service.
func (x *EventServiceClientBroker) PushEvent(ctx context.Context, r *connect.Request[runtimev1.PushEventRequest]) (*connect.Response[runtimev1.PushEventResponse], error) {
	// prepare the request
	request := r.Msg

	// prepare the message
	message := &pubsub.Message{
		Data:        request.GetData(),
		Attributes:  request.GetAttributes(),
		OrderingKey: request.GetOrderingKey(),
	}

	// prepare the logger attr
	attr := slog.Group("event",
		slog.String("id", request.Event.Id),
		slog.String("type", request.Event.GetType()),
		slog.String("source", request.Event.GetSource()),
		slog.String("subject", request.Event.GetSubject()),
	)

	logger := slogr.FromContext(ctx)
	// prepare the logger message
	logger.Info("push an event", attr)

	topic := x.client.Topic(x.topic)
	topic.EnableMessageOrdering = true
	// publish the message
	if _, err := topic.Publish(ctx, message).Get(ctx); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	response := &runtimev1.PushEventResponse{}
	// done!
	return connect.NewResponse(response), nil
}

var _ EventServiceClient = &EventServiceClientDiscard{}

// EventServiceClientDiscard is a client on which all operations succeed without doing anything.
type EventServiceClientDiscard struct{}

// PushEvent pushes a given event to connect.runtime.v1.EventService service.
func (*EventServiceClientDiscard) PushEvent(ctx context.Context, r *connect.Request[runtimev1.PushEventRequest]) (*connect.Response[runtimev1.PushEventResponse], error) {
	// prepare the request
	request := r.Msg
	// prepare the logger attr
	attr := slog.Group("event",
		slog.String("id", request.Event.Id),
		slog.String("type", request.Event.GetType()),
		slog.String("source", request.Event.GetSource()),
		slog.String("subject", request.Event.GetSubject()),
	)

	logger := slogr.FromContext(ctx)
	// prepare the logger message
	logger.Info("push an event", attr)

	response := &runtimev1.PushEventResponse{}
	// done!
	return connect.NewResponse(response), nil

}
