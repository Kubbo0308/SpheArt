package articles

import (
	"backend/database"
	"backend/infrastrcuture/persistence"
	handler "backend/interface/handler/http"
	"backend/usecase"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	db := database.NewPostgreSQLDB()
	defer database.CloseDB(db)
	ap := persistence.NewBookmarkPersistence(db)
	au := usecase.NewBookmarkUsecase(ap)
	ah := handler.NewBookmarkHandler(au)
	switch r.Method {
	case http.MethodGet:
		ah.BookmarkPerPage(w, r)
	case http.MethodPost:
		ah.PostBookmark(w, r)
	}
}
