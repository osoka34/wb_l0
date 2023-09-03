package controller

import (
	"github.com/gofiber/fiber/v2"
	"wb_l0/internal/models"
	"wb_l0/pkg/utils"
)

func (h *OrderHandler) GetOrder() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params = models.GetParams{}

		if err := utils.ReadRequestHeaderJson(ctx, &params); err != nil {
			h.logger.Errorf("err is: %v", err)
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		toFront, err := h.uc.GetOrder(&params)
		if err != nil {
			h.logger.Errorf("err is: %v", err)
			return ctx.Status(fiber.StatusInternalServerError).JSON(toFront)
		}

		return ctx.JSON(toFront)
	}
}
