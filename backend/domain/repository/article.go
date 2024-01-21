package repository

import "backend/domain/model"

type ArticleRepository interface {
	GetAllQiitaArticles(articles *[]model.QiitaResponse) error
}
