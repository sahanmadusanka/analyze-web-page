package v1

import (
	"net/http"

	m "web-page-analyzer/models"
	"web-page-analyzer/service/analyze"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// @Summary Analyze given Web page URL
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/analyze [post]
func AnalyzeUrl(c *gin.Context) {
	var request m.Request

	if err := c.BindJSON(&request); err != nil {
		log.Error("Error in decording request payload; Error: ", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if _, err := analyze.Analyze(&request); err != nil {

	}
}
