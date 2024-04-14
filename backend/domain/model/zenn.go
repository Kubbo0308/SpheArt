package model

import (
	"fmt"
	"strconv"
	"time"
)

type ZennUser struct {
	Id             int    `json:"id"`
	UserName       string `json:"username"`
	Name           string `json:"name"`
	AvatarSmallUrl string `json:"avatar_small_url"`
}

type ZennArticles struct {
	Id                  int       `json:"id"`
	PostType            string    `json:"post_type"`
	Title               string    `json:"title"`
	Slug                string    `json:"slug"`
	CommentsCount       int       `json:"comments_count"`
	LikedCount          int       `json:"liked_count"`
	BodyLettersCount    int       `json:"body_letters_count"`
	ArticleType         string    `json:"article_type"`
	Emoji               string    `json:"emoji"`
	IsSuspendingPrivate bool      `json:"is_suspending_private"`
	PublishedAt         time.Time `json:"published_at"`
	BodyUpdatedAt       time.Time `json:"body_updated_at"`
	SourceRepoUpdatedAt *string   `json:"source_repo_updated_at"`
	Pinned              bool      `json:"pinned"`
	Path                string    `json:"path"`
	User                ZennUser  `json:"user"`
}

type ZennResponse struct {
	Articles []ZennArticles `json:"articles"`
	NextPage int            `json:"next_page"`
}

func (za ZennArticles) GetUserId() string {
	return strconv.Itoa(za.User.Id)
}

func (za ZennArticles) GetUrl() string {
	return fmt.Sprintf("https://zenn.dev/%v", za.Path)
}

func (za ZennArticles) GetId() string {
	return strconv.Itoa(za.Id)
}
