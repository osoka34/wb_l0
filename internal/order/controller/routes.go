package controller

import "github.com/gofiber/fiber/v2"

func MapOrderRoutes(router fiber.Router, h *OrderHandler) {
	router.Get("/order", h.GetOrder())
}
