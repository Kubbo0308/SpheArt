package main

import (
	"backend/database"
	"backend/di"
	"backend/router"
	"fmt"
)

func main() {
	fmt.Println("run!!!")
	db := database.NewDB()
	defer database.CloseDB(db)

	e := router.NewRouter(di.Article(db))
	e.Logger.Fatal(e.Start(":8080"))
}
