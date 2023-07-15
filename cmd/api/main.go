package main

import (
	"go.uber.org/zap"

	"greenbone-task/internal/api"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	api.Run(logger, "")
}
