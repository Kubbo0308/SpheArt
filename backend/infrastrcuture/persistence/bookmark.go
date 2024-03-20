package persistence

import (
	"backend/domain/model"
	"backend/domain/repository"
	"errors"

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

func (bp *bookmarkPersistence) PostBookmark(bookmark *model.Bookmark) error {
	// ブックマークがすでに存在するかを確認
	existingBookmark := model.Bookmark{}
	if err := bp.db.Where("user_id = ? AND article_id = ?", bookmark.UserID, bookmark.ArticleID).First(&existingBookmark).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			// 予期しないエラーが発生した場合
			return err
		}
		// レコードが存在しない場合はブックマークを作成
		if err := bp.db.Create(bookmark).Error; err != nil {
			return err
		}
	} else {
		// レコードが存在する場合はブックマークを削除
		if err := bp.db.Delete(&existingBookmark).Error; err != nil {
			return err
		}
	}
	return nil
}
