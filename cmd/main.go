package main

import (
	"github.com/ilyakaznacheev/cleanenv"
	"mediasoft-customer/internal/app"
	"mediasoft-customer/internal/config"
	"mediasoft-customer/pkg/logger"
)

func main() {
	log := logger.New()

	cfg := config.Config{}
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatal("failed to retrieve env variables %v", err)
	}

	if err := app.Run(log, cfg); err != nil {
		log.Fatal("error running gateway server %v", err)
	}
}
