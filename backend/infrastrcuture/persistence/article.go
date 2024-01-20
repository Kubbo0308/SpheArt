package persistence

import (
	"backend/domain/model"
	"backend/domain/repository"

	"gorm.io/gorm"
)

type articlePersistence struct {
	database *gorm.DB
}

func NewArticlePersistence(db *gorm.DB) repository.ArticleRepository {
	return &articlePersistence{database: db}
}

func (ap *articlePersistence) GetAllArticles(articles *[]model.Article) error {
	return nil
}
