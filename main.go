package main

import (
	"context"

	"github.com/jybbang/go-core-architecture/core"
	"github.com/jybbang/nexinterface/src/books/commands"
	"go.uber.org/zap"
)

var Log *zap.SugaredLogger

func init() {
	logger, _ := zap.NewProduction()
	Log = logger.Sugar()

	core.SetLogger(logger)
}

func main() {
	Log.Info("program start")
	defer Log.Info("program end")

	applicationSetup()
	infrastructureSetup()

	// demo
	cmd := &commands.CreateBookCommand{
		Title:  "test title",
		Author: "JYB",
		Price:  123,
	}

	ctx := context.Background()
	book := core.GetMediator().Send(ctx, cmd)

	Log.Info(book.V)
}
