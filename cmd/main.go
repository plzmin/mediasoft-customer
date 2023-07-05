package main

import (
	"github.com/caarlos0/env/v9"
	"mediasoft-customer/internal/app"
	"mediasoft-customer/internal/config"
	"mediasoft-customer/pkg/logger"
)

func main() {
	log := logger.New()

	cfg := config.Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal("failed to retrieve env variables %v", err)
	}

	if err := app.Run(log, cfg); err != nil {
		log.Fatal("error running gateway server %v", err)
	}
}
