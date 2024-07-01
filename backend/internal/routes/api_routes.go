package routes

import (
	"backend/internal/config"
	"backend/internal/handlers"
	"backend/internal/repositories"
	"backend/pkg/database"
	"github.com/gin-gonic/gin"
	"github.com/logto-io/go/client"
	"github.com/sirupsen/logrus"
)

func SetupRoutes(router *gin.Engine, logger *logrus.Logger) {
	// Connect to PostgreSQL database
	err := database.Connect(logger)
	db := database.DB
	if err != nil {
		panic("Failed to connect to the database!")
	}
	err = config.Load()
	if err != nil {
		return
	}

	//logto config
	logtoConfig := &client.LogtoConfig{
		Endpoint:  config.AppConfig.OpenIDEndpoint,
		AppId:     config.AppConfig.OpenIDClientID,
		AppSecret: config.AppConfig.OpenIDClientSecret,
	}

	// Initialize repositories and handlers
	customerRepo := repositories.NewCustomerRepository(db, logger)
	customerHandler := handlers.NewCustomerHandler(customerRepo, logger)

	orderRepo := repositories.NewOrderRepository(db, logger)
	orderHandler := handlers.NewOrderHandler(orderRepo, logger)

	authHandler := handlers.NewAuthenticationHandler(logtoConfig, logger)

	// Setup routes
	v1 := router.Group("/api/v1")
	{
		customers := v1.Group("/customers")
		{
			customers.GET("", customerHandler.GetAllCustomers)
			customers.POST("", customerHandler.CreateCustomer)
			customers.PUT("/:id", customerHandler.UpdateCustomer)
		}
		orders := v1.Group("/orders")
		{
			orders.POST("", orderHandler.CreateOrder)
			orders.PUT("/:id", orderHandler.UpdateOrder)
			orders.DELETE("/:id", orderHandler.DeleteOrder)
			orders.GET("/:id", orderHandler.GetOrderByID)
			orders.GET("", orderHandler.GetAllOrders)
		}
		users := v1.Group("/users")
		{
			users.GET("/:user_id/orders", orderHandler.GetOrdersByUserID)
		}

		authentication := v1.Group("/")
		{
			authentication.GET("", authHandler.Home)
			authentication.GET("/callback", authHandler.CallBack)
			authentication.GET("/sign-in", authHandler.SignIn)

		}
	}
}
