package queries

import (
	"context"

	"github.com/google/uuid"
	"github.com/jybbang/go-core-architecture/core"
	"github.com/jybbang/nexinterface/src/entities"
)

type GetBookQuery struct {
	Id string `validate:"required,uuid4"`
}

func GetBookQueryHandler(ctx context.Context, request interface{}) core.Result {
	query := request.(*GetBookQuery)
	dto := new(entities.Book)
	repository := core.GetRepositoryService(new(entities.Book))
	return repository.Find(ctx, uuid.MustParse(query.Id), dto)

}
