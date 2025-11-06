package middleware

import (
	"time"
	"github.com/gin-gonic/gin"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		c.Writer.Header().Set("X-Response-Time", time.Since(start).String())
	}
}
