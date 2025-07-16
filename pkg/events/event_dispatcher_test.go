package events_test

import (
	"errors"
	"testing"
	"time"

	"github.com/luizmoitinho/events_utils/pkg/events"
	"github.com/stretchr/testify/suite"
)

type TestEvent struct {
	Name    string
	Payload any
}

func (e *TestEvent) GetName() string {
	return e.Name
}

func (e *TestEvent) GetPayload() any {
	return e.Payload
}

func (e *TestEvent) GetDateTime() time.Time {
	return time.Now()
}

type TestEventHandler struct {
	ID int
}

func (h *TestEventHandler) Handle(event events.EventInterface) {}

type EventDispatcherTestSuite struct {
	suite.Suite

	event  TestEvent
	event2 TestEvent

	handler  TestEventHandler
	handler2 TestEventHandler
	handler3 TestEventHandler

	eventDispatcher events.EventDispatcherInterface
}

func (s *EventDispatcherTestSuite) SetupTest() {
	s.eventDispatcher = events.NewEventDispatcher()

	s.handler = TestEventHandler{ID: 1}
	s.handler2 = TestEventHandler{ID: 2}
	s.handler3 = TestEventHandler{ID: 3}

	s.event = TestEvent{Name: "test", Payload: "test"}
	s.event2 = TestEvent{Name: "test-2", Payload: "test-2"}
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	err := s.eventDispatcher.Register(s.event.GetName(), &s.handler)
	s.Nil(err)
	s.Equal(true, s.eventDispatcher.Has(s.event.GetName(), &s.handler))

	err = s.eventDispatcher.Register(s.event.GetName(), &s.handler2)
	s.Nil(err)
	s.Equal(true, s.eventDispatcher.Has(s.event.GetName(), &s.handler2))

	s.Equal(2, s.eventDispatcher.Length(s.event.GetName()))
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Register_WithSamHandler() {
	err := s.eventDispatcher.Register(s.event.GetName(), &s.handler)
	s.Nil(err)
	s.Equal(true, s.eventDispatcher.Has(s.event.GetName(), &s.handler))

	s.Equal(1, s.eventDispatcher.Length(s.event.GetName()))

	err = s.eventDispatcher.Register(s.event.GetName(), &s.handler)
	s.Equal(errors.New("handler already registered"), err)

	s.Equal(1, s.eventDispatcher.Length(s.event.GetName()))
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Clear() {
	//Event one
	err := s.eventDispatcher.Register(s.event.GetName(), &s.handler)
	s.Nil(err)
	s.Equal(true, s.eventDispatcher.Has(s.event.GetName(), &s.handler))
	s.Equal(1, s.eventDispatcher.Length(s.event.GetName()))

	err = s.eventDispatcher.Register(s.event.GetName(), &s.handler2)
	s.Nil(err)
	s.Equal(true, s.eventDispatcher.Has(s.event.GetName(), &s.handler2))
	s.Equal(2, s.eventDispatcher.Length(s.event.GetName()))

	//Event two
	err = s.eventDispatcher.Register(s.event2.GetName(), &s.handler3)
	s.Nil(err)
	s.Equal(1, s.eventDispatcher.Length(s.event2.GetName()))

	//act
	s.eventDispatcher.Clear()

	s.Nil(err)
	s.Equal(false, s.eventDispatcher.Has(s.event.GetName(), &s.handler))
	s.Equal(false, s.eventDispatcher.Has(s.event2.GetName(), &s.handler2))

	s.Equal(0, s.eventDispatcher.Length(s.event.GetName()))
	s.Equal(0, s.eventDispatcher.Length(s.event2.GetName()))
}
