package database

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID         primitive.ObjectID `bson:"id"`
	Name       string             `bson:"name"`
	Lastname   string             `bson:"lastname"`
	Username   string             `bson:"username"`
	Password   string             `bson:"password"`
	Token      string             `bson:"token"`
	ExpireAt   time.Time          `bson:"expire_at"`
	Motorcycle []Motorcycle       `bson:"motorcycles"`
	CreatedAt  time.Time          `bson:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at"`
}
