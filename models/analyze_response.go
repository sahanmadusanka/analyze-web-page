package models

type AnalyzeResponse struct {
	HtmlVersion    string         `json:"htmlVersion"`
	PageTitle      string         `json:"pageTitle"`
	Headings       map[string]int `json:"headings"`
	Link           Link           `json:"link"`
	LoginPageExsit bool           `json:"loginPageExsit"`
}

type Link struct {
	InternalLinkCount int
	ExternalLinkCount int
	InaccessibleLinks []string
}
