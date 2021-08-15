package commands

import (
	"context"

	"github.com/google/uuid"
	"github.com/jybbang/go-core-architecture/core"
	"github.com/jybbang/nexinterface/src/books/events"
	"github.com/jybbang/nexinterface/src/entities"
)

type CreateBookCommand struct {
	Title  string
	Author string
	Price  float64
}

func CreateBookCommandHandler(ctx context.Context, request interface{}) core.Result {
	command := request.(*CreateBookCommand)

	dto := new(entities.Book)
	dto.ID = uuid.New()
	dto.Title = command.Title
	dto.Author = command.Author
	dto.Price = command.Price

	repository := core.GetRepositoryService(dto)
	repository.Add(ctx, dto)

	eventBus := core.GetEventBus()
	eventBus.AddDomainEvent(events.NewBookCreatedEvent(dto))
	defer eventBus.PublishDomainEvents(ctx)

	repository.Find(ctx, dto, dto.ID)
	return core.Result{V: dto}
}
