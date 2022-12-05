package main

import (
	"log"

	"github.com/lugavin/go-scaffold/config"
	"github.com/lugavin/go-scaffold/internal/app"
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
