package analyze

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockWebHttpClient struct {
	mock.Mock
}

func (m *MockWebHttpClient) GetContent(url string, timeout int) (*http.Response, error) {
	args := m.Called(url, timeout)
	return args.Get(0).(*http.Response), args.Error(1)
}

func TestAnalyzeLinks(t *testing.T) {
	mockClient := MockWebHttpClient{}

	pageUrl := "https://abc.com/"

	links := []string{"https://abc.com/xyz", "https://qwe.com/xyz"}

	mockClient.On("GetContent", "https://abc.com/xyz", 2).Return(&http.Response{StatusCode: 200, Body: http.NoBody}, nil)
	mockClient.On("GetContent", "https://qwe.com/xyz", 2).Return(&http.Response{StatusCode: 200, Body: http.NoBody}, nil)

	resp := AnalyzeLinks(&mockClient, pageUrl, links)
	assert.NotNil(t, resp)
	assert.Equal(t, 2, len(resp))
	assert.False(t, resp["https://abc.com/xyz"].IsExternal)
	assert.True(t, resp["https://qwe.com/xyz"].IsExternal)
}
