package api

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"greenbone-task/internal/api/router"
	"greenbone-task/internal/config"
	"greenbone-task/internal/persistence/db"
)

func Run(logger *zap.Logger, configPath string) {
	config.Setup(configPath, logger)
	db.Setup(logger)
	gin.SetMode(viper.GetString("SERVER_MODE"))

	web := router.Setup(logger)
	logger.Info("greenbone task running on port " + viper.GetString("SERVER_PORT"))
	logger.Info("==================>")
	_ = web.Run(":" + viper.GetString("SERVER_PORT"))
}
