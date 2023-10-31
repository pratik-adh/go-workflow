package service

import (
	"CRUD/config"
	models "CRUD/models"
	"fmt"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

// Insert New data into the database
func RegisterNewUser(user *models.Registration) (err error) {
	newUUID := uuid.New()
	v := validator.New()

	    // Validate the user struct
    if err := v.Struct(user); err != nil {
        return fmt.Errorf(err.Error())
    } 

    user.ID = newUUID
	if err = config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}