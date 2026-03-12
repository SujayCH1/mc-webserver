package router

import (
	"mc-webserver/internal/api"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/health", api.HealthCheck)

	return r
}
