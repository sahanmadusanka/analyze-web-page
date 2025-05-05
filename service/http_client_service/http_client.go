package http

import (
	"io"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type WebClient interface {
	GetContent(url string, timeout int) (*http.Response, error)
}

type WebHttpClient struct{}

type WebPageContent struct {
	Content    string
	StatusCode int
	Errors     map[string]string
}

func (WebHttpClient) GetContent(url string, timeout int) (*http.Response, error) {
	client := http.Client{
		Timeout: time.Second * time.Duration(timeout),
	}

	resp, err := client.Get(url)
	if err != nil {
		log.Errorf("error fetching web page; url: %v, error: %v", url, err)
		return nil, err
	}
	return resp, nil
}

func FetchWebPage(w WebClient, url string) (*WebPageContent, error) {
	resp, err := w.GetContent(url, 3)
	if err != nil {
		return nil, nil
	}

	return buildResponse(resp), nil
}

func buildResponse(resp *http.Response) *WebPageContent {
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("error reading web page body: %v", err)
		return nil
	}

	return &WebPageContent{
		Content:    string(bytes),
		StatusCode: resp.StatusCode,
	}
}
