package port

import (
	"github.com/Aorts/fiber-boiler/internal/core/domain"
	"github.com/gofiber/fiber/v2"
)

type UserRepository interface {
	GetById(c *fiber.Ctx, id string) (*domain.User, error)
	GetByUsername(c *fiber.Ctx, username string) (*domain.User, error)
	Create(c *fiber.Ctx, in *domain.User) (*domain.User, error)
	Update(c *fiber.Ctx, in *domain.User) error
	Delete(c *fiber.Ctx, id string) error
}
