package articles

import (
	"backend/batch"
	"backend/database"
	"backend/domain/model"
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		db := database.NewPostgreSQLDB()
		defer database.CloseDB(db)
		// マイグレーション
		if err := db.AutoMigrate(&model.Article{}, &model.User{}, &model.Bookmark{}); err != nil {
			return
		}
		batch.RunQiitaAPIBatch(db)
		batch.RunZennAPIBatch(db)
		fmt.Fprintf(w, "<h1>Run Batch!</h1>")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
