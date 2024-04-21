package articles

import (
	"backend/database"
	"backend/infrastrcuture/persistence"
	handler "backend/interface/handler/http"
	"backend/usecase"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// OPTIONSリクエストへの対応
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	switch r.Method {
	case http.MethodPost:
		db := database.NewPostgreSQLDB()
		defer database.CloseDB(db)
		ap := persistence.NewUserPersistence(db)
		au := usecase.NewUserUsecase(ap)
		ah := handler.NewUserHandler(au)
		ah.SignUp(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
