package repository

import "backend/domain/model"

type ZennRepository interface {
	GetAllZennArticles() ([]model.Article, error)
}
