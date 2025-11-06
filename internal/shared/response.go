package shared

import "github.com/gin-gonic/gin"

func JSONErr(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{"error": msg})
}
