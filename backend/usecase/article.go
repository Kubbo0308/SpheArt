package usecase

import (
	"backend/domain/model"
	"backend/domain/repository"
)

type ArticleUsecase interface {
	GetAllArticles() ([]model.Article, error)
}

type articleUsecase struct {
	qr repository.QiitaRepository
}

func NewArticleUsecase(qr repository.QiitaRepository) ArticleUsecase {
	return &articleUsecase{qr}
}

func (au *articleUsecase) GetAllArticles() ([]model.Article, error) {
	resp, err := au.qr.GetAllQiitaArticles()
	if err != nil {
		return []model.Article{}, err
	}
	return resp, nil
}
