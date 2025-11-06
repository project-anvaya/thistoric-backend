package leads

import (
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(r *gin.RouterGroup, db *gorm.DB) {
	svc := NewService(NewRepo(db))
	r.POST("/leads", func(c *gin.Context) {
		var in Lead
		if err := c.BindJSON(&in); err != nil {
			c.JSON(400, gin.H{"error":"invalid json"}); return
		}
		if in.VendorID == 0 || strings.TrimSpace(in.OrganizerName) == "" || strings.TrimSpace(in.OrganizerEmail) == "" {
			c.JSON(400, gin.H{"error":"vendor_id, organizer_name, organizer_email required"}); return
		}
		if err := svc.Create(&in); err != nil {
			c.JSON(500, gin.H{"error":"db error"}); return
		}
		c.JSON(201, gin.H{"ok": true, "lead_id": in.ID})
	})
}
