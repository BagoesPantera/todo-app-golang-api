package models

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

type DatabaseConfig struct {
	name string
	user string
	pass string
	host string
	port string
}

func getEnv(key, defaults string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaults
}

func ConnectDatabase() {
	if err := godotenv.Load(".env"); err != nil {
		panic("Cannot load .env file")
	}
	var dbconfig DatabaseConfig
	dbconfig.name = getEnv("DB_NAME", "golang_todo")
	dbconfig.user = getEnv("DB_USER", "root")
	dbconfig.pass = getEnv("DB_PASS", "")
	dbconfig.host = getEnv("DB_HOST", "localhost")
	dbconfig.port = getEnv("DB_PORT", "3306")
	config := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbconfig.user, dbconfig.pass, dbconfig.host, dbconfig.port, dbconfig.name)

	database, err := gorm.Open(mysql.Open(config))

	if err != nil {
		panic(err)
	}

	err = database.AutoMigrate(&Todo{})
	if err != nil {
		panic(err)
	}

	DB = database
}
