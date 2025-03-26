package authsvc

import (
	"github.com/Aorts/fiber-boiler/internal/core/domain"
	"github.com/Aorts/fiber-boiler/internal/core/port"
)

type service struct {
}

func NewAuthService() port.AuthService {
	return &service{}
}

func (s *service) ForgetPassword(username string) error {
	panic("unimplemented")
}

func (s *service) Login(username string, password string) (string, error) {
	panic("unimplemented")
}

func (s *service) Register(user *domain.User) error {
	panic("unimplemented")
}
