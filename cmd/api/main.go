package main

import (
	"log"

	"github.com/AbdalrhmanAmmar/erp-dahboard-golang/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config error: %v", err)
	}
	_ = cfg // هتستخدمها في الاتصال بالداتابيز
	// ... باقي الإعداد
}