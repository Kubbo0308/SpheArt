package main

import (
	"backend/database"
	"fmt"
)

func main() {
	fmt.Println("run!!!")
	db := database.NewDB()
	database.CloseDB(db)
}
