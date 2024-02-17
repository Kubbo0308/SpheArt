package repository

import "backend/domain/model"

type ArticleRepository interface {
	GetAllArticles() ([]model.Article, error)
}