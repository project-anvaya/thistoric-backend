package vendors

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(r *gin.RouterGroup, db *gorm.DB) {
	svc := NewService(NewRepo(db))

	r.GET("/vendors", func(c *gin.Context) {
		limit, offset := 50, 0
		if v := c.Query("limit"); v != "" { fmt.Sscanf(v, "%d", &limit) }
		if v := c.Query("offset"); v != "" { fmt.Sscanf(v, "%d", &offset) }

		items, err := svc.List(c.Query("q"), c.Query("category"), c.Query("location"), limit, offset)
		if err != nil { c.JSON(500, gin.H{"error":"db error"}); return }

		type dto struct {
			ID uint `json:"id"`
			Name, Category, Location, Description, Phone, Website string
		}
		out := make([]dto, 0, len(items))
		for _, v := range items {
			out = append(out, dto{ID:v.ID, Name:v.Name, Category:v.Category, Location:v.Location,
				Description:v.Description, Phone:v.Phone, Website:v.Website})
		}
		c.JSON(200, out)
	})

	r.GET("/vendors/:id", func(c *gin.Context) {
		id64, _ := strconv.ParseUint(c.Param("id"), 10, 64)
		v, err := svc.Get(uint(id64))
		if err != nil { c.JSON(404, gin.H{"error":"not found"}); return }
		c.JSON(200, gin.H{
			"id": v.ID, "name": v.Name, "category": v.Category, "location": v.Location,
			"description": v.Description, "phone": v.Phone, "website": v.Website,
			"portfolio": []string{},
		})
	})
}
