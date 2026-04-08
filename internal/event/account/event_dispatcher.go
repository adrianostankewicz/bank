package event

import (
	"context"
	"fmt"
)

type Event struct {
	Type        string  `json:"type"`
	Origin      string  `json:"origin"`
	Destination string  `json:"destination"`
	Amount      float64 `json:"amount"`
}

type EventHandler interface {
	Handle(ctx context.Context, event Event) (interface{}, error)
}

type EventDispatcher struct {
	handlers map[string]EventHandler
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string]EventHandler),
	}
}

func (d *EventDispatcher) Register(eventType string, handler EventHandler) {
	d.handlers[eventType] = handler
}

func (d *EventDispatcher) Dispatch(ctx context.Context, event Event) (interface{}, error) {
	handler, ok := d.handlers[event.Type]
	if !ok {
		return nil, fmt.Errorf("unsupported event: %s", event.Type)
	}
	return handler.Handle(ctx, event)
}
