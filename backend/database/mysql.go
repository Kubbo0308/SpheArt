package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQLDB() *gorm.DB {
	// MySQLに接続
	dsn := fmt.Sprintf(`%s:%s@tcp(db:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local`,
		os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_DATABASE"))
	if err := RetryConnectDB(mysql.Open(dsn), &gorm.Config{}, 100); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Connected")
	return db
}
