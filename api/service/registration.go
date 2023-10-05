package service

import (
	"CRUD/config"
	models "CRUD/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

// Insert New data into the database
func RegisterNewUser(user *models.Registration) (err error) {
	newUUID := uuid.New()
    user.ID = newUUID
	if err = config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}