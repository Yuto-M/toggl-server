package main

import (
	"log"

	"github.com/Yuto-M/go-clean-template/config"
	"github.com/Yuto-M/go-clean-template/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
