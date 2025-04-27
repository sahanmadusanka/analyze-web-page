package http

import (
	"io"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type WebPageContent struct {
	Content    string
	StatusCode int
	Errors     map[string]string
}

func GetUrl(url string, timeout int) (*http.Response, error) {
	client := http.Client{
		Timeout: time.Second * time.Duration(timeout),
	}

	resp, err := client.Get(url)
	if err != nil {
		log.Errorf("error in fetching a web page, url: %v error: %v", url, err)
		return nil, err
	}
	return resp, nil
}

func FetchWebPage(url string) *WebPageContent {
	resp, err := GetUrl(url, 3)
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
