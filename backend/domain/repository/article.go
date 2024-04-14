package repository

import "backend/domain/model"

type ArticleRepository interface {
	ArticlesPerPages(pageNum int) ([]model.Article, error)
	AllArticles() ([]model.Article, error)
	SearchInArticleTitle(searchTitle string, pageNum int) ([]model.Article, error)
}
