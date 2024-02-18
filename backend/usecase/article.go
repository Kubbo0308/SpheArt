package usecase

import (
	"backend/domain/model"
	"backend/domain/repository"
)

type ArticleUsecase interface {
	GetAllArticles() ([]model.Article, error)
	SearchInArticleTitle(searchTitle string) ([]model.Article, error)
}

type articleUsecase struct {
	ar repository.ArticleRepository
}

func NewArticleUsecase(ar repository.ArticleRepository) ArticleUsecase {
	return &articleUsecase{ar}
}

func (au *articleUsecase) GetAllArticles() ([]model.Article, error) {
	articles, err := au.ar.GetAllArticles()
	if err != nil {
		return []model.Article{}, err
	}
	return articles, nil
}

func (au *articleUsecase) SearchInArticleTitle(searchTitle string) ([]model.Article, error) {
	articles, err := au.ar.SearchInArticleTitle(searchTitle)
	if err != nil {
		return []model.Article{}, err
	}
	return articles, nil
}
