package main

import (
	"backend/batch"
	"backend/database"
	"backend/di"
	"backend/router"
	"fmt"
	"time"
)

func main() {
	fmt.Println("run!!!")
	db := database.NewDB()
	defer database.CloseDB(db)

	ticker := time.NewTicker(3 * time.Hour)
	defer ticker.Stop()

	ch1 := make(chan bool)
	ch2 := make(chan bool)

	fmt.Println("Start!")

	go func() {
		batch.RunQiitaAPIBatch(db)
		ch1 <- true
	}()

	go func() {
		batch.RunZennAPIBatch(db)
		ch2 <- true
	}()

	<-ch1
	<-ch2

	fmt.Println("Finish!")

	e := router.NewRouter(di.Article(db), di.User(db))
	e.Logger.Fatal(e.Start(":8080"))
}
