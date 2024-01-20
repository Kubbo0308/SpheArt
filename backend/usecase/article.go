package usecase

import (
	"backend/domain/model"
	"backend/domain/repository"
)

type ArticleUsecase interface {
	GetAllArticles() ([]model.Article, error)
}

type articleUsecase struct {
	Repository repository.ArticleRepository
}

func NewArticleUsecase(repository repository.ArticleRepository) ArticleUsecase {
	return &articleUsecase{Repository: repository}
}

func (au *articleUsecase) GetAllArticles() ([]model.Article, error) {
	return []model.Article{}, nil
}
