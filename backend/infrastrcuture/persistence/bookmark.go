package persistence

import (
	"backend/domain/model"
	"backend/domain/repository"
	"fmt"

	"gorm.io/gorm"
)

type bookmarkPersistence struct {
	db *gorm.DB
}

func NewBookmarkPersistence(db *gorm.DB) repository.BookmarkRepository {
	return &bookmarkPersistence{db}
}

func (bp *bookmarkPersistence) AllBookmarkByUserId(userId uint) ([]model.Bookmark, error) {
	bookmarks := []model.Bookmark{}
	res := bp.db.Where("user_id = ?", userId).Find(&bookmarks)
	if res.Error != nil {
		return []model.Bookmark{}, res.Error
	}
	return bookmarks, nil
}

func (bp *bookmarkPersistence) CreateBookmark(bookmark *model.Bookmark) error {
	if err := bp.db.Create(bookmark).Error; err != nil {
		return err
	}
	return nil
}

func (bp *bookmarkPersistence) DeleteBookmark(bookmark *model.Bookmark) error {
	res := bp.db.Where("user_id = ? AND article_id = ?", bookmark.UserID, bookmark.ArticleID).Delete(&model.Bookmark{})
	if res.Error != nil {
		return res.Error
	}

	// データが見つからない場合
	if res.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
