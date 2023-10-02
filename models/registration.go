package models

type Registration struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      string `json:"age"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}

func (b *Registration) TableName() string {
	return "registration_details"
}
