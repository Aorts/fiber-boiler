package userhdl

import (
	"github.com/Aorts/fiber-boiler/internal/core/port"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usersvc port.UserService
	authsvc port.AuthService
}

func NewUserHandler(authsvc port.AuthService, usersvc port.UserService) port.UserHandler {
	return &handler{
		authsvc: authsvc,
		usersvc: usersvc,
	}
}

func (h *handler) Me(c *fiber.Ctx) error {
	panic("unimplemented")
}

func (h *handler) UpdateUser(c *fiber.Ctx) error {
	panic("unimplemented")
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

func (h *handler) RefreshToken(c *fiber.Ctx) error {
	panic("unimplemented")
}
