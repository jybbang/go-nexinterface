package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jybbang/go-core-architecture/core"
	v1 "github.com/jybbang/nexinterface/src/controllers/v1"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"go.uber.org/zap"
)

var logger *zap.Logger
var log *zap.SugaredLogger

func init() {
	logger, _ = zap.NewProduction()
	log = logger.Sugar()

	applicationSetup()
	infrastructureSetup()

	core.UseTracing(core.TracingSettings{
		ServiceName: "books",
		Endpoint:    "http://localhost:9411/api/v2/spans",
	})
}

func main() {
	log.Info("program start")
	defer log.Info("program end")

	r := gin.Default()

	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)

	apiV1 := r.Group("/api/v1")
	{
		v1.AddBookController(apiV1)
	}

	r.Run(":9000")
	core.Close()
}
