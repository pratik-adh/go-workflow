package main

import (
	"CRUD/api/routes"
	"CRUD/config"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	res, _ := json.Marshal(config.DB)
	fmt.Println("database configurations", string(res))
	defer config.DB.Close()
	r := routes.SetupRouter()
	// //running
	r.Run()
}
