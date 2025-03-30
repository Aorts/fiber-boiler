package userhdl

import (
	"github.com/Aorts/fiber-boiler/internal/core/domain"
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
	req := new(domain.RegisterRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	user := domain.NewUser(req.Username, req.Email, req.Password)
	if err := h.authsvc.Register(c, user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "user created"})
}

func (h *handler) RefreshToken(c *fiber.Ctx) error {
	panic("unimplemented")
}
