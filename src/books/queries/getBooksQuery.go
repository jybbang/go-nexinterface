package queries

import (
	"context"

	"github.com/jybbang/go-core-architecture/core"
	"github.com/jybbang/nexinterface/src/entities"
)

type GetBooksQuery struct {
}

func GetBooksQueryHandler(ctx context.Context, request interface{}) core.Result {
	repository := core.GetRepositoryService(new(entities.Book))
	dtos := make([]*entities.Book, 0)
	return repository.List(ctx, &dtos)
}
