package authsvc

import (
	"github.com/Aorts/fiber-boiler/internal/core/domain"
	"github.com/Aorts/fiber-boiler/internal/core/port"
	"github.com/gofiber/fiber/v2"
)

type service struct {
}

func NewAuthService() port.AuthService {
	return &service{}
}

func (s *service) ForgetPassword(c *fiber.Ctx, username string) error {
	panic("unimplemented")
}

func (s *service) Login(c *fiber.Ctx, username string, password string) (string, error) {
	panic("unimplemented")
}

func (s *service) Register(c *fiber.Ctx, user *domain.User) error {
	panic("unimplemented")
}
