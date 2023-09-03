package main

import (
	"context"
	"fmt"
	"log"
	"wb_l0/config"
	"wb_l0/internal/httpServer"
	"wb_l0/pkg/logger"
)

func main() {

	ctx := context.Background()

	viperInstance, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Cannot load config. Error: {%s}", err.Error())
	}

	cfg, err := config.ParseConfig(viperInstance)
	if err != nil {
		log.Fatalf("Cannot parse config. Error: {%s}", err.Error())
	}

	logger := logger.NewLogger(cfg).Sugar()
	defer logger.Sync()

	s := httpServer.NewServer(cfg, logger)
	if err = s.Run(ctx); err != nil {
		logger.Error(fmt.Sprint(err))
	}

}
