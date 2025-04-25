package models

type AnalyzeResponse struct {
	HtmlVersion string
	PageTitle   string
	Headings    map[string]int
}
