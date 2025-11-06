package httpx

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"thistoric-backend/internal/adapters/log"
	"thistoric-backend/internal/config"
	"thistoric-backend/internal/health"
	"thistoric-backend/internal/platform/http/middleware"
	"thistoric-backend/internal/domains/vendors"
	"thistoric-backend/internal/domains/leads"
	"thistoric-backend/internal/domains/users"
)

func NewRouter(cfg config.Config, logger *logz.Logger, db *gorm.DB) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.RequestLogger())
	r.Use(middleware.CORS())

	api := r.Group("")
	health.Register(api)
	vendors.Register(api, db)
	leads.Register(api, db)
	users.Register(api) // /me (protected)

	return r
}
