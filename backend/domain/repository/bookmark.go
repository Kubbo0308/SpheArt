package repository

import "backend/domain/model"

type BookmarkRepository interface {
	AllBookmarkByUserId(userId uint) ([]model.Bookmark, error)
	CreateBookmark(bookmark *model.Bookmark) error
	DeleteBookmark(bookmark *model.Bookmark) error
}
