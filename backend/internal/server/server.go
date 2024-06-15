package server

import (
	"backend/internal/routes"
	"backend/pkg/logging"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Savannah Informatics Interview
// @version 1.0
// @description Interview application for Savannah Informatics Backend Role.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email paulodhiambo962@gmail.com

// Run @host localhost:8080
// @BasePath /v1
func Run() error {
	logger := logging.GetLogger()
	router := gin.Default()

	// Setup routes
	routes.SetupRoutes(router, logger)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	logger.Info("Starting server on :8080")

	return router.Run(":8080")
}
