package usecase

import (
	"backend/domain/model"
	"backend/domain/repository"
)

type ArticleUsecase interface {
	ArticlesPerPage(pageNum int) ([]model.Article, error)
	AllArticles() ([]model.Article, error)
	SearchInArticleTitle(searchTitle string, pageNum int) ([]model.Article, error)
}

type articleUsecase struct {
	ar repository.ArticleRepository
}

func NewArticleUsecase(ar repository.ArticleRepository) ArticleUsecase {
	return &articleUsecase{ar}
}

func (au *articleUsecase) ArticlesPerPage(pageNum int) ([]model.Article, error) {
	articles, err := au.ar.ArticlesPerPages(pageNum)
	if err != nil {
		return []model.Article{}, err
	}
	return articles, nil
}

func (au *articleUsecase) AllArticles() ([]model.Article, error) {
	articles, err := au.ar.AllArticles()
	if err != nil {
		return []model.Article{}, err
	}
	return articles, nil
}

func (au *articleUsecase) SearchInArticleTitle(searchTitle string, pageNum int) ([]model.Article, error) {
	articles, err := au.ar.SearchInArticleTitle(searchTitle, pageNum)
	if err != nil {
		return []model.Article{}, err
	}
	return articles, nil
}
