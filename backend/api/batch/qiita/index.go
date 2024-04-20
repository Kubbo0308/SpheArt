package articles

import (
	"backend/batch"
	"backend/database"
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		db := database.NewPostgreSQLDB()
		defer database.CloseDB(db)
		batch.RunQiitaAPIBatch(db)
		fmt.Fprintf(w, "<h1>Run Qiita Batch!</h1>")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
