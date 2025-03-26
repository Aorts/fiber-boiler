package userhdl

import (
	"github.com/Aorts/fiber-boiler/internal/core/port"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usersvc port.UserService
}

func NewUserHandler(usersvc port.UserService) port.UserHandler {
	return &handler{
		usersvc: usersvc,
	}
}

func (h *handler) Me(c *fiber.Ctx) error {
	panic("unimplemented")
}

func (h *handler) UpdateUser(c *fiber.Ctx) error {
	panic("unimplemented")
}
