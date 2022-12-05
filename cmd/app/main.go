package main

import (
	"flag"
	"log"

	"github.com/lugavin/go-scaffold/config"
	"github.com/lugavin/go-scaffold/internal/app"
)

var cfgPath = flag.String("c", "./config/config.yml", "Path to config file")

func main() {
	// Parse the command-line flags
	flag.Parse()

	// Configuration
	cfg, err := config.NewConfig(*cfgPath)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
