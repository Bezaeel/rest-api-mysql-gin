package main

import (
	"github.com/bezaeel/rest-api-mysql-gin/Config"
	"github.com/bezaeel/rest-api-mysql-gin/Models"
	Routes "github.com/bezaeel/rest-api-mysql-gin/Routes"

	"fmt"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql",
		Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Contact{})

	r := Routes.SetUpRoutes()
	r.Run(":5000")
}
