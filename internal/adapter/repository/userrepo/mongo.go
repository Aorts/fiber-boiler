package userrepo

import (
	"context"

	"github.com/Aorts/fiber-boiler/config"
	"github.com/Aorts/fiber-boiler/internal/core/domain"
	"github.com/Aorts/fiber-boiler/internal/core/port"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type repository struct {
	mc  *mongo.Client
	col *mongo.Collection
}

func NewUserRepository(mc *mongo.Client) port.UserRepository {
	col := mc.Database(config.Get().Mongo.Database).Collection("users")
	if _, err := col.Indexes().CreateMany(context.Background(), []mongo.IndexModel{
		{
			Keys: bson.D{
				{Key: "username", Value: 1},
			},
			Options: options.Index().SetUnique(true),
		},
	}); err != nil {
		panic(err)
	}
	return &repository{mc: mc, col: col}
}

func (r *repository) Create(in *domain.User) error {
	panic("unimplemented")
}

func (r *repository) Delete(id string) error {
	panic("unimplemented")
}

func (r *repository) GetByEmail(email string) (*domain.User, error) {
	panic("unimplemented")
}

func (r *repository) GetById(id string) (*domain.User, error) {
	panic("unimplemented")
}

func (r *repository) GetByUsername(username string) (*domain.User, error) {
	panic("unimplemented")
}

func (r *repository) Update(in *domain.User) error {
	panic("unimplemented")
}
