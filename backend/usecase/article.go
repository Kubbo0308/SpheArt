package usecase

import (
	"backend/domain/model"
	"backend/domain/repository"
)

type ArticleUsecase interface {
	GetAllArticles() ([]model.Article, error)
}

type articleUsecase struct {
	articleRepository repository.ArticleRepository
}

func NewArticleUsecase(ar repository.ArticleRepository) ArticleUsecase {
	return &articleUsecase{articleRepository: ar}
}

func (au *articleUsecase) GetAllArticles() ([]model.Article, error) {
	return []model.Article{}, nil
}
