package port

import "github.com/Aorts/fiber-boiler/internal/core/domain"

type UserRepository interface {
	GetById(id string) (*domain.User, error)
	GetByUsername(username string) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
	Create(in *domain.User) error
	Update(in *domain.User) error
	Delete(id string) error
}
