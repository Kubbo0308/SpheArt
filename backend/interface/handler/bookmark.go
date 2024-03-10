package handler

import (
	"backend/usecase"
	"net/http"
	"strconv"

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
	userId, err := strconv.Atoi(ctx.Param("userId"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	res, err := bh.bu.AllBookmark(uint(userId))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, res)
}

func (bh *bookmarkHandler) PostBookmark(ctx echo.Context) error {
	userId, err := strconv.Atoi(ctx.Param("userId"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	articleId := ctx.Param("articleId")

	res, err := bh.bu.PostBookmark(uint(userId), articleId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusCreated, res)
}

func (bh *bookmarkHandler) DeleteBookmark(ctx echo.Context) error {
	userIdInt, err := strconv.Atoi(ctx.Param("userId"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	articleId := ctx.Param("articleId")
	res, err := bh.bu.DeleteBookmark(uint(userIdInt), articleId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, res)
}
