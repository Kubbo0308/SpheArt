package persistence

import (
	"backend/domain/model"
	"backend/domain/repository"
	"encoding/json"
	"io"
	"net/http"

	"gorm.io/gorm"
)

type articlePersistence struct {
	database *gorm.DB
}

func NewArticlePersistence(db *gorm.DB) repository.ArticleRepository {
	return &articlePersistence{database: db}
}

func (ap *articlePersistence) GetAllQiitaArticles(articles *[]model.QiitaResponse) error {
	err := getQiitaArticleFromAPI(*articles)
	if err != nil {
		return err
	}
	return nil
}

func getQiitaArticleFromAPI(jsonData []model.QiitaResponse) error {
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
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		return err
	}
	return nil
}
