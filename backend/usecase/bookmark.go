package usecase

import (
	"backend/domain/model"
	"backend/domain/repository"
)

type BookmarkUsecase interface {
	AllBookmark(userId uint) ([]model.Bookmark, error)
	PostBookmark(userId uint, articleId string) (model.Bookmark, error)
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

func (bu *bookmarkUsecase) PostBookmark(userId uint, articleId string) (model.Bookmark, error) {
	newBookmark := model.Bookmark{
		UserID:    userId,
		ArticleID: articleId,
	}
	err := bu.br.PostBookmark(&newBookmark)
	if err != nil {
		return model.Bookmark{}, err
	}
	return newBookmark, nil
}
