package infrastructure

import (
	"github.com/verbruggenjesse/grpc-consumer/domain/abstract"
	"github.com/verbruggenjesse/grpc-consumer/protos/eventstore"
)

// EventClient is a client that can handle miultiple streams of events
type EventClient struct {
	subscriber abstract.ISubscriber
	eventChan  chan *eventstore.Event
	errChan    chan error
}

// NewEventClient is the constructor for EventClient
func NewEventClient(subscriber abstract.ISubscriber) *EventClient {
	return &EventClient{
		subscriber: subscriber,
		eventChan:  make(chan *eventstore.Event),
		errChan:    make(chan error),
	}
}

// Subscribe will fetch events specified by a subscription object
func (e *EventClient) Subscribe(subscription abstract.ISubscription) {
	messageChan := make(chan abstract.IMessage)

	go func() {
		for message := range messageChan {
			event := &eventstore.Event{
				Key:     message.Key(),
				Id:      message.ID(),
				Payload: []byte(message.Values()["payload"].(string)),
			}

			if subscription.IncludeMetadata() {
				event.Metadata = []byte(message.Values()["metadata"].(string))
			}

			e.eventChan <- event
		}
	}()

	go e.subscriber.Subscribe(subscription, &messageChan, &e.errChan)
}

// EventChan is used to transport events
func (e *EventClient) EventChan() *chan *eventstore.Event {
	return &e.eventChan
}

// ErrorChan is used to transport errors
func (e *EventClient) ErrorChan() *chan error {
	return &e.errChan
}
