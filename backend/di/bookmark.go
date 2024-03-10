package di

import (
	"backend/infrastrcuture/persistence"
	"backend/interface/handler"
	"backend/usecase"

	"gorm.io/gorm"
)

func Bookmark(db *gorm.DB) handler.BookmarkHandler {
	bp := persistence.NewBookmarkPersistence(db)
	bu := usecase.NewBookmarkUsecase(bp)
	bh := handler.NewBookmarkHandler(bu)
	return bh
}
