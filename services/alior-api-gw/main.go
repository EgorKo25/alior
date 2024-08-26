package main

import (
	"context"
	l "log"

	"API-GW/src/config"
	"API-GW/src/server"
	"API-GW/src/server/command"

	"github.com/EgorKo25/common/logger"
)

func main() {
	log, err := logger.NewLogger(logger.PRODUCTION)
	if err != nil {
		l.Fatalf("cannot create logger: %s", err.Error())
	}
	cfg, err := config.NewAPIGWConfig()
	if err != nil {
		log.Fatal("cannot load configuration with error: %s", err.Error())
	}
	manager := command.NewManager(log)

	s := server.NewServer(cfg.ServerConfig, manager, log)
	if err = s.Run(context.Background()); err != nil {
		log.Fatal("%s", err.Error())
	}
}
