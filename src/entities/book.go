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
