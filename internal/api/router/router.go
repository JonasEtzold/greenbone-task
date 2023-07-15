package router

import (
	"time"

	middleware "github.com/deepmap/oapi-codegen/pkg/gin-middleware"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"greenbone-task/internal/api/models"
	"greenbone-task/internal/api/services"
)

func Setup(logger *zap.Logger) *gin.Engine {
	swagger, err := models.GetSwagger()
	if err != nil {
		logger.Panic("Failed to load openAPI service specification.", zap.Error(err))
	}

	notifier := services.NewNotifierService(logger)

	app := gin.New()
	// Middlewares
	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	app.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	// Logs all panic to error log
	//   - stack means whether output the stack info.
	app.Use(ginzap.RecoveryWithZap(logger, true))
	app.Use(middleware.OapiRequestValidator(swagger))
	app.NoRoute(services.NoRouteHandler())

	// Routes
	services.RegisterHandlers(app, services.NewComputerService(logger, notifier))

	return app
}
