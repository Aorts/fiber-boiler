package userrepo

import "go.mongodb.org/mongo-driver/v2/bson"

type model struct {
	Id       bson.ObjectID `bson:"_id,omitempty"`
	Username string        `bson:"username"`
	Email    string        `bson:"email"`
	Password string        `bson:"password"`
}
