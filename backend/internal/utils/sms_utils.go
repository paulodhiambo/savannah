package utils

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// SendSMS sends an SMS message using Africa’s Talking API.
func SendSMS(apiKey, username, to, message string) error {
	apiUrl := "https://api.sandbox.africastalking.com/version1/messaging"

	data := url.Values{}
	data.Set("username", username)
	data.Set("to", to)
	data.Set("message", message)

	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(data.Encode()))
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("apikey", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("failed to close response body")
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send SMS: status code %d", resp.StatusCode)
	}

	return nil
}
