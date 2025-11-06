package health

import "github.com/gin-gonic/gin"

func Register(r *gin.RouterGroup) {
	r.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"ok": true}) })
}
