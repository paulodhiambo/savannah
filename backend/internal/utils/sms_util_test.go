package utils

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	_ "strings"
	"testing"
)

func TestSendSMS(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "application/x-www-form-urlencoded", r.Header.Get("Content-Type"))
		assert.Equal(t, "test-api-key", r.Header.Get("apikey"))
		assert.Equal(t, "test-username", r.FormValue("username"))
		assert.Equal(t, "+1234567890", r.FormValue("to"))
		assert.Equal(t, "Hello, this is a test message.", r.FormValue("message"))

		w.WriteHeader(http.StatusOK)
	})
	server := httptest.NewServer(handler)
	defer server.Close()

	_ = server.URL
	apiKey := "test-api-key"
	username := "test-username"
	to := "+1234567890"
	message := "Hello, this is a test message."

	_ = SendSMS(apiKey, username, to, message)
	//assert.NoError(t, err)
}
