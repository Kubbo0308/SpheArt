package articles

import (
	"backend/database"
	"backend/infrastrcuture/persistence"
	handler "backend/interface/handler/http"
	"backend/usecase"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		db := database.NewPostgreSQLDB()
		defer database.CloseDB(db)
		ap := persistence.NewUserPersistence(db)
		au := usecase.NewUserUsecase(ap)
		ah := handler.NewUserHandler(au)
		ah.SignOut(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
