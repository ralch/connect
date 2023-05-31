package runtimev1connect

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
	"github.com/bufbuild/connect-go"
	runtimev1 "github.com/ralch/connect/internal/proto/connect/runtime/v1"
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
		return nil, fmt.Errorf("connect: event broker requires a project id")
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
		return nil, fmt.Errorf("connect: event broker requires a topic name")
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

	// publish the message
	if _, err := x.client.Topic(x.topic).Publish(ctx, message).Get(ctx); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	response := &runtimev1.PushEventResponse{}
	// done!
	return connect.NewResponse(response), nil
}
