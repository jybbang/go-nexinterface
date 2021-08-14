package entities

import (
	"github.com/google/uuid"
	"github.com/jybbang/go-core-architecture/core"
)

type Book struct {
	core.Entity
	Title  string  `bson:"title,omitempty"`
	Author string  `bson:"author,omitempty"`
	Price  float64 `bson:"price,omitempty"`
}

type BookState struct {
	ID       uuid.UUID
	Stock    int
	Discount float64
}

func (b *Book) CopyWith(src interface{}) bool {
	if source, ok := src.(*Book); ok {
		b.ID = source.ID
		b.CreateUser = source.CreateUser
		b.UpdateUser = source.UpdateUser
		b.CreatedAt = source.CreatedAt
		b.UpdatedAt = source.UpdatedAt
		b.Title = source.Title
		b.Author = source.Author
		b.Price = source.Price
	}
	return false
}
