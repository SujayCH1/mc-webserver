package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
)

func ServerOnly() gin.HandlerFunc {

	return func(c *gin.Context) {

		serverIP := os.Getenv("MC_SERVER_IP")

		if c.ClientIP() != serverIP {
			c.AbortWithStatusJSON(403, gin.H{
				"error": "forbidden",
			})
			return
		}

		c.Next()
	}
}
