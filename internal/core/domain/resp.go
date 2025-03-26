package domain

import "github.com/gofiber/fiber/v2"

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func SuccessResponse(ctx *fiber.Ctx, data any, statusCode int, message string) error {
	return ctx.Status(statusCode).JSON(Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(ctx *fiber.Ctx, statusCode int, message string, err error) error {
	return ctx.Status(statusCode).JSON(Response{
		Success: false,
		Message: message,
		Error:   err.Error(),
	})
}
