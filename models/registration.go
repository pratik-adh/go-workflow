package models

import "github.com/google/uuid"

type Registration struct {
	ID   uuid.UUID `json:"id"`
	Name     string `json:"name"`
	Email    string `form:"email" validate:"required,email"`
	Password string `json:"password"`
	Age      string `json:"age"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}

func (b *Registration) TableName() string {
	return "registration_details"
}
