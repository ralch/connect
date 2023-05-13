package runtimev1connect

import (
	"context"

	connect "github.com/bufbuild/connect-go"
	runtimev1 "github.com/ralch/connect/internal/proto/connect/runtime/v1"
)

var _ MessageServiceHandler = &EventServiceHandlerBroker{}

// EventServiceHandlerBroker represents a message broker for connect.runtime.v1.EventServiceHandler handler.
type EventServiceHandlerBroker struct {
	// EventServiceHandler contains an instance of connect.runtime.v1.EventServiceHandler handler.
	EventServiceHandler EventServiceHandler
}

// PushMessage pushes a given event to connect.runtime.v1.MessageService service.
func (x *EventServiceHandlerBroker) PushMessage(ctx context.Context, r *connect.Request[runtimev1.PushMessageRequest]) (*connect.Response[runtimev1.PushMessageResponse], error) {
	// prepare the request
	request := r.Msg

	// prepare the argument
	argument := &runtimev1.PushEventRequest{
		Event: &runtimev1.Event{},
	}

	// prepare the event
	argument.SetAttributes(request.Message.Attributes)
	argument.SetData(request.Message.Data)

	// push the event
	if _, err := x.EventServiceHandler.PushEvent(ctx, connect.NewRequest(argument)); err != nil {
		return nil, err
	}

	response := &runtimev1.PushMessageResponse{}
	// done!
	return connect.NewResponse(response), nil
}
