package userrepo

import (
	"context"
	"time"

	"github.com/Aorts/fiber-boiler/config"
	"github.com/Aorts/fiber-boiler/internal/core/domain"
	"github.com/Aorts/fiber-boiler/internal/core/port"
	"github.com/gofiber/fiber/v2"
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
				{Key: "password", Value: 1},
			},
			Options: options.Index().SetUnique(true),
		},
	}); err != nil {
		panic(err)
	}
	return &repository{mc: mc, col: col}
}

func (r *repository) Create(c *fiber.Ctx, in *domain.User) (*domain.User, error) {
	model, err := r.toModel(in)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	model.CreatedAt = now
	model.UpdatedAt = now

	result, err := r.col.InsertOne(c.Context(), in)
	if err != nil {
		return nil, err
	}
	model.Id = result.InsertedID.(bson.ObjectID)
	return r.toDomain(*model), nil
}

func (r *repository) Delete(ictx *fiber.Ctx, d string) error {
	panic("unimplemented")
}

func (r *repository) GetById(ctx *fiber.Ctx, id string) (*domain.User, error) {
	panic("unimplemented")
}

func (r *repository) GetByUsername(ctx *fiber.Ctx, username string) (*domain.User, error) {
	panic("unimplemented")
}

func (r *repository) Update(ctx *fiber.Ctx, in *domain.User) error {
	panic("unimplemented")
}
