package dtos

import "github.com/google/uuid"

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Surname  string `json:"surname" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterResponse struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Surname  string    `json:"surname"`
	Username string    `json:"username"`
}
