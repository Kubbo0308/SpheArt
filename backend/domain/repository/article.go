package repository

import "backend/domain/model"

type ArticleRepository interface {
	GetAllArticles(articles *[]model.Article) error
}
