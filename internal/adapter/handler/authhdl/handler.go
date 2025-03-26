package authhdl

import (
	"github.com/Aorts/fiber-boiler/internal/core/port"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	ausvc port.AuthService
}

func NewAuthHandler(ausvc port.AuthService) port.AuthHandler {
	return &handler{
		ausvc: ausvc,
	}
}

func (h *handler) ForgetPassword(c *fiber.Ctx) error {
	panic("unimplemented")
}

func (h *handler) Login(c *fiber.Ctx) error {
	panic("unimplemented")
}

func (h *handler) Register(c *fiber.Ctx) error {
	panic("unimplemented")
}
