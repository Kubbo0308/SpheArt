package model

import "time"

type ZennUser struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	DisplayName      string `json:"display_name"`
	AvatarSmallUrl   string `json:"avatar_small_url"`
	Pro              bool   `json:"pro"`
	AvatarRegistered bool   `json:"avatar_registered"`
}

type ZennResponse struct {
	Id                  int      `json:"id"`
	PostType            string   `json:"post_type"`
	Title               string   `json:"title"`
	Slug                string   `json:"slug"`
	CommentsCount       int      `json:"comments_count"`
	LikedCount          int      `json:"liked_count"`
	BodyLettersCount    int      `json:"body_letters_count"`
	ArticleType         string   `json:"article_type"`
	Emoji               string   `json:"emoji"`
	IsSuspendingPrivate bool     `json:"is_suspending_private"`
	PublishedAt         string   `json:"published_at"`
	BodyUpdatedAt       string   `json:"body_updated_at"`
	SourceRepoUpdatedAt *string  `json:"source_repo_updated_at"`
	Pinned              bool     `json:"pinned"`
	Path                string   `json:"path"`
	User                ZennUser `json:"user"`
}

func (zq ZennResponse) GetPublishedAt() time.Time {
	parsedTime, _ := time.Parse("2006-01-02T15:04:05", zq.PublishedAt)
	return parsedTime
}

func (zq ZennResponse) GetBodyUpdatedAt() time.Time {
	parsedTime, _ := time.Parse("2006-01-02T15:04:05", zq.BodyUpdatedAt)
	return parsedTime
}
