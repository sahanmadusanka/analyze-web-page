package analyze

import (
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
</ul>  

<h3>An Ordered HTML List</h3>

<ol>
  <li>Coffee</li>
  <li>Tea</li>
  <li>Milk</li>
</ol> 

</body>
</html>
`

func TestAnalyze(t *testing.T) {
	request := &models.Request{
		Url: "https://google.com",
	}

	_, err := Analyze(request)

	if err != nil {
		t.Fatalf("error in analyze request : %v", err)
	}

	assert.Nil(t, err)
}

func TestAnalyze_invalidUrl(t *testing.T) {
	request := &models.Request{
		Url: "abc://google.com",
	}

	_, err := Analyze(request)
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

	htmlVersion := GetHtmlVersion(doc)
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
