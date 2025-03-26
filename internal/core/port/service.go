package port

import "github.com/Aorts/fiber-boiler/internal/core/domain"

type UserService interface {
	GetById(id string) (*domain.User, error)
	GetByUsername(username string) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
	Create(user *domain.User) error
	Update(user *domain.User) error
	Delete(id string) error
}

type AuthService interface {
	Login(username, password string) (string, error)
	Register(user *domain.User) error
	ForgetPassword(username string) error
}
