package runtimev1connect

//go:generate counterfeiter -generate

// Option represents an option for connect.runtime.v1.Option
type Option interface {
	Apply(interface{})
}

var _ Option = OptionFunc(nil)

// ConfigFunc represents a function that is an option for connect.runtime.v1.
type OptionFunc func(interface{})

// Apply applies the options to the client.
func (x OptionFunc) Apply(args interface{}) {
	x(args)
}

// WithTopic returns a option that sets the connect.runtime.v1.EventService service broker topic.
func WithTopic(name string) OptionFunc {
	fn := func(args interface{}) {
		if broker, ok := args.(*EventServiceClientBroker); ok {
			broker.topic = name
		}
	}

	return OptionFunc(fn)
}
