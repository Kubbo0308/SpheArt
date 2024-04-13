package batch

import (
	"backend/domain/model"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dyatlov/go-opengraph/opengraph"
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

// URLからOGPのインスタンスを取得する
func GetOGPImageFromURL(url string) (string, error) {
	// http.Client の設定
	client := &http.Client{
		// 10秒のタイムアウト
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	// ライブラリのOGPインスタンス
	og := opengraph.NewOpenGraph()
	// HTMLの解析
	err = og.ProcessHTML(strings.NewReader(string(body)))
	if err != nil {
		return "", err
	}

	// OGPの画像URLを取得
	if len(og.Images) > 0 {
		return og.Images[0].URL, nil
	}
	// 画像が見つからない場合は空文字を返す
	return "", nil

}

// 型を変換する
func ConvertQiitaResponsesToArticles(qiitaResponses []model.QiitaResponse) []model.Article {
	var articles []model.Article
	for _, qiitaResp := range qiitaResponses {
		// URLからOGP画像のURL取得
		ogpImageUrl, err := GetOGPImageFromURL(qiitaResp.Url)
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
