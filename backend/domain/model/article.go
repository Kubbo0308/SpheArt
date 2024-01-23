package model

import "time"

type Article struct {
	ID                uint      `json:"id" gorm:"primaryKey"`
	Title             string    `json:"title"`
	Url               string    `json:"url"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	PublisherId       string    `json:"publisher_id"`
	PublisherName     string    `json:"publisher_name"`
	PublisherImageURL string    `json:"publisher_image_url"`
	Likes_count       int       `json:"likes_count"`
}
