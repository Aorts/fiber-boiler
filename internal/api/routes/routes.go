package routes

import (
	"github.com/Aorts/fiber-boiler/internal/api/middlerware"
	"github.com/Aorts/fiber-boiler/internal/core/port"
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	App         *fiber.App
	UserHandler port.UserHandler
	AuthHandler port.AuthHandler
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
	api.Post("/auths/login", c.AuthHandler.Login)
	api.Post("/auths/register", c.AuthHandler.Register)
	api.Get("/users/me", c.Middleware.AuthMiddleware(), c.UserHandler.Me)
	api.Put("/users/me", c.Middleware.AuthMiddleware(), c.UserHandler.UpdateUser)
}
