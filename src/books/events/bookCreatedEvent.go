package events

import (
	"github.com/jybbang/go-core-architecture/core"
	"github.com/jybbang/nexinterface/src/entities"
)

type BookCreatedEvent struct {
	core.DomainEvent
	Source *entities.Book
}

func NewBookCreatedEvent(source *entities.Book) *BookCreatedEvent {
	event := new(BookCreatedEvent)
	event.ID = source.ID
	event.Topic = "BookCreatedEvent"
	event.Source = source

	return event
}

func BookCreatedEventHandler(notification interface{}) {
	event, ok := notification.(*BookCreatedEvent)
	if !ok {
		panic(core.ErrConflict.Error())
	}
	core.Log.Info("event handler", event.Topic)
}
