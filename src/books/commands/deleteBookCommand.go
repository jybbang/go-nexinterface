package commands

import (
	"context"

	"github.com/google/uuid"
	"github.com/jybbang/go-core-architecture/core"
	"github.com/jybbang/nexinterface/src/books/events"
	"github.com/jybbang/nexinterface/src/entities"
)

type DeleteBookCommand struct {
	Id string `validate:"required,uuid4"`
}

func DeleteBookCommandHandler(ctx context.Context, request interface{}) core.Result {
	command := request.(*DeleteBookCommand)

	id := uuid.MustParse(command.Id)
	repository := core.GetRepositoryService(new(entities.Book))
	result := repository.Remove(ctx, id)

	core.GetEventbus().AddDomainEvent(events.NewBookDeletedEvent(id))
	return result
}
