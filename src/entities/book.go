package entities

import (
	"github.com/jybbang/go-core-architecture/core"
)

type Book struct {
	core.Entity
	Title  string `bson:"title,omitempty"`
	Author string `bson:"author,omitempty"`
	Price  int    `bson:"price,omitempty"`
}
