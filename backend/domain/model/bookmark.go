package model

type Bookmark struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	UserID    uint   `json:"user_id" gorm:"index"`
	ArticleID string `json:"article_id" gorm:"type:varchar(191);index"` // VARCHARに変更
}
