package users

import (
	"github.com/gin-gonic/gin"
	"thistoric-backend/internal/security/jwt"
)

func Register(r *gin.RouterGroup) {
	r.GET("/me", jwt.RequireUser(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"user_id": c.GetString("sub"),
			"email":   c.GetString("email"),
		})
	})
}
