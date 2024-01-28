package di

import (
	"backend/handler"
	"backend/infrastrcuture/persistence"
	"backend/usecase"

	"gorm.io/gorm"
)

func Article(db *gorm.DB) handler.ArticleHandler {
	qp := persistence.NewQiitaArticlePersistence(db)
	zp := persistence.NewZennArticlePersistence(db)
	au := usecase.NewArticleUsecase(qp, zp)
	ah := handler.NewArticleHandler(au)
	return ah
}
