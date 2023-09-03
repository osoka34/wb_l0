package utils

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate *validator.Validate = validator.New()

func ReadRequestHeader(c *fiber.Ctx, request interface{}) error {
	if err := c.ReqHeaderParser(request); err != nil {
		return err
	}

	return validate.StructCtx(c.Context(), request)
}

func ReadRequestHeaderJson(c *fiber.Ctx, request interface{}) error {
	if err := json.Unmarshal(c.Request().Body(), request); err != nil {
		return err
	}

	return validate.StructCtx(c.Context(), request)
}
