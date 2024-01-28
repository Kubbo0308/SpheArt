package persistence

import (
	"backend/domain/model"
	"backend/domain/repository"
	"encoding/json"
	"io"
	"net/http"

	"gorm.io/gorm"
)

type zennPersistence struct {
	db *gorm.DB
}

func NewZennArticlePersistence(db *gorm.DB) repository.ZennRepository {
	return &zennPersistence{db}
}

func (zp *zennPersistence) GetAllZennArticles() ([]model.Article, error) {
	var zennResp model.ZennResponse
	err := GetZennArticleFromAPI(&zennResp)
	if err != nil {
		return []model.Article{}, err
	}
	articles := ConvertZennResponsesToArticles(zennResp.Articles)
	return articles, nil
}

func GetZennArticleFromAPI(jsonData *model.ZennResponse) error {
	res, err := http.Get(`https://zenn.dev/api/articles?page=1&per_page=100`)
	if err != nil {
		return err
	}
	// リクエストを読み込む。
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	// 必ず閉じる。
	defer res.Body.Close()
	// リクエストを引数に受け取った構造体にマッピングする
	err = json.Unmarshal(body, jsonData)
	if err != nil {
		return err
	}
	return nil
}

func ConvertZennResponsesToArticles(zennResponses []model.ZennArticles) []model.Article {
	var articles []model.Article
	for _, zennResp := range zennResponses {
		articles = append(articles, model.Article{
			ID:                uint(zennResp.Id),
			Title:             zennResp.Title,
			Url:               zennResp.GetUrl(),
			CreatedAt:         zennResp.PublishedAt,
			UpdatedAt:         zennResp.BodyUpdatedAt,
			PublisherId:       zennResp.GetUserId(),
			PublisherName:     zennResp.User.Name,
			PublisherImageURL: zennResp.User.AvatarSmallUrl,
			Likes_count:       zennResp.LikedCount,
		})
	}
	return articles
}
