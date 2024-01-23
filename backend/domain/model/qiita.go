package model

import "time"

type QiitaTag struct {
	Name     string   `json:"name"`
	Versions []string `json:"versions"`
}

type User struct {
	Description       string  `json:"description"`
	FacebookId        string  `json:"facebook_id"`
	FolloweesCount    int     `json:"followees_count"`
	FollowersCount    int     `json:"followers_count"`
	GithubLoginName   *string `json:"github_login_name"`
	UserId            string  `json:"id"`
	ItemsCount        int     `json:"items_count"`
	LinkedinId        string  `json:"linkedin_id"`
	Location          string  `json:"location"`
	Name              string  `json:"name"`
	Organization      string  `json:"organization"`
	PermanentId       int     `json:"permanent_id"`
	ProfileImageUrl   string  `json:"profile_image_url"`
	TeamOnly          bool    `json:"team_only"`
	TwitterScreenName *string `json:"twitter_screen_name"`
	WebsiteUrl        string  `json:"website_url"`
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
	Tags           QiitaTag  `json:"tags"`
	Title          string    `json:"title"`
	UpdatedAt      time.Time `json:"updated_at"`
	Url            string    `json:"url"`
	User           User      `json:"user"`
}
