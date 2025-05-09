package routers

import (
	"net/http"
	"time"
	v1 "web-page-analyzer/routers/api/v1"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

func initLogging() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		latency := time.Since(startTime)
		status := c.Writer.Status()

		log.WithFields(log.Fields{
			"status":   status,
			"method":   c.Request.Method,
			"path":     c.Request.URL.Path,
			"latency":  latency,
			"clientIP": c.ClientIP(),
		}).Info("Incoming request")
	}
}

func InitRouter() *gin.Engine {
	log.Info("Initilizing GIN router")

	r := gin.New()

	r.Use(initLogging(), gin.Recovery())

	r.LoadHTMLFiles("html/index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.POST("/api/v1/analyze", v1.AnalyzeUrl)

	err := r.Run(":3000")
	if err != nil {
		log.Fatalf("[Error] failed to start Gin server due to: %v", err)
	}

	return r
}
