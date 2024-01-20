package handler

import (
	"backend/usecase"

	"github.com/labstack/echo/v4"
)

type ArticleHandler interface {
	GetAllArticles(ctx echo.Context) error
}

type articleHandler struct {
	articleUsecase usecase.ArticleUsecase
}

func NewArticleHandler(au usecase.ArticleUsecase) ArticleHandler {
	return &articleHandler{articleUsecase: au}
}

func (ah *articleHandler) GetAllArticles(ctx echo.Context) error {
	return nil
}
