package analyze

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnalyzeLinks(t *testing.T) {
	pageUrl := "https://abc.com/"

	links := []string{"https://abc.com/xyz", "https://qwe.com/xyz"}

	resp := AnalyzeLinks(pageUrl, links)
	assert.NotNil(t, resp)
	assert.Equal(t, 2, len(resp))
	assert.False(t, resp["https://abc.com/xyz"].IsExternal)
	assert.True(t, resp["https://qwe.com/xyz"].IsExternal)
}
