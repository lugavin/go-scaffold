package main

import (
	"log"

	"github.com/lugavin/go-easy/config"
	"github.com/lugavin/go-easy/internal/app"
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
