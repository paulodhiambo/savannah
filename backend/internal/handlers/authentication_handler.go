package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	"github.com/sirupsen/logrus"
	"net/http"
)

type AuthenticationHandler struct {
	logger *logrus.Logger
}

func NewAuthenticationHandler(logger *logrus.Logger) *AuthenticationHandler {
	return &AuthenticationHandler{logger: logger}
}

// Home godoc
// @Summary Display home page
// @Description Display home page with user token
// @Produce html
// @Success 200 {string} string "OK"
// @Router / [get]
func (h *AuthenticationHandler) Home(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	data := struct {
		Token interface{}
	}{
		Token: user,
	}
	// Render the HTML template with the variable
	c.HTML(http.StatusOK, "home.html", data)
}

// SignIn godoc
// @Summary Initiate sign-in process
// @Description Redirects the user to the Logto sign-in page
// @Produce json
// @Success 302 {string} string "Redirect"
// @Router /api/v1/auth/login [get]
func (h *AuthenticationHandler) SignIn(c *gin.Context) {
	q := c.Request.URL.Query()
	q.Add("provider", "github")
	c.Request.URL.RawQuery = q.Encode()
	gothic.BeginAuthHandler(c.Writer, c.Request)
}

// CallBack godoc
// @Summary Handle sign-in callback
// @Description Handle callback from Logto after user signs in
// @Produce json
// @Success 302 {string} string "Redirect"
// @Router /api/v1/auth/callback [get]
func (h *AuthenticationHandler) CallBack(c *gin.Context) {
	htmlFormat := `<html><body>%v</body></html>`
	q := c.Request.URL.Query()
	q.Add("provider", "github")
	c.Request.URL.RawQuery = q.Encode()
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		err = c.AbortWithError(http.StatusInternalServerError, err)
		h.logger.Error(err)
		return
	}
	res, err := json.Marshal(user)
	if err != nil {
		err = c.AbortWithError(http.StatusInternalServerError, err)
		h.logger.Error(err)
		return
	}
	session := sessions.Default(c)
	session.Set("user", user)
	err = session.Save()

	if err != nil {
		err = c.AbortWithError(http.StatusInternalServerError, err)
		h.logger.Error(err)
		return
	}
	jsonString := string(res)
	html := fmt.Sprintf(htmlFormat, jsonString)
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))
}
