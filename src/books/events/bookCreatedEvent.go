package events

import (
	"context"
	"fmt"

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

func BookCreatedEventHandler(ctx context.Context, notification interface{}) error {
	event := notification.(*BookCreatedEvent)

	fmt.Printf("%+v\n", event)

	return nil
}
