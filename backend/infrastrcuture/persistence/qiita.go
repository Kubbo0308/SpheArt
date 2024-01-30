package persistence

import (
	"backend/domain/model"
	"backend/domain/repository"
	"encoding/json"
	"io"
	"net/http"

	"gorm.io/gorm"
)

type qiitaPersistence struct {
	db *gorm.DB
}

func NewQiitaArticlePersistence(db *gorm.DB) repository.QiitaRepository {
	return &qiitaPersistence{db}
}

func (qp *qiitaPersistence) GetAllQiitaArticles() ([]model.Article, error) {
	var qiitaResp []model.QiitaResponse
	err := GetQiitaArticleFromAPI(&qiitaResp)
	if err != nil {
		return []model.Article{}, err
	}
	articles := ConvertQiitaResponsesToArticles(qiitaResp)
	return articles, nil
}

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

func ConvertQiitaResponsesToArticles(qiitaResponses []model.QiitaResponse) []model.Article {
	var articles []model.Article
	for _, qiitaResp := range qiitaResponses {
		articles = append(articles, model.Article{
			ID:                qiitaResp.GetId(),
			Title:             qiitaResp.Title,
			Url:               qiitaResp.Url,
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
