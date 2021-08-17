package events

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jybbang/go-core-architecture/core"
	"github.com/jybbang/nexinterface/src/entities"
)

type BookDeletedEvent struct {
	core.DomainEvent
	Source *entities.Book
}

func NewBookDeletedEvent(id uuid.UUID) *BookDeletedEvent {
	event := new(BookDeletedEvent)
	event.ID = id
	event.Topic = "BookDeletedEvent"

	return event
}

func BookDeletedEventHandler(ctx context.Context, notification interface{}) error {
	fmt.Println("📘 book deleted 🎈")
	return nil
}
