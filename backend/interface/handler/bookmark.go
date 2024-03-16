package handler

import (
	"backend/usecase"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type BookmarkHandler interface {
	AllBookmark(ctx echo.Context) error
	PostBookmark(ctx echo.Context) error
	DeleteBookmark(ctx echo.Context) error
}

type bookmarkHandler struct {
	bu usecase.BookmarkUsecase
}

func NewBookmarkHandler(bu usecase.BookmarkUsecase) BookmarkHandler {
	return &bookmarkHandler{bu}
}

func (bh *bookmarkHandler) AllBookmark(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	res, err := bh.bu.AllBookmark(uint(userId.(float64)))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
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
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusCreated, res)
}

func (bh *bookmarkHandler) DeleteBookmark(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	articleId := ctx.Param("articleId")

	res, err := bh.bu.DeleteBookmark(uint(userId.(float64)), articleId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, res)
}
