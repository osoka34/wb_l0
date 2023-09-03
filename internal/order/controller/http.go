package controller

import (
	"go.uber.org/zap"
	"wb_l0/config"
	"wb_l0/internal/order"
)

type OrderHandler struct {
	uc     order.Usecase
	logger *zap.SugaredLogger
	cfg    *config.Config
}

func NewOrderHandler(uc order.Usecase, cfg *config.Config, logger *zap.SugaredLogger) *OrderHandler {
	return &OrderHandler{uc: uc, cfg: cfg, logger: logger}
}
