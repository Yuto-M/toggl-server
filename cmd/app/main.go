package main

import (
	"log"

	"github.com/Yuto-M/toggl-server/config"
	"github.com/Yuto-M/toggl-server/internal/app"
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
