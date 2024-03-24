package repository

import "backend/domain/model"

type BookmarkRepository interface {
	AllBookmarkedArticleByUserId(userId uint) ([]model.Article, error)
	PostBookmark(bookmark *model.Bookmark) error
}
