package handler

import (
	"backend/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ArticleHandler interface {
	ArticlesPerPage(ctx echo.Context) error
	AllArticles(ctx echo.Context) error
	SearchInArticleTitle(ctx echo.Context) error
}

type articleHandler struct {
	au usecase.ArticleUsecase
}

func NewArticleHandler(au usecase.ArticleUsecase) ArticleHandler {
	return &articleHandler{au}
}

func (ah *articleHandler) ArticlesPerPage(ctx echo.Context) error {
	queryParam := ctx.QueryParam("per_page")
	perPage, err := strconv.Atoi(queryParam)
	if err != nil {
		// 変換に失敗した場合のエラーハンドリング
		// 例えば、デフォルト値を設定する、エラーレスポンスを返すなど
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid per_page parameter"})
	}

	res, err := ah.au.ArticlesPerPage(perPage)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, res)
}

func (ah *articleHandler) AllArticles(ctx echo.Context) error {
	res, err := ah.au.AllArticles()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, res)
}

func (ah *articleHandler) SearchInArticleTitle(ctx echo.Context) error {
	queryParam := ctx.QueryParam("title")
	searchTitle := "%" + queryParam + "%"

	res, err := ah.au.SearchInArticleTitle(searchTitle)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, res)
}
