package model

import "time"

type Article struct {
	ID                string     `json:"id" gorm:"primaryKey"`
	Title             string     `json:"title"`
	Url               string     `json:"url"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
	PublisherId       string     `json:"publisher_id"`
	PublisherName     string     `json:"publisher_name"`
	PublisherImageURL string     `json:"publisher_image_url"`
	LikesCount        int        `json:"likes_count"`
	QuoteSource       string     `json:"quote_source"`
	Bookmarks         []Bookmark `json:"foreignKey:ArticleID"`
}
