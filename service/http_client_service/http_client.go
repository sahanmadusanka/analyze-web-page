package http

import (
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type WebPageContent struct {
	Content    string
	StatusCode int
	Errors     map[string]string
}

func GetUrl(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Errorf("error in fetching a web page, url: %v error: %v", url, err)
		return nil, err
	}
	return resp, nil
}

func FetchWebPage(url string) *WebPageContent {
	resp, err := GetUrl(url)
	if err != nil {
		return nil
	}

	return buildResponse(resp)
}

func buildResponse(resp *http.Response) *WebPageContent {
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("error in read web page body, error: %v", err)
		return nil
	}

	return &WebPageContent{
		Content:    string(bytes),
		StatusCode: resp.StatusCode,
	}
}
