package v1

import (
	"net/http"

	m "web-page-analyzer/models"
	"web-page-analyzer/service/analyze"
	client "web-page-analyzer/service/http_client_service"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Request struct {
	Url string `json:"url" binding:"required" required:"$field is required"`
}

// @Summary Analyze given Web page URL
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/analyze [post]
func AnalyzeUrl(c *gin.Context) {
	var request m.Request

	if err := c.BindJSON(&request); err != nil {
		log.Error("Error in decoding request payload; Error: ", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	webClient := client.WebHttpClient{}

	resp, err := analyze.Analyze(webClient, &request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
