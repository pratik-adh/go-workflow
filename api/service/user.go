package service

import (
	"CRUD/config"
	"CRUD/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllUsers(user *[]models.User) (err error) {
	if err = config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

type Error struct {
    Number  int    `json:"Number"`
    Message string `json:"Message"`
}

func CreateUser(c *gin.Context, user *models.User) (err error) {
	if err = config.DB.Create(user).Error; err != nil {
		byteErr, _ := json.Marshal(err)
		var newError Error
		json.Unmarshal((byteErr), &newError)
		c.JSON(http.StatusNotFound, gin.H{"error": newError.Message})
		// errors.GetErrorResponse(newError.Number, newError.Message)
		fmt.Println(newError.Message)
		return err
	}
	return nil
}

func GetUserByID(user *models.User, id string) (err error) {
	if err = config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *models.User, id string) (err error) {
	fmt.Println(user)
	config.DB.Save(user)
	return nil
}

func DeleteUser(user *models.User, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(user)
	return nil
}
