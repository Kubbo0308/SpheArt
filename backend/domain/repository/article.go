package repository

import "backend/domain/model"

type ArticleRepository interface {
	AllArticles() ([]model.Article, error)
	ArticlesPerPages(pageNum int) ([]model.Article, error)
	SearchInArticleTitle(searchTitle string) ([]model.Article, error)
}
