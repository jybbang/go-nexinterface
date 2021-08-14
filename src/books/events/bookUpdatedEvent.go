package events

import (
	"github.com/jybbang/go-core-architecture/core"
	"github.com/jybbang/nexinterface/src/entities"
)

type BookUpdatedEvent struct {
	core.DomainEvent
	Source *entities.Book
}

func NewBookUpdatedEvent(source *entities.Book) *BookUpdatedEvent {
	event := new(BookUpdatedEvent)
	event.ID = source.ID
	event.Topic = "BookUpdatedEvent"
	event.Source = source

	return event
}
