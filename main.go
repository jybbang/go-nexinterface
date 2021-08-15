package main

import (
	"context"
	"net/http"

	"github.com/jybbang/go-core-architecture/core"
	"github.com/jybbang/nexinterface/src/books/commands"
	"go.uber.org/zap"
)

var Log *zap.SugaredLogger

func init() {
	logger, _ := zap.NewProduction()
	Log = logger.Sugar()

	core.SetLogger(logger)
	core.AddMetrics(core.MetricsSettings{Endpoint: "/metrics"})
	core.AddTracing(core.TracingSettings{ServiceName: "demo"})
}

func main() {
	Log.Info("program start")
	defer Log.Info("program end")

	applicationSetup()
	infrastructureSetup()

	// demo
	cmd := &commands.CreateBookCommand{
		Title: "test title",
		// Author: "JYB",
		Price: -1,
	}

	ctx := context.Background()
	book := core.GetMediator().Send(ctx, cmd)

	Log.Info(book.V)

	http.ListenAndServe(":9000", nil)
}
