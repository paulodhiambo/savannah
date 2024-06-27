package server

import (
	_ "backend/docs"
	"backend/internal/routes"
	"backend/pkg/authentication"
	"backend/pkg/logging"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	_ "github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"github.com/logto-io/go/client"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
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
	router.LoadHTMLGlob("templates/*.html")

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	logtoConfig := &client.LogtoConfig{
		Endpoint:  "https://auth.streempoint.com/",
		AppId:     "nsg5qgl54ysk1sysa7mcq",
		AppSecret: "cpmnqXe122G0jlMk6tT5Kj06EMvdCZUd",
	}

	// Add a link to perform a sign-in request on the home page
	router.GET("/", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		logtoClient := client.NewLogtoClient(
			logtoConfig,
			&authentication.SessionStorage{Session: session},
		)

		_, err := logtoClient.GetAccessToken("https://savannah-api-dot-streempoint.ue.r.appspot.com/api")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(logtoClient.GetIdToken())
		//authState := "You are not logged in to this website. :("
		//
		//if logtoClient.IsAuthenticated() {
		//	authState = "You are logged in to this website! :)"
		//}

		// ...
		//homePage := `<h1>Hello Logto</h1>` +
		//	"<div>" + authState + "</div>" +
		//	"<div>" + logtoClient.GetIdToken() + "</div>"

		//ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(homePage))

		data := struct {
			Token string
		}{
			Token: logtoClient.GetIdToken(),
		}

		// Render the HTML template with the variable
		ctx.HTML(http.StatusOK, "home.html", data)
	})

	// Add a route for handling sign-in requests
	router.GET("/sign-in", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		logtoClient := client.NewLogtoClient(
			logtoConfig,
			&authentication.SessionStorage{Session: session},
		)

		// The sign-in request is handled by Logto.
		// The user will be redirected to the Redirect URI on signed in.
		signInUri, err := logtoClient.SignIn("https://6c84-105-163-1-253.ngrok-free.app/callback")
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		// Redirect the user to the Logto sign-in page.
		ctx.Redirect(http.StatusTemporaryRedirect, signInUri)
	})

	router.GET("/callback", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		logtoClient := client.NewLogtoClient(
			logtoConfig,
			&authentication.SessionStorage{Session: session},
		)
		// The sign-in callback request is handled by Logto
		err := logtoClient.HandleSignInCallback(ctx.Request)
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		// Jump to the page specified by the developer.
		// This example takes the user back to the home page.
		ctx.Redirect(http.StatusTemporaryRedirect, "/")
	})

	// Setup routes
	routes.SetupRoutes(router, logger)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	logger.Info("Starting server on :8080")

	return router.Run(":8080")
}
