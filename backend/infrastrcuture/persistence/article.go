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

// 1ページあたりの記事数
const pageSize = 20

func (ap *articlePersistence) ArticlesPerPages(pageNum int) ([]model.Article, error) {
	articles := []model.Article{}
	// ページ番号から、OFFSETの計算を行う（ページ番号は1から始まる）
	offset := (pageNum - 1) * pageSize
	// LimitとOffsetメソッドを使ってページネーションを適用
	res := ap.db.Limit(pageSize).Offset(offset).Find(&articles)
	if res.Error != nil {
		return []model.Article{}, res.Error
	}
	return articles, nil
}

func (ap *articlePersistence) AllArticles() ([]model.Article, error) {
	articles := []model.Article{}
	res := ap.db.Find(&articles)
	if res.Error != nil {
		return []model.Article{}, res.Error
	}
	return articles, nil
}

func (ap *articlePersistence) SearchInArticleTitle(searchTitle string, pageNum int) ([]model.Article, error) {
	articles := []model.Article{}
	// ページ番号から、OFFSETの計算を行う（ページ番号は1から始まる）
	offset := (pageNum - 1) * pageSize
	res := ap.db.Where("title LIKE ?", searchTitle).Limit(pageSize).Offset(offset).Find(&articles)
	if res.Error != nil {
		return []model.Article{}, res.Error
	}
	return articles, nil
}
