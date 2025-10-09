package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Surname   string             `bson:"surname" json:"surname"`
	Username  string             `bson:"username" json:"username"`
	Password  string             `bson:"password" json:"-"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

func NewUser(name, surname, username, password string) User {
	return User{
		ID:        primitive.NewObjectID(),
		Name:      name,
		Surname:   surname,
		Username:  username,
		Password:  password,
		CreatedAt: time.Now(),
	}
}
