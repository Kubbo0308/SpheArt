package persistence

import (
	"backend/domain/model"
	"backend/domain/repository"
	"encoding/json"
	"fmt"
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
	var zennResp []model.ZennResponse
	err := GetZennArticleFromAPI(&zennResp)
	fmt.Println(zennResp)
	if err != nil {
		return []model.Article{}, err
	}
	articles := ConvertZennResponsesToArticles(zennResp)
	return articles, nil
}

func GetZennArticleFromAPI(jsonData *[]model.ZennResponse) error {
	res, err := http.Get(`https://zenn.dev/api/articles`)
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
	fmt.Println(jsonData)
	if err != nil {
		return err
	}
	return nil
}

func ConvertZennResponsesToArticles(zennResponses []model.ZennResponse) []model.Article {
	var articles []model.Article
	for _, zennResp := range zennResponses {
		articles = append(articles, model.Article{
			Title:             zennResp.Title,
			Url:               zennResp.Path,
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
