package batch

import (
	"backend/domain/model"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"gorm.io/gorm"
)

func RunZennAPIBatch(db *gorm.DB) {
	// Zenn APIからデータを取得
	var zennResp model.ZennResponse
	err := GetZennArticleFromAPI(&zennResp)
	if err != nil {
		log.Println("Failed to fetch Qiita items:", err)
	}
	articles := ConvertZennResponsesToArticles(zennResp.Articles)

	// データベースに保存
	for _, item := range articles {
		var existArticle model.Article

		if err := db.Where("id = ?", item.ID).First(&existArticle).Error; err != nil {
			// データベースに存在していないデータのみ保存
			if err == gorm.ErrRecordNotFound {
				if err = db.Create(&item).Error; err != nil {
					fmt.Println("Failed to create article:", err)
				}
			} else {
				fmt.Println("Error checking for existing article:", err)
			}
		} else {
			fmt.Println("Article with ID", item.ID, "already exists")
		}
	}
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
			ID:                zennResp.GetId(),
			Title:             zennResp.Title,
			Url:               zennResp.GetUrl(),
			CreatedAt:         zennResp.PublishedAt,
			UpdatedAt:         zennResp.BodyUpdatedAt,
			PublisherId:       zennResp.GetUserId(),
			PublisherName:     zennResp.User.Name,
			PublisherImageURL: zennResp.User.AvatarSmallUrl,
			LikesCount:        zennResp.LikedCount,
			QuoteSource:       "zenn",
		})
	}
	return articles
}
