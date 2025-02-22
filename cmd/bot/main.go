package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/timohahaa/storm-bot/config"
	"github.com/timohahaa/storm-bot/internal/service"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("[service] %v", err)
	}

	app, err := service.New(cfg)
	if err != nil {
		log.Fatalf("[service] %v", err)
	}

	if err := app.Run(); err != nil {
		log.Fatalf("[service] %v", err)
	}
}
