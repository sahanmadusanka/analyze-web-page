package analyze

import (
	"net/url"
	"strings"
	"sync"
	client "web-page-analyzer/service/http_client_service"

	log "github.com/sirupsen/logrus"
)

type LinkResponse struct {
	Url        string
	IsExternal bool
	Error      error
}

// Analyze whether the given links are accessible and identify which ones are internal or external relative to the 'pageUrl'
func AnalyzeLinks(pageUrl string, links []string) map[string]LinkResponse {
	var wg sync.WaitGroup

	ch := make(chan LinkResponse, len(links))

	hostname, err := getHostname(pageUrl)
	if err != nil {
		log.Errorf("cannot get hostname from page URL %s: %v", pageUrl, err)
		return map[string]LinkResponse{}
	}

	for _, link := range links {
		wg.Add(1)
		go checkLink(hostname, link, &wg, ch)
	}

	wg.Wait()
	close(ch)

	results := make(map[string]LinkResponse)

	for result := range ch {
		results[result.Url] = result
	}

	return results
}

func getHostname(pageUrl string) (string, error) {
	u, err := url.Parse(pageUrl)
	if err != nil {
		log.Errorf("cannot parse page url %s: %v", pageUrl, err)
		return "", err
	}
	return u.Scheme + "://" + u.Hostname(), nil
}

func checkLink(hostname string, link string, wg *sync.WaitGroup, ch chan<- LinkResponse) {
	defer wg.Done()

	resp, err := client.GetUrl(link, 2)
	if err != nil {
		ch <- LinkResponse{Url: link, Error: err}
		return
	}

	defer resp.Body.Close()

	isExternalLink := !strings.HasPrefix(link, hostname)

	ch <- LinkResponse{Url: link, IsExternal: isExternalLink}
}
