package events

import (
	"sync"
	"time"
)

type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() any
}

type EventHandlerInterface interface {
	Handle(event EventInterface, wg *sync.WaitGroup)
}

type EventDispatcherInterface interface {
	Register(eventName string, event EventHandlerInterface) error
	Dispatch(event EventInterface) error
	Unregister(eventName string, event EventHandlerInterface) error
	Has(eventName string, event EventHandlerInterface) bool
	Length(eventName string) int
	Clear()
}
