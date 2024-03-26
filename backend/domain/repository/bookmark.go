package repository

import "backend/domain/model"

type BookmarkRepository interface {
	BookmarkedArticlesPerPages(userId uint, pageNum int) ([]model.Article, error)
	AllBookmarkedArticleByUserId(userId uint) ([]model.Article, error)
	PostBookmark(bookmark *model.Bookmark) error
}
