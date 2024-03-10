package persistence

import (
	"backend/domain/model"
	"backend/domain/repository"

	"gorm.io/gorm"
)

type bookmarkPersistence struct {
	db *gorm.DB
}

func NewBookmarkPersistence(db *gorm.DB) repository.BookmarkRepository {
	return &bookmarkPersistence{db}
}

func (bp *bookmarkPersistence) AllBookmarkByUserId(userId uint) ([]model.Bookmark, error) {
	return []model.Bookmark{}, nil
}

func (bp *bookmarkPersistence) CreateBookmark(userId uint, articleId uint) error {
	return nil
}
