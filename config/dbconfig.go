package config

import (
	"fmt"
	"log"
	"os"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"todolist-BE/models"
)

var DB *gorm.DB

func SetupDatabaseConnection() *gorm.DB {
	err := godotenv.Load(".env")
	if err!= nil {
        log.Fatal("Error loading environment")
    }
	
	db_username := os.Getenv("MYSQL_USER")
	db_password := os.Getenv("MYSQL_PASSWORD")
	db_host := os.Getenv("MYSQL_HOST")
	db_port := os.Getenv("MYSQL_PORT")
	db_name := os.Getenv("MYSQL_DBNAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", db_username, db_password, db_host, db_port, db_name)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
        panic("failed connect to database")
    }

	database.AutoMigrate(&models.Activity{})
	database.AutoMigrate(&models.Todos{})

	DB = database
	return DB
}