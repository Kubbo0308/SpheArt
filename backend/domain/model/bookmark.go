package model

type Bookmark struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	UserID    uint   `json:"user_id"`
	ArticleID string `json:"article_id"`
}
