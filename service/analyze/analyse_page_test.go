package analyze

import (
	"testing"
	"web-page-analyzer/models"

	"github.com/stretchr/testify/assert"
)

func TestAnalyze(t *testing.T) {
	request := models.Request{
		Url: "https://google.com",
	}

	_, err := Analyze(request)

	if err != nil {
		t.Fatalf("error in analyze request : %v", err)
	}

	assert.Nil(t, err)
}
