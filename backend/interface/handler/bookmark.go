package handler

import (
	"backend/usecase"

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
	return nil
}

func (bh *bookmarkHandler) PostBookmark(ctx echo.Context) error {
	return nil
}

func (bh *bookmarkHandler) DeleteBookmark(ctx echo.Context) error {
	return nil
}
