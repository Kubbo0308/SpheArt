package persistence

import (
	"backend/domain/model"
	"backend/domain/repository"

	"gorm.io/gorm"
)

type articlePersistence struct {
	Db *gorm.DB
}

func NewArticlePersistence(db *gorm.DB) repository.ArticleRepository {
	return &articlePersistence{Db: db}
}

func (ap *articlePersistence) GetAllArticles(articles *[]model.Article) error {
	return nil
}
