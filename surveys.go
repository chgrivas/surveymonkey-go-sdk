package surveymonkey

import (
	"context"
	"fmt"
	"net/http"
)

// Survey is a survey type
type Survey struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Nickname string `json:"nickname"`
	Href     string `json:"href"`
}

// GetSurveysOptions is a set of optional properties that can be used for surveys pagination.
type GetSurveysOptions struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
}

// GetSurveys returns a list of the existing surveys.
func (c *Client) GetSurveys(ctx context.Context, options *GetSurveysOptions) ([]Survey, error) {
	page := 1
	perPage := 50
	if options != nil {
		page = options.Page
		perPage = options.PerPage
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/surveys?page=%d&per_page=%d", c.BaseURL, page, perPage), nil)
	if err != nil {
		return nil, err
	}

	var res []Survey
	if err := c.sendRequest(ctx, req, &res); err != nil {
		return nil, err
	}

	return res, err
}
