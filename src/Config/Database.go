package Config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

type DbConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	password string
}

func BuildDBConfig() *DbConfig {
	dbConfig := DbConfig{
		Host:     "localhost",
		Port:     3306,
		User:     "root",
		password: "",
		DBName:   "ContactManager",
	}

	return &dbConfig
}

func DbURL(dbConfig *DbConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
