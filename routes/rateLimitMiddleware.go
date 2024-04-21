package routes

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type VisitorData struct {
	visitors map[string]*rate.Limiter
	mtx      sync.Mutex
}

func (v *VisitorData) getVisitor(ip string) *rate.Limiter {
	v.mtx.Lock()
	defer v.mtx.Unlock()

	limiter, exists := v.visitors[ip]
	if !exists {
		limiter = rate.NewLimiter(2, 2) // 2 requests per second
		v.visitors[ip] = limiter
	}

	return limiter
}

func RateLimitMiddleware(v *VisitorData) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		limiter := v.getVisitor(ip)

		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"status": 429,
				"error":  "Too many requests",
			})
			return
		}

		c.Next()
	}
}
