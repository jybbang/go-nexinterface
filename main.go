package main

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/jybbang/nexinterface/src/controllers/v1"
	"go.uber.org/zap"
)

var logger *zap.Logger
var log *zap.SugaredLogger

func init() {
	logger, _ = zap.NewProduction()
	log = logger.Sugar()

	applicationSetup()
	infrastructureSetup()
}

func main() {
	log.Info("program start")
	defer log.Info("program end")

	router := gin.Default()

	apiV1 := router.Group("/api/v1")
	{
		v1.AddBookController(apiV1)
	}

	router.Run(":9000")
}
