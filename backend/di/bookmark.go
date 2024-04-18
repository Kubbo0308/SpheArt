package di

import (
	"backend/infrastrcuture/persistence"
	handler "backend/interface/handler/echo"
	"backend/usecase"

	"gorm.io/gorm"
)

func Bookmark(db *gorm.DB) handler.BookmarkHandler {
	bp := persistence.NewBookmarkPersistence(db)
	bu := usecase.NewBookmarkUsecase(bp)
	bh := handler.NewBookmarkHandler(bu)
	return bh
}
