package analyze

import (
	"net/http"
	"testing"
	"web-page-analyzer/models"

	client "web-page-analyzer/service/http_client_service"

	"github.com/stretchr/testify/assert"
)

const VALID_HTML string = `
<!DOCTYPE html>
<html>
<title>Sample title</title>
<body>

<h2>An Unordered HTML List</h2>

<ul>
  <li>Coffee</li>
  <li>Tea</li>
  <li>Milk</li>
  <li><a href="https://abc.com/ert">link 3</a></li>
</ul>  

<h3>An Ordered HTML List</h3>

<ol>
  <li>Coffee</li>
  <li>Tea</li>
  <li>Milk</li>
  <li><a href="https://abc.com/xyz">link 1</a></li>
  <li><a href="https://qwe.com/xyz">link 2</a></li>
  <li><a href="https://abc.com/xyz">link 3</a></li>
</ol> 



<form>
	<input type="email" />
	<input type="password" />
	<input type="submit" >Login</input>
</form>

</body>
</html>
`

func TestAnalyze(t *testing.T) {
	mockClient := MockWebHttpClient{}

	request := &models.Request{
		Url: "https://google.com",
	}

	mockClient.On("GetContent", "https://google.com", 3).Return(&http.Response{StatusCode: 200, Body: http.NoBody}, nil)

	_, err := Analyze(&mockClient, request)

	if err != nil {
		t.Fatalf("error in analyze request : %v", err)
	}

	assert.Nil(t, err)
}

func TestAnalyze_invalidUrl(t *testing.T) {
	mockClient := MockWebHttpClient{}

	request := &models.Request{
		Url: "abc://google.com",
	}

	mockClient.On("GetContent", "https://google.com", 3).Return(&http.Response{StatusCode: 404, Body: http.NoBody}, nil)

	_, err := Analyze(&mockClient, request)
	assert.NotNil(t, err)
}

func TestPassDocument(t *testing.T) {

	page := &client.WebPageContent{
		Content: VALID_HTML,
	}

	doc := ParseDocument(page)
	assert.NotNil(t, doc)
	assert.NotEmpty(t, doc.Nodes)
}

func TestGetHtmlVersion(t *testing.T) {
	page := &client.WebPageContent{
		Content: VALID_HTML,
	}

	doc := ParseDocument(page)

	htmlVersion := GetHTMLVersion(doc)
	assert.Equal(t, "HTML5", htmlVersion)
}

func TestGetPageTitle(t *testing.T) {
	page := &client.WebPageContent{
		Content: VALID_HTML,
	}

	doc := ParseDocument(page)

	title := GetPageTitle(doc)

	assert.Equal(t, "Sample title", title)
}

func TestGetHeadings(t *testing.T) {
	page := &client.WebPageContent{
		Content: VALID_HTML,
	}

	doc := ParseDocument(page)

	headings := GetHeadings(doc)

	assert.NotEmpty(t, headings)
	assert.Equal(t, 6, len(headings))
	assert.Equal(t, 1, headings["h2"])
	assert.Equal(t, 1, headings["h3"])
}

func TestGetLinks(t *testing.T) {
	mockClient := MockWebHttpClient{}

	page := &client.WebPageContent{
		Content: VALID_HTML,
	}

	mockClient.On("GetContent", "https://qwe.com/xyz", 2).Return(&http.Response{StatusCode: 200, Body: http.NoBody}, nil)
	mockClient.On("GetContent", "https://abc.com", 2).Return(&http.Response{StatusCode: 200, Body: http.NoBody}, nil)
	mockClient.On("GetContent", "https://abc.com/ert", 2).Return(&http.Response{StatusCode: 200, Body: http.NoBody}, nil)
	mockClient.On("GetContent", "https://abc.com/xyz", 2).Return(&http.Response{StatusCode: 200, Body: http.NoBody}, nil)

	doc := ParseDocument(page)

	resp := GetLinks(&mockClient, "https://abc.com", doc)

	assert.NotNil(t, resp)
	assert.Equal(t, 3, len(resp))
	assert.False(t, resp["https://abc.com/xyz"].IsExternal)
	assert.True(t, resp["https://qwe.com/xyz"].IsExternal)
}

func TestLoginPageExsit(t *testing.T) {
	page := &client.WebPageContent{
		Content: VALID_HTML,
	}

	doc := ParseDocument(page)

	loginExsit := LoginPageExist(doc)

	assert.True(t, loginExsit)
}
