package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type clientLimiter struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var (
	clients = make(map[string]*clientLimiter)
	mu      sync.Mutex
)

func getLimiter(ip string, r rate.Limit, b int) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	client, exists := clients[ip]

	if !exists {
		limiter := rate.NewLimiter(r, b)

		clients[ip] = &clientLimiter{
			limiter:  limiter,
			lastSeen: time.Now(),
		}

		return limiter
	}

	client.lastSeen = time.Now()
	return client.limiter
}

func RateLimit(r rate.Limit, b int) gin.HandlerFunc {

	return func(c *gin.Context) {

		ip := c.ClientIP()

		limiter := getLimiter(ip, r, b)

		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "too many requests",
			})
			return
		}

		c.Next()
	}
}
