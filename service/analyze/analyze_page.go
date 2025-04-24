package analyze

import (
	"errors"
	"regexp"
	m "web-page-analyzer/models"

	log "github.com/sirupsen/logrus"
)

// Make it external config
const URL_VALIDATION_REGEX = "((http|https)://)(www.)?[a-zA-Z0-9@:%._\\+~#?&//=]{2,256}\\.[a-z]{2,6}\\b([-a-zA-Z0-9@:%._\\+~#?&//=]*)"

func Analyze(request m.Request) (*m.AnalyzeResponse, error) {
	//Validate input
	if err := validateRequest(&request); err != nil {
		return nil, err
	}

	return nil, nil
}

func validateRequest(request *m.Request) error {
	urlRegex := regexp.MustCompile(URL_VALIDATION_REGEX)

	if !urlRegex.MatchString(request.Url) {
		log.Warn("Input URL is not valid; URL: ", request.Url)
		return errors.New("input url is not valid")
	}
	return nil
}
