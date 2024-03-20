package repository

import "backend/domain/model"

type BookmarkRepository interface {
	AllBookmarkByUserId(userId uint) ([]model.Bookmark, error)
	PostBookmark(bookmark *model.Bookmark) error
}
