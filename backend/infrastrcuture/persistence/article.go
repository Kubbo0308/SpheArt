package persistence

import (
	"backend/domain/model"
	"backend/domain/repository"

	"gorm.io/gorm"
)

type articlePersistence struct {
	db *gorm.DB
}

func NewArticlePersistence(db *gorm.DB) repository.ArticleRepository {
	return &articlePersistence{db}
}

func (ap *articlePersistence) GetAllArticles() ([]model.Article, error) {
	articles := []model.Article{}
	res := ap.db.Find(&articles)
	if res.Error != nil {
		return []model.Article{}, res.Error
	}
	return articles, nil
}

func (ap *articlePersistence) SearchInArticleTitle(searchTitle string) ([]model.Article, error) {
	articles := []model.Article{}
	res := ap.db.Where("title LIKE ?", searchTitle).Find(&articles)
	if res.Error != nil {
		return []model.Article{}, res.Error
	}
	return articles, nil
}
