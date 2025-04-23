package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

func InitRouter() *gin.Engine {
	log.Info("initilizing GIN router")

	r := gin.Default()

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	err := r.Run(":8080")
	if err != nil {
		panic("[Error] failed to start Gin server due to: " + err.Error())
	}

	return r
}
