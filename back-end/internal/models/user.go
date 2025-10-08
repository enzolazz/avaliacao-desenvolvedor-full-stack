package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `bson:"_id" json:"id"`
	Name     string    `bson:"name" json:"name"`
	Surname  string    `bson:"surname" json:"surname"`
	Username string    `bson:"username" json:"username"`
	Password string    `bson:"password" json:"-"`
}

func NewUser(name, surname, username, password string) User {
	return User{
		ID:       uuid.New(),
		Name:     name,
		Surname:  surname,
		Username: username,
		Password: password,
	}
}
