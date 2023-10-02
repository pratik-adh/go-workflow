package models

type User struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Age     string `json:"age"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

func (b *User) TableName() string {
	return "user_details"
}
