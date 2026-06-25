package main

import (
	"log"

	"github.com/AbdalrhmanAmmar/erp-dahboard-golang/internal/config"
	"github.com/AbdalrhmanAmmar/erp-dahboard-golang/internal/database"
	"github.com/AbdalrhmanAmmar/erp-dahboard-golang/internal/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	if err := database.Connect(cfg.DatabaseURL); err != nil {
		log.Fatalf("database error: %v", err)
	}

	srv := server.New()
	log.Printf("server running on port %s", cfg.AppPort)
	if err := srv.Run(":" + cfg.AppPort); err != nil {
		log.Fatalf("server error: %v", err)
	}
}