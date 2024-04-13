package batch

import (
	"backend/domain/model"
	"backend/utils"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"gorm.io/gorm"
)

func RunQiitaAPIBatch(db *gorm.DB) {
	// Qiita APIからデータを取得
	var qiitaResp []model.QiitaResponse
	err := GetQiitaArticleFromAPI(&qiitaResp)
	if err != nil {
		log.Println("Failed to fetch Qiita items:", err)
	}

	articles := ConvertQiitaResponsesToArticles(qiitaResp)

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

// Qiita APIから記事を取得する
func GetQiitaArticleFromAPI(jsonData *[]model.QiitaResponse) error {
	res, err := http.Get(`https://qiita.com/api/v2/items?page=1&per_page=100`)
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

// 型を変換する
func ConvertQiitaResponsesToArticles(qiitaResponses []model.QiitaResponse) []model.Article {
	var articles []model.Article
	for _, qiitaResp := range qiitaResponses {
		// URLからOGP画像のURL取得
		ogpImageUrl, err := utils.GetOGPImageFromURL(qiitaResp.Url)
		if err != nil {
			ogpImageUrl = ""
		}

		articles = append(articles, model.Article{
			ID:                qiitaResp.Id,
			Title:             qiitaResp.Title,
			Url:               qiitaResp.Url,
			OgpImageUrl:       ogpImageUrl,
			CreatedAt:         qiitaResp.CreatedAt,
			UpdatedAt:         qiitaResp.UpdatedAt,
			PublisherId:       qiitaResp.User.UserId,
			PublisherName:     qiitaResp.User.Name,
			PublisherImageURL: qiitaResp.User.ProfileImageUrl,
			LikesCount:        qiitaResp.LikesCount,
			QuoteSource:       "qiita",
		})
	}
	return articles
}
