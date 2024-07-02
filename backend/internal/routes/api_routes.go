package routes

import (
	"backend/internal/config"
	"backend/internal/handlers"
	"backend/internal/middleware"
	"backend/internal/repositories"
	"backend/pkg/database"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	gorrilla "github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
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

	store := cookie.NewStore([]byte(config.AppConfig.Secret))
	router.Use(sessions.Sessions("session", store))

	githubProvider := github.New(config.AppConfig.GithubClientID, config.AppConfig.GithubClientSecret, config.AppConfig.CallbackUrl)
	goth.UseProviders(githubProvider)
	gothic.Store = gorrilla.NewCookieStore([]byte(config.AppConfig.GithubClientID))

	// Initialize repositories and handlers
	customerRepo := repositories.NewCustomerRepository(db, logger)
	customerHandler := handlers.NewCustomerHandler(customerRepo, logger)

	orderRepo := repositories.NewOrderRepository(db, logger)
	orderHandler := handlers.NewOrderHandler(orderRepo, logger)

	authHandler := handlers.NewAuthenticationHandler(logger)

	// Setup routes
	v1 := router.Group("/api/v1")
	{
		customers := v1.Group("/customers")
		customers.Use(middleware.AuthMiddleware())
		{
			customers.GET("", customerHandler.GetAllCustomers)
			customers.POST("", customerHandler.CreateCustomer)
			customers.PUT("/:id", customerHandler.UpdateCustomer)
		}
		orders := v1.Group("/orders")
		orders.Use(middleware.AuthMiddleware())
		{
			orders.POST("", orderHandler.CreateOrder)
			orders.PUT("/:id", orderHandler.UpdateOrder)
			orders.DELETE("/:id", orderHandler.DeleteOrder)
			orders.GET("/:id", orderHandler.GetOrderByID)
			orders.GET("", orderHandler.GetAllOrders)
		}
		users := v1.Group("/users")
		users.Use(middleware.AuthMiddleware())
		{
			users.GET("/:user_id/orders", orderHandler.GetOrdersByUserID)
		}

		authentication := v1.Group("/auth")
		{
			authentication.GET("/", authHandler.Home)
			authentication.GET("/callback", authHandler.CallBack)
			authentication.GET("/login", authHandler.SignIn)

		}
	}
}
