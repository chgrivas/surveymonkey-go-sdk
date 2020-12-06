package surveymonkey

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// BaseURLV3 is the base default base URL of the client.
const BaseURLV3 = "https://api.surveymonkey.com/v3"

// Client is an http client type
type Client struct {
	BaseURL     string
	accessToken string
	HTTPClient  *http.Client
}

// NewClient returns a client for reaching the SurveyMonkey API.
func NewClient(accessToken string) *Client {
	return &Client{
		BaseURL:     BaseURLV3,
		accessToken: accessToken,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

type successResponse struct {
	Data    interface{} `json:"data"`
	PerPage int         `json:"per_page"`
	Page    int         `json:"page"`
	Total   int         `json:"total"`
}

type errorResponse struct {
	Error struct {
		ID             interface{} `json:"id"`
		Name           string      `json:"name"`
		Docs           string      `json:"docs"`
		Message        string      `json:"message"`
		HTTPStatusCode int         `json:"http_status_code"`
	} `json:"error"`
}

func (c *Client) sendRequest(ctx context.Context, req *http.Request, v interface{}) error {
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("bearer %s", c.accessToken))

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes errorResponse
		if err := json.NewDecoder(res.Body).Decode(&errRes); err != nil {
			return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
		}

		return errors.New(errRes.Error.Message)
	}

	fullResponse := successResponse{
		Data: v,
	}
	if err := json.NewDecoder(res.Body).Decode(&fullResponse); err != nil {
		return err
	}

	return nil
}
