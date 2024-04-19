package main

import (
	"backend/database"
	"backend/domain/model"
	"fmt"
)

func main() {
	// DBのインスタンスアドレスを取得
	// dbConn := database.NewMySQLDB()
	dbConn := database.NewPostgreSQLDB()
	defer fmt.Println("Successfully Migrated")
	defer database.CloseDB(dbConn)
	if err := dbConn.AutoMigrate(&model.Article{}, &model.User{}, &model.Bookmark{}); err != nil {
		return
	}
}
