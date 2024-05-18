package handler

import (
	"backend/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type BookmarkHandler interface {
	BookmarkPerPage(ctx echo.Context) error
	AllBookmark(ctx echo.Context) error
	PostBookmark(ctx echo.Context) error
}

type bookmarkHandler struct {
	bu usecase.BookmarkUsecase
}

func NewBookmarkHandler(bu usecase.BookmarkUsecase) BookmarkHandler {
	return &bookmarkHandler{bu}
}

func (bh *bookmarkHandler) BookmarkPerPage(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	queryParam := ctx.QueryParam("per_page")
	perPage, err := strconv.Atoi(queryParam)
	if err != nil {
		// 変換に失敗した場合のエラーハンドリング
		// 例えば、デフォルト値を設定する、エラーレスポンスを返すなど
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid per_page parameter"})
	}

	res, err := bh.bu.BookmarkedArticlePerPage(uint(userId.(float64)), perPage)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, res)
}

func (bh *bookmarkHandler) AllBookmark(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	res, err := bh.bu.AllBookmarkedArticle(uint(userId.(float64)))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, res)
}

func (bh *bookmarkHandler) PostBookmark(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	articleId := ctx.Param("articleId")

	res, err := bh.bu.PostBookmark(uint(userId.(float64)), articleId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, res)
}
