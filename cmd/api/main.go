package main

import (
	"log"
	"thistoric-backend/internal/config"
	"thistoric-backend/internal/platform/http"
	"thistoric-backend/internal/adapters/db"
	"thistoric-backend/internal/adapters/log"
)

func main() {
	cfg := config.Load()
	logger := logz.New()
	gormDB := db.Connect(cfg.PostgresDSN)

	r := httpx.NewRouter(cfg, logger, gormDB)

	addr := ":" + cfg.Port
	logger.Infof("listening on %s", addr)
	if err := r.Run(addr); err != nil {
		logger.Errorf("server error: %v", err)
		log.Fatal(err)
	}
}
