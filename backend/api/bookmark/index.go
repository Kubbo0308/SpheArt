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
	case http.MethodGet:
		db := database.NewPostgreSQLDB()
		defer database.CloseDB(db)
		ap := persistence.NewBookmarkPersistence(db)
		au := usecase.NewBookmarkUsecase(ap)
		ah := handler.NewBookmarkHandler(au)
		ah.BookmarkPerPage(w, r)
	case http.MethodPost:
		db := database.NewPostgreSQLDB()
		defer database.CloseDB(db)
		ap := persistence.NewBookmarkPersistence(db)
		au := usecase.NewBookmarkUsecase(ap)
		ah := handler.NewBookmarkHandler(au)
		ah.PostBookmark(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
