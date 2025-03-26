package usersvc

import (
	"github.com/Aorts/fiber-boiler/internal/core/domain"
	"github.com/Aorts/fiber-boiler/internal/core/port"
)

type service struct {
	ur port.UserRepository
}

func NewUserService(ur port.UserRepository) port.UserService {
	return &service{
		ur: ur,
	}
}

func (s *service) Create(user *domain.User) error {
	panic("unimplemented")
}

func (s *service) Delete(id string) error {
	panic("unimplemented")
}

func (s *service) GetByEmail(email string) (*domain.User, error) {
	panic("unimplemented")
}

func (s *service) GetById(id string) (*domain.User, error) {
	panic("unimplemented")
}

func (s *service) GetByUsername(username string) (*domain.User, error) {
	panic("unimplemented")
}

func (s *service) Update(user *domain.User) error {
	panic("unimplemented")
}
