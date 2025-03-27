package port

import "github.com/gofiber/fiber/v2"

type UserHandler interface {
	Me(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
	ForgetPassword(c *fiber.Ctx) error
	RefreshToken(c *fiber.Ctx) error
}
