package articles

import (
	"backend/database"
	"backend/domain/model"
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	db := database.NewPostgreSQLDB()
	defer database.CloseDB(db)
	// マイグレーション
	if err := db.AutoMigrate(&model.Article{}, &model.User{}, &model.Bookmark{}); err != nil {
		return
	}
	fmt.Fprintf(w, "<h1>Success Migration!</h1>")
}
