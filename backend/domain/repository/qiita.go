package repository

import "backend/domain/model"

type QiitaRepository interface {
	GetAllQiitaArticles() ([]model.Article, error)
}
