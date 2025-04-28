package analyze

import (
	"errors"
	"regexp"
	"slices"
	"strings"
	"web-page-analyzer/constant"
	m "web-page-analyzer/models"
	client "web-page-analyzer/service/http_client_service"

	"golang.org/x/net/html"

	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
)

func Analyze(request *m.Request) (*m.AnalyzeResponse, error) {

	if !validateUrl(request.Url) {
		return nil, errors.New("input url is not valid")
	}

	resp := client.FetchWebPage(request.Url)
	doc := ParseDocument(resp)

	return buildResponse(request, doc), nil
}

func validateUrl(url string) bool {
	return regexp.MustCompile(constant.URL_VALIDATION_REGEX).MatchString(url)
}

func ParseDocument(c *client.WebPageContent) *goquery.Document {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(c.Content))
	if err != nil {
		log.Errorf("error in parsing document, error: %v", err)
		return nil
	}
	return doc
}

func buildResponse(request *m.Request, doc *goquery.Document) *m.AnalyzeResponse {
	htmlVersion := GetHtmlVersion(doc)

	title := GetPageTitle(doc)

	headings := GetHeadings(doc)

	links := GetLinks(request.Url, doc)

	loginPageExsit := LoginPageExsit(doc)

	return &m.AnalyzeResponse{
		HtmlVersion:    htmlVersion,
		PageTitle:      title,
		Headings:       headings,
		Link:           buildLinkResponse(&links),
		LoginPageExsit: loginPageExsit,
	}
}

func GetHtmlVersion(doc *goquery.Document) string {
	version := ""
	root := doc.Nodes[0]
	for child := root.FirstChild; child != nil; child = child.NextSibling {
		if child.Type == html.DoctypeNode {
			version = child.Data
			break
		}
	}
	d := strings.ToLower(version)
	switch {
	case d == "html":
		return "HTML5"
	case strings.Contains(d, "xhtml 1.0"):
		return "XHTML 1.0"
	case strings.Contains(d, "xhtml 1.1"):
		return "XHTML 1.1"
	case strings.Contains(d, "html 4.01"):
		return "HTML 4.01"
	case strings.Contains(d, "html 3.2"):
		return "HTML 3.2"
	default:
		return "Unknown HTML version"
	}
}

func GetPageTitle(doc *goquery.Document) string {
	title := doc.Find("title")
	return title.Text()
}

func GetHeadings(doc *goquery.Document) map[string]int {
	h1Count := getNodeCount(doc.Find("h1"))
	h2Count := getNodeCount(doc.Find("h2"))
	h3Count := getNodeCount(doc.Find("h3"))
	h4Count := getNodeCount(doc.Find("h4"))
	h5Count := getNodeCount(doc.Find("h5"))
	h6Count := getNodeCount(doc.Find("h26"))
	return map[string]int{"h1": h1Count, "h2": h2Count, "h3": h3Count, "h4": h4Count, "h5": h5Count, "h6": h6Count}
}

func getNodeCount(s *goquery.Selection) int {
	return len(s.Nodes)
}

func GetLinks(pageUrl string, doc *goquery.Document) map[string]LinkResponse {
	var links []string = []string{}

	doc.Find("a").Each(func(index int, tag *goquery.Selection) {
		val, exists := tag.Attr("href")

		if exists && validateUrl(val) && !slices.Contains(links, val) {
			links = append(links, val)
		}
	})

	log.Info("start analyzing links, count :", len(links))

	return AnalyzeLinks(pageUrl, links)
}

func buildLinkResponse(links *map[string]LinkResponse) m.Link {
	var internalLinkCount int
	var externalLinkCount int
	var inaccessibleLinks []string = []string{}

	for _, link := range *links {

		if link.Error != nil {
			inaccessibleLinks = append(inaccessibleLinks, link.Url)
			continue
		}

		if link.IsExternal {
			externalLinkCount++
		} else {
			internalLinkCount++
		}

	}

	return m.Link{
		InternalLinkCount: internalLinkCount,
		ExternalLinkCount: externalLinkCount,
		InaccessibleLinks: inaccessibleLinks,
	}
}

// Only look for a page that contains a form with three inputs: email, password, and a submit button
func LoginPageExsit(doc *goquery.Document) bool {
	email := doc.Find("form input[type^=\"email\"]").Length() > 0
	password := doc.Find("form input[type^=\"password\"]").Length() > 0
	submit := doc.Find("form input[type^=\"submit\"]").Length() > 0

	return email && password && submit
}
