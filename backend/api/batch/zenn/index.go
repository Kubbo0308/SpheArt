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
		batch.RunZennAPIBatch(db)
		fmt.Fprintf(w, "<h1>Run Zenn Batch!</h1>")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
