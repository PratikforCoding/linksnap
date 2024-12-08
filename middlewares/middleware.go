package middlewares

import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

// RateLimitMiddleware is a middleware that limits the number of requests a user can make in a given period
func RateLimitMiddleware() gin.HandlerFunc {
	rate := limiter.Rate {
		Period: 1 * time.Minute,
		Limit: 10,
	}
	store := memory.NewStore()
	instance := limiter.New(store, rate)

	// Return the middleware
	return func(c *gin.Context) {
		ip := c.ClientIP()

		context, err := instance.Get(c, ip) 
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			c.Abort()
			return
		}
		if context.Reached {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
			c.Abort()
			return
		}
		c.Next()
	}
}