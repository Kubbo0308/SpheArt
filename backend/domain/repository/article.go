package repository

import "backend/domain/model"

type ArticleRepository interface {
	GetAllQiitaArticles() ([]model.Article, error)
}
