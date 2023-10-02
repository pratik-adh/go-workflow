package service

import (
	"CRUD/config"
	models "CRUD/models"

	_ "github.com/go-sql-driver/mysql"
)

//Register NEW USER ... Insert New data
func RegisterNewUser(user *models.Registration) (err error) {
	if err = config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}
