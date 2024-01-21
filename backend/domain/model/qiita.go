package model

import "time"

type Tag struct {
	Name     string   `json:"name"`
	Versions []string `json:"versions"`
}

type User struct {
	Description     string `json:"description"`
	FolloweesCount  int    `json:"followees_count"`
	FollowersCount  int    `json:"followers_count"`
	UserId          string `json:"id"`
	Name            string `json:"name"`
	ProfileImageUrl string `json:"profile_image_url"`
}

type QiitaResponse struct {
	RenderdBody    string    `json:"renderd_body"`
	Body           string    `json:"body"`
	Coediting      bool      `json:"coediting"`
	CommentsCount  int       `json:"comments_count"`
	CreatedAt      time.Time `json:"created_at"`
	Id             string    `json:"id"`
	LikesCount     int       `json:"likes_count"`
	Private        bool      `json:"private"`
	ReactionsCount int       `json:"reactions_count"`
	StocksCount    int       `json:"stocks_count"`
	Tags           Tag       `json:"tags"`
	Title          string    `json:"title"`
	UpdatedAt      time.Time `json:"updated_at"`
	Url            string    `json:"url"`
	User           User      `json:"user"`
}
