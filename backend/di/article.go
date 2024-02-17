package di

import (
	"backend/handler"
	"backend/infrastrcuture/persistence"
	"backend/usecase"

	"gorm.io/gorm"
)

func Article(db *gorm.DB) handler.ArticleHandler {
	ap := persistence.NewArticlePersistence(db)
	au := usecase.NewArticleUsecase(ap)
	ah := handler.NewArticleHandler(au)
	return ah
}
