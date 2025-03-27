package usersvc

import (
	"github.com/Aorts/fiber-boiler/internal/core/domain"
	"github.com/Aorts/fiber-boiler/internal/core/port"
	"github.com/gofiber/fiber/v2"
)

type service struct {
	ur port.UserRepository
}

func NewUserService(ur port.UserRepository) port.UserService {
	return &service{
		ur: ur,
	}
}

// Delete implements port.UserService.
func (s *service) Delete(c *fiber.Ctx, id string) error {
	panic("unimplemented")
}

func (s *service) Create(c *fiber.Ctx, user *domain.User) error {
	panic("unimplemented")
}

// GetByEmail implements port.UserService.
func (s *service) GetByEmail(c *fiber.Ctx, email string) (*domain.User, error) {
	panic("unimplemented")
}

// GetById implements port.UserService.
func (s *service) GetById(c *fiber.Ctx, id string) (*domain.User, error) {
	panic("unimplemented")
}

// GetByUsername implements port.UserService.
func (s *service) GetByUsername(c *fiber.Ctx, username string) (*domain.User, error) {
	panic("unimplemented")
}

// Update implements port.UserService.
func (s *service) Update(c *fiber.Ctx, user *domain.User) error {
	panic("unimplemented")
}
