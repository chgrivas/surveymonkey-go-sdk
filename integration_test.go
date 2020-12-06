// +build integration

package surveymonkey

import (
	"context"
	"os"
	"testing"
)

func TestGetSurveys(t *testing.T) {
	c := NewClient(os.Getenv("SURVEYMONKEY_ACCESS_TOKEN"))

	ctx := context.Background()
	res, err := c.GetSurveys(ctx, nil)

	if err != nil {
		t.Fatal("expected nil error")
	}

	if res == nil {
		t.Fatal("expected non-nil response")
	}

	if len(res) != 1 {
		t.Fatal("expected 1 survey")
	}
}
