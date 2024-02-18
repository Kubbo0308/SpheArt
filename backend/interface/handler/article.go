package handler

import (
	"backend/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ArticleHandler interface {
	GetAllArticles(ctx echo.Context) error
	SearchInArticleTitle(ctx echo.Context) error
}

type articleHandler struct {
	au usecase.ArticleUsecase
}

func NewArticleHandler(au usecase.ArticleUsecase) ArticleHandler {
	return &articleHandler{au}
}

func (ah *articleHandler) GetAllArticles(ctx echo.Context) error {
	resp, err := ah.au.GetAllArticles()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (ah *articleHandler) SearchInArticleTitle(ctx echo.Context) error {
	queryParam := ctx.QueryParam("title")
	searchTitle := "%" + queryParam + "%"

	resp, err := ah.au.SearchInArticleTitle(searchTitle)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, resp)
}
