package usecase

import (
	"backend/domain/model"
	"backend/domain/repository"
)

type BookmarkUsecase interface {
	AllBookmark(userId uint) ([]model.Bookmark, error)
	PostBookmark(userId uint, articleId uint) error
	DeleteBookmark(userId uint, articleId uint) error
}

type bookmarkUsecase struct {
	br repository.BookmarkRepository
}

func NewBookmarkUsecase(br repository.BookmarkRepository) BookmarkUsecase {
	return &bookmarkUsecase{br}
}

func (bu *bookmarkUsecase) AllBookmark(userId uint) ([]model.Bookmark, error) {
	bookmarks, err := bu.br.AllBookmarkByUserId(userId)
	if err != nil {
		return []model.Bookmark{}, err
	}
	return bookmarks, nil
}

func (bu *bookmarkUsecase) PostBookmark(userId uint, articleId uint) error {
	err := bu.br.CreateBookmark(userId, articleId)
	if err != nil {
		return err
	}
	return nil
}

func (bu *bookmarkUsecase) DeleteBookmark(userId uint, articleId uint) error {
	err := bu.br.DeleteBookmark(userId, articleId)
	if err != nil {
		return err
	}
	return nil
}
