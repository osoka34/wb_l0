package httpServer

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"wb_l0/config"
)

type Server struct {
	logger *zap.SugaredLogger
	cfg    *config.Config
	fiber  *fiber.App
}

func NewServer(cfg *config.Config, logger *zap.SugaredLogger) *Server {
	return &Server{
		fiber:  fiber.New(fiber.Config{DisableStartupMessage: true}),
		logger: logger,
		cfg:    cfg,
	}
}

func (s *Server) Run(ctx context.Context) error {
	if err := s.MapHandlers(ctx, s.fiber, s.logger); err != nil {
		s.logger.Fatal(fmt.Sprintf("Cannot map handlers. Error: {%s}", err))
	}

	s.logger.Info(fmt.Sprintf("Start server on {host:port - %s:%s}", s.cfg.Server.Host, s.cfg.Server.Port))

	if err := s.fiber.Listen(fmt.Sprintf("%s:%s", s.cfg.Server.Host, s.cfg.Server.Port)); err != nil {
		s.logger.Fatal(fmt.Sprintf("Cannot listen. Error: {%s}", err))
	}

	return nil
}
