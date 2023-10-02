package service

import (
	"CRUD/config"
	"CRUD/models"
	"fmt"
)

type LoginService interface {
	LoginUser(email string, password string) bool
}

func LoginUser(email string, password string) bool {
	var user models.Registration
	fmt.Println("user hai hai", user)
	if err := config.DB.Where("email = ?", email).Where("password = ?", password).First(&user).Error; err != nil {
		return false
	}
	return true
}
