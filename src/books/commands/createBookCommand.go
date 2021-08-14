package commands

import (
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

func CreateBookCommandHandler(request interface{}) interface{} {
	command, ok := request.(*CreateBookCommand)
	if !ok {
		panic(core.ErrConflict.Error())
	}

	dto := new(entities.Book)
	dto.ID = uuid.New()
	dto.Title = command.Title
	dto.Author = command.Author
	dto.Price = command.Price

	repository := core.GetRepositoryService(dto)
	repository.Add(dto)

	eventBus := core.GetEventBus()
	eventBus.AddDomainEvent(events.NewBookCreatedEvent(dto))
	defer eventBus.PublishDomainEvents()

	repository.Find(dto, dto.ID)
	return dto
}
