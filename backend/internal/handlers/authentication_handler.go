package handlers

import (
	"backend/pkg/authentication"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/logto-io/go/client"
	"github.com/sirupsen/logrus"
	"net/http"
)

type AuthenticationHandler struct {
	logger *logrus.Logger
	client *client.LogtoConfig
}

func NewAuthenticationHandler(client *client.LogtoConfig, logger *logrus.Logger) *AuthenticationHandler {
	return &AuthenticationHandler{logger: logger, client: client}
}

func (h *AuthenticationHandler) Home(c *gin.Context) {
	session := sessions.Default(c)
	logtoClient := client.NewLogtoClient(
		h.client,
		&authentication.SessionStorage{Session: session},
	)

	_, err := logtoClient.GetAccessToken("https://savannah-api-dot-streempoint.ue.r.appspot.com/api")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(logtoClient.GetIdToken())
	data := struct {
		Token string
	}{
		Token: logtoClient.GetIdToken(),
	}

	// Render the HTML template with the variable
	c.HTML(http.StatusOK, "home.html", data)
}

func (h *AuthenticationHandler) SignIn(c *gin.Context) {
	session := sessions.Default(c)
	logtoClient := client.NewLogtoClient(
		h.client,
		&authentication.SessionStorage{Session: session},
	)

	// The user will be redirected to the Redirect URI on signed in.
	signInUri, err := logtoClient.SignIn("https://6c84-105-163-1-253.ngrok-free.app/callback")
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	// Redirect the user to the Logto sign-in page.
	c.Redirect(http.StatusTemporaryRedirect, signInUri)
}

func (h *AuthenticationHandler) CallBack(c *gin.Context) {
	session := sessions.Default(c)
	logtoClient := client.NewLogtoClient(
		h.client,
		&authentication.SessionStorage{Session: session},
	)
	// The sign-in callback request is handled by Logto
	err := logtoClient.HandleSignInCallback(c.Request)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	// This example takes the user back to the home page.
	c.Redirect(http.StatusTemporaryRedirect, "/")
}
