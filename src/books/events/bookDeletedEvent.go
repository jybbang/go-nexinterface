package events

import (
	"github.com/jybbang/go-core-architecture/core"
	"github.com/jybbang/nexinterface/src/entities"
)

type BookDeletedEvent struct {
	core.DomainEvent
	Source *entities.Book
}

func NewBookDeletedEvent(source *entities.Book) *BookDeletedEvent {
	event := new(BookDeletedEvent)
	event.ID = source.ID
	event.Topic = "BookDeletedEvent"
	event.Source = source

	return event
}
