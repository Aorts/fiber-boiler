package userrepo

import (
	"time"

	"github.com/Aorts/fiber-boiler/internal/core/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type model struct {
	Id        bson.ObjectID `bson:"_id,omitempty"`
	Username  string        `bson:"username"`
	Password  string        `bson:"password"`
	CreatedAt time.Time     `bson:"createdAat"`
	UpdatedAt time.Time     `bson:"updatedAt"`
}

func (r *repository) toModel(in *domain.User) (*model, error) {
	oid, err := bson.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, err
	}
	return &model{
		Id:        oid,
		Username:  in.Username,
		Password:  in.Password,
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
	}, nil
}

func (r *repository) toDomain(in model) *domain.User {
	return &domain.User{
		Id:        in.Id.Hex(),
		Username:  in.Username,
		Password:  in.Password,
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
	}
}
