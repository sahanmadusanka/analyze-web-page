# Web Page Analyzer

## Description

This is a simple web application that analyzes a web page based on a provided URL.

The application allows users to enter a URL and retrieve key information about the page, such as HTML version, headings, links, and login form exsist.

## Requirements
* Go lang 1.24
* Docker


## How to build the application

Run `docker build . -t web-page-analyzer` from the root of the project.

## How to run the application

Run `docker run -d -p 3000:3000 web-page-analyzer` to start the application on port `3000`

Navigate to the [localhost:3000](localhost:3000) in the browser to access the analyze page.


## Available APIs

### Analyze web page
`POST /api/v1/analyze`

##### Request body
```
{
    "url":<url need to analyze>
}
```
##### Response schema
```
{
  "htmlVersion": "string",
  "pageTitle": "string",
  "headings": {
    "h1": "integer",
    "h2": "integer",
    "h3": "integer",
    "h4": "integer",
    "h5": "integer",
    "h6": "integer"
  },
  "link": {
    "InternalLinkCount": "integer",
    "ExternalLinkCount": "integer",
    "InaccessibleLinks": ["string"]
  }
}
```

### Prometheus metrics
`GET /metrics`


## Assumptions/Decisions
* Analyzing web page URL is open and can access without authentication.
* Fetch only the links available in `<a>` tags on the page.
* Internal link detection works by checking if the link has the same hostname as the given URL to analyze the page.
* Login form detection only checks if the email and password fields are available under a form.

## Possible Improvments
* The Gin router can use groups for URLs with the same prefix.
* Externalize the property to a URL validation regex.
* Improvements for the analyzer UI to make it more intuitive and attractive (UI/UX).

