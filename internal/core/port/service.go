package port

import (
	"github.com/Aorts/fiber-boiler/internal/core/domain"
	"github.com/gofiber/fiber/v2"
)

type UserService interface {
	GetById(c *fiber.Ctx, id string) (*domain.User, error)
	GetByUsername(c *fiber.Ctx, username string) (*domain.User, error)
	GetByEmail(c *fiber.Ctx, email string) (*domain.User, error)
	Create(c *fiber.Ctx, user *domain.User) error
	Update(c *fiber.Ctx, user *domain.User) error
	Delete(c *fiber.Ctx, id string) error
}

type AuthService interface {
	Login(c *fiber.Ctx, username, password string) (string, error)
	Register(c *fiber.Ctx, user *domain.User) error
	ForgetPassword(c *fiber.Ctx, username string) error
}
