package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func RateLimitMiddleware(redisClient *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.ClientIP()
		value, err := redisClient.Get(c.Request.Context(), key).Result()
		if err == redis.Nil {
			// Key doesn't exist, set it with expiration
			err = redisClient.Set(c.Request.Context(), key, 1, time.Minute).Err()
			if err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
		} else if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		} else {
			// Key exists, increment it
			count, err := redisClient.Incr(c.Request.Context(), key).Result()
			if err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			if count > 60 {
				c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
				return
			}
		}
		c.Next()
	}
}

func SecurityHeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Referrer-Policy", "strict-origin-when-cross-origin")
		c.Header("Content-Security-Policy", "default-src 'self'; script-src 'self' 'unsafe-inline' https://cdn.tailwindcss.com; style-src 'self' 'unsafe-inline';")
		c.Next()
	}
}
