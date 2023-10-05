package models

import "github.com/google/uuid"

type User struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Age     int `json:"age"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	ReportingTo uuid.UUID `json:"reporting_to"`
}

func (b *User) TableName() string {
	return "user_details"
}
