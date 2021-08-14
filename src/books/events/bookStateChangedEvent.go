package events

import (
	"github.com/jybbang/go-core-architecture/core"
	"github.com/jybbang/nexinterface/src/entities"
)

type BookStateChangedEvent struct {
	core.DomainEvent
	State *entities.BookState
}

func NewBookStateChangedEvent(state *entities.BookState) *BookStateChangedEvent {
	event := new(BookStateChangedEvent)
	event.ID = state.ID
	event.Topic = "BookStateChangedEvent"
	event.State = state

	return event
}
