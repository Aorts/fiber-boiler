package routes

import (
	"github.com/Aorts/fiber-boiler/internal/api/middlerware"
	"github.com/Aorts/fiber-boiler/internal/core/port"
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	App         *fiber.App
	UserHandler port.UserHandler
	Middleware  middlerware.Middleware
}

func (c *Config) Setup() {
	c.App.Use(c.Middleware.CORSMiddleware())
	c.Routes()
}

func (c *Config) Routes() {
	c.App.Get("/healthz", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "ok",
		})
	})
	api := c.App.Group("api/v1")
	api.Post("/user/refresh", c.UserHandler.Login)
	api.Post("/user/login", c.UserHandler.Login)
	api.Post("/user/register", c.UserHandler.Register)
	api.Post("/user/forget-password", c.UserHandler.ForgetPassword)
	api.Get("/user/me", c.Middleware.AuthMiddleware(), c.UserHandler.Me)
	api.Put("/user/me", c.Middleware.AuthMiddleware(), c.UserHandler.UpdateUser)
}
