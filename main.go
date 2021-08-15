package main

import (
	"context"

	"github.com/jybbang/go-core-architecture/core"
	"github.com/jybbang/nexinterface/src/books/commands"
	"go.uber.org/zap"
)

var logger *zap.Logger
var log *zap.SugaredLogger

func init() {
	logger, _ = zap.NewDevelopment()
	log = logger.Sugar()

	core.AddMetrics(core.MetricsSettings{Endpoint: "/metrics"})
	core.AddTracing(core.TracingSettings{ServiceName: "demo"})

	applicationSetup()
	infrastructureSetup()
}

func main() {
	log.Info("program start")
	defer log.Info("program end")

	// demo
	cmd := &commands.CreateBookCommand{
		Title:  "test title",
		Author: "JYB",
		Price:  10000,
	}

	ctx := context.Background()
	book := core.GetMediator().Send(ctx, cmd)
	log.Info(book.V, book.E)

	// http.ListenAndServe(":9000", nil)
}
