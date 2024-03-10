package repository

import "backend/domain/model"

type BookmarkRepository interface {
	AllBookmarkByUserId(userId uint) ([]model.Bookmark, error)
	CreateBookmark(userId uint, articleId uint) error
}
