package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgreSQLDB() *gorm.DB {
	// MySQLに接続
	dsn := fmt.Sprintf(`host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai`,
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DATABASE"), os.Getenv("POSTGRES_PORT"))
	if err := RetryConnectDB(postgres.Open(dsn), &gorm.Config{}, 100); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Connected")
	return db
}
