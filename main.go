package main

import (
	"github.com/jybbang/go-core-architecture/core"
	"github.com/jybbang/nexinterface/src/books/commands"
	"github.com/jybbang/nexinterface/src/entities"
	"go.uber.org/zap"
)

var Log *zap.SugaredLogger

func init() {
	logger, _ := zap.NewProduction()
	Log = logger.Sugar()
}

func main() {
	Log.Info("program start")
	defer Log.Info("program end")

	applicationSetup()
	infrastructureSetup()

	Log.Debug("debuging")
	cmd := &commands.CreateBookCommand{
		Title:  "test title",
		Author: "JYB",
		Price:  123,
	}

	book, _ := core.GetMediator().Send(cmd)

	Log.Info(book.(*entities.Book))
}
