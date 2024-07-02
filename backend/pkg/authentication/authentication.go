package authentication

import (
	"errors"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/github"
	"time"
)

// Authentication struct to store session details
type Authentication struct {
	AuthURL      string
	AccessToken  string
	RefreshToken string
	ExpiresAt    time.Time
	UserID       string
}

// ValidateToken validates the access token using the Goth provider.
func ValidateToken(accessToken string) (bool, error) {
	// Fetch the provider
	provider, err := goth.GetProvider("github")
	if err != nil {
		return false, err
	}

	// Create a session with the access token
	session := &github.Session{
		AuthURL:     "",
		AccessToken: accessToken,
	}

	// Use the provider to get the user and validate the token
	user, err := provider.FetchUser(session)
	if err != nil {
		return false, err
	}

	// Check if the user is valid
	if user.AccessToken == "" {
		return false, errors.New("invalid access token")
	}

	// Additional validation logic can be added here if necessary
	return true, nil
}
