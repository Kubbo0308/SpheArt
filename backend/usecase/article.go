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
	zr repository.ZennRepository
}

func NewArticleUsecase(qr repository.QiitaRepository, zr repository.ZennRepository) ArticleUsecase {
	return &articleUsecase{qr, zr}
}

func (au *articleUsecase) GetAllArticles() ([]model.Article, error) {
	qiitaResp, err := au.qr.GetAllQiitaArticles()
	if err != nil {
		return []model.Article{}, err
	}
	zennResp, err := au.zr.GetAllZennArticles()
	if err != nil {
		return []model.Article{}, err
	}
	articles := append(qiitaResp, zennResp...)
	return articles, nil
}
