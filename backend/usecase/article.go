package usecase

import (
	"backend/domain/model"
	"backend/domain/repository"
)

type ArticleUsecase interface {
	GetAllArticles() ([]model.Article, error)
}

type articleUsecase struct {
	ar repository.ArticleRepository
}

func NewArticleUsecase(ar repository.ArticleRepository) ArticleUsecase {
	return &articleUsecase{ar}
}

func (au *articleUsecase) GetAllArticles() ([]model.Article, error) {
	resp, err := au.ar.GetAllQiitaArticles()
	if err != nil {
		return []model.Article{}, err
	}
	return resp, nil
}
