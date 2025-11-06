package middleware

import "github.com/gin-gonic/gin"

// Placeholder for role-based checks later
func RequireRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) { c.Next() }
}
