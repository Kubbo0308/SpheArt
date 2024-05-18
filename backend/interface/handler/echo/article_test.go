package handler

import (
	"backend/domain/model"
	mock "backend/testutils/mock/usecase"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestArticleHandler_ArticlesPerPage(t *testing.T) {
	t.Run(
		"準正常系： 不正なパラメータが渡された場合",
		func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			articleUsecase := mock.NewMockArticleUsecase(ctrl)
			articleHandler := NewArticleHandler(articleUsecase)

			echo := echo.New()

			req := httptest.NewRequest(http.MethodGet, "/articles?per_page=invalid", nil)
			rec := httptest.NewRecorder()
			ctx := echo.NewContext(req, rec)

			if assert.NoError(t, articleHandler.ArticlesPerPage(ctx)) {
				assert.Equal(t, http.StatusBadRequest, rec.Code)
				assert.JSONEq(t, `{"error":"Invalid per_page parameter"}`, rec.Body.String())
			}
		},
	)
	t.Run(
		"準正常系： サーバーエラーが発生した場合",
		func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			articleUsecase := mock.NewMockArticleUsecase(ctrl)
			articleHandler := NewArticleHandler(articleUsecase)

			echo := echo.New()

			req := httptest.NewRequest(http.MethodGet, "/articles?per_page=10", nil)
			rec := httptest.NewRecorder()
			ctx := echo.NewContext(req, rec)

			articleUsecase.EXPECT().ArticlesPerPage(10).Return(nil, assert.AnError)

			if assert.NoError(t, articleHandler.ArticlesPerPage(ctx)) {
				assert.Equal(t, http.StatusInternalServerError, rec.Code)
				assert.JSONEq(t, `{"error":"assert.AnError general error for testing"}`, rec.Body.String())
			}
		},
	)
	t.Run(
		"正常系: 正常にレスポンスが取得できた場合",
		func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			articleUsecase := mock.NewMockArticleUsecase(ctrl)
			articleHandler := NewArticleHandler(articleUsecase)

			echo := echo.New()

			req := httptest.NewRequest(http.MethodGet, "/articles?per_page=10", nil)
			rec := httptest.NewRecorder()
			ctx := echo.NewContext(req, rec)

			layout := time.RFC3339
			createdAtExample := "2006-01-02T15:04:06Z"
			time, _ := time.Parse(layout, createdAtExample)

			expectedArticles := []model.Article{
				{
					ID:                "1",
					Title:             "Article 1",
					Url:               "https://example.com/article1",
					OgpImageUrl:       "https://example.com/image1.jpg",
					CreatedAt:         time,
					UpdatedAt:         time,
					PublisherId:       "publisher1",
					PublisherName:     "Publisher 1",
					PublisherImageURL: "https://example.com/publisher1.jpg",
					LikesCount:        100,
					QuoteSource:       "Source 1",
					Bookmarks:         nil,
				},
				{
					ID:                "2",
					Title:             "Article 2",
					Url:               "https://example.com/article2",
					OgpImageUrl:       "https://example.com/image2.jpg",
					CreatedAt:         time,
					UpdatedAt:         time,
					PublisherId:       "publisher2",
					PublisherName:     "Publisher 2",
					PublisherImageURL: "https://example.com/publisher2.jpg",
					LikesCount:        200,
					QuoteSource:       "Source 2",
					Bookmarks:         nil,
				},
			}

			expectedJSON := `[{
				"id": "1",
				"title": "Article 1",
				"url": "https://example.com/article1",
				"ogp_image_url": "https://example.com/image1.jpg",
				"created_at": "2006-01-02T15:04:06Z",
				"updated_at": "2006-01-02T15:04:06Z",
				"publisher_id": "publisher1",
				"publisher_name": "Publisher 1",
				"publisher_image_url": "https://example.com/publisher1.jpg",
				"likes_count": 100,
				"quote_source": "Source 1",
				"foreignKey:ArticleID": null
			}, {
				"id": "2",
				"title": "Article 2",
				"url": "https://example.com/article2",
				"ogp_image_url": "https://example.com/image2.jpg",
				"created_at": "2006-01-02T15:04:06Z",
				"updated_at": "2006-01-02T15:04:06Z",
				"publisher_id": "publisher2",
				"publisher_name": "Publisher 2",
				"publisher_image_url": "https://example.com/publisher2.jpg",
				"likes_count": 200,
				"quote_source": "Source 2",
				"foreignKey:ArticleID": null
			}]`

			articleUsecase.EXPECT().ArticlesPerPage(10).Return(expectedArticles, nil)

			if assert.NoError(t, articleHandler.ArticlesPerPage(ctx)) {
				assert.Equal(t, http.StatusOK, rec.Code)
				assert.JSONEq(t, expectedJSON, rec.Body.String())
			}
		},
	)
}

func TestArticleHandler_AllArticles(t *testing.T) {
	t.Run(
		"準正常系： サーバーエラーが発生した場合",
		func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			articleUsecase := mock.NewMockArticleUsecase(ctrl)
			articleHandler := NewArticleHandler(articleUsecase)

			echo := echo.New()

			req := httptest.NewRequest(http.MethodGet, "/articles", nil)
			rec := httptest.NewRecorder()
			ctx := echo.NewContext(req, rec)

			articleUsecase.EXPECT().AllArticles().Return(nil, assert.AnError)

			if assert.NoError(t, articleHandler.AllArticles(ctx)) {
				assert.Equal(t, http.StatusInternalServerError, rec.Code)
				assert.JSONEq(t, `{"error":"assert.AnError general error for testing"}`, rec.Body.String())
			}
		},
	)
	t.Run(
		"正常系: 正常にレスポンスが取得できた場合",
		func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			articleUsecase := mock.NewMockArticleUsecase(ctrl)
			articleHandler := NewArticleHandler(articleUsecase)

			echo := echo.New()

			req := httptest.NewRequest(http.MethodGet, "/articles", nil)
			rec := httptest.NewRecorder()
			ctx := echo.NewContext(req, rec)

			layout := time.RFC3339
			createdAtExample := "2006-01-02T15:04:06Z"
			time, _ := time.Parse(layout, createdAtExample)

			expectedArticles := []model.Article{
				{
					ID:                "1",
					Title:             "Article 1",
					Url:               "https://example.com/article1",
					OgpImageUrl:       "https://example.com/image1.jpg",
					CreatedAt:         time,
					UpdatedAt:         time,
					PublisherId:       "publisher1",
					PublisherName:     "Publisher 1",
					PublisherImageURL: "https://example.com/publisher1.jpg",
					LikesCount:        100,
					QuoteSource:       "Source 1",
					Bookmarks:         nil,
				},
				{
					ID:                "2",
					Title:             "Article 2",
					Url:               "https://example.com/article2",
					OgpImageUrl:       "https://example.com/image2.jpg",
					CreatedAt:         time,
					UpdatedAt:         time,
					PublisherId:       "publisher2",
					PublisherName:     "Publisher 2",
					PublisherImageURL: "https://example.com/publisher2.jpg",
					LikesCount:        200,
					QuoteSource:       "Source 2",
					Bookmarks:         nil,
				},
			}

			expectedJSON := `[{
				"id": "1",
				"title": "Article 1",
				"url": "https://example.com/article1",
				"ogp_image_url": "https://example.com/image1.jpg",
				"created_at": "2006-01-02T15:04:06Z",
				"updated_at": "2006-01-02T15:04:06Z",
				"publisher_id": "publisher1",
				"publisher_name": "Publisher 1",
				"publisher_image_url": "https://example.com/publisher1.jpg",
				"likes_count": 100,
				"quote_source": "Source 1",
				"foreignKey:ArticleID": null
			}, {
				"id": "2",
				"title": "Article 2",
				"url": "https://example.com/article2",
				"ogp_image_url": "https://example.com/image2.jpg",
				"created_at": "2006-01-02T15:04:06Z",
				"updated_at": "2006-01-02T15:04:06Z",
				"publisher_id": "publisher2",
				"publisher_name": "Publisher 2",
				"publisher_image_url": "https://example.com/publisher2.jpg",
				"likes_count": 200,
				"quote_source": "Source 2",
				"foreignKey:ArticleID": null
			}]`

			articleUsecase.EXPECT().AllArticles().Return(expectedArticles, nil)

			if assert.NoError(t, articleHandler.AllArticles(ctx)) {
				assert.Equal(t, http.StatusOK, rec.Code)
				assert.JSONEq(t, expectedJSON, rec.Body.String())
			}
		},
	)
}

func TestArticleHandler_SearchInArticleTitle(t *testing.T) {
	t.Run(
		"準正常系： 不正なパラメータが渡された場合",
		func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			articleUsecase := mock.NewMockArticleUsecase(ctrl)
			articleHandler := NewArticleHandler(articleUsecase)

			echo := echo.New()

			req := httptest.NewRequest(http.MethodGet, "/articles/search?title=example&per_page=invalid", nil)
			rec := httptest.NewRecorder()
			ctx := echo.NewContext(req, rec)

			if assert.NoError(t, articleHandler.SearchInArticleTitle(ctx)) {
				assert.Equal(t, http.StatusBadRequest, rec.Code)
				assert.JSONEq(t, `{"error":"Invalid per_page parameter"}`, rec.Body.String())
			}
		},
	)
	t.Run(
		"準正常系： サーバーエラーが発生した場合",
		func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			articleUsecase := mock.NewMockArticleUsecase(ctrl)
			articleHandler := NewArticleHandler(articleUsecase)

			echo := echo.New()

			req := httptest.NewRequest(http.MethodGet, "/articles/search?title=example&per_page=10", nil)
			rec := httptest.NewRecorder()
			ctx := echo.NewContext(req, rec)

			articleUsecase.EXPECT().SearchInArticleTitle("%example%", 10).Return(nil, assert.AnError)

			if assert.NoError(t, articleHandler.SearchInArticleTitle(ctx)) {
				assert.Equal(t, http.StatusInternalServerError, rec.Code)
				assert.JSONEq(t, `{"error":"assert.AnError general error for testing"}`, rec.Body.String())
			}
		},
	)
	t.Run(
		"正常系: 正常にレスポンスが取得できた場合",
		func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			articleUsecase := mock.NewMockArticleUsecase(ctrl)
			articleHandler := NewArticleHandler(articleUsecase)

			echo := echo.New()

			req := httptest.NewRequest(http.MethodGet, "/articles/search?title=example&per_page=10", nil)
			rec := httptest.NewRecorder()
			ctx := echo.NewContext(req, rec)

			layout := time.RFC3339
			createdAtExample := "2006-01-02T15:04:06Z"
			time, _ := time.Parse(layout, createdAtExample)

			expectedArticles := []model.Article{
				{
					ID:                "1",
					Title:             "Article 1",
					Url:               "https://example.com/article1",
					OgpImageUrl:       "https://example.com/image1.jpg",
					CreatedAt:         time,
					UpdatedAt:         time,
					PublisherId:       "publisher1",
					PublisherName:     "Publisher 1",
					PublisherImageURL: "https://example.com/publisher1.jpg",
					LikesCount:        100,
					QuoteSource:       "Source 1",
					Bookmarks:         nil,
				},
				{
					ID:                "2",
					Title:             "Article 2",
					Url:               "https://example.com/article2",
					OgpImageUrl:       "https://example.com/image2.jpg",
					CreatedAt:         time,
					UpdatedAt:         time,
					PublisherId:       "publisher2",
					PublisherName:     "Publisher 2",
					PublisherImageURL: "https://example.com/publisher2.jpg",
					LikesCount:        200,
					QuoteSource:       "Source 2",
					Bookmarks:         nil,
				},
			}

			expectedJSON := `[{
				"id": "1",
				"title": "Article 1",
				"url": "https://example.com/article1",
				"ogp_image_url": "https://example.com/image1.jpg",
				"created_at": "2006-01-02T15:04:06Z",
				"updated_at": "2006-01-02T15:04:06Z",
				"publisher_id": "publisher1",
				"publisher_name": "Publisher 1",
				"publisher_image_url": "https://example.com/publisher1.jpg",
				"likes_count": 100,
				"quote_source": "Source 1",
				"foreignKey:ArticleID": null
			}, {
				"id": "2",
				"title": "Article 2",
				"url": "https://example.com/article2",
				"ogp_image_url": "https://example.com/image2.jpg",
				"created_at": "2006-01-02T15:04:06Z",
				"updated_at": "2006-01-02T15:04:06Z",
				"publisher_id": "publisher2",
				"publisher_name": "Publisher 2",
				"publisher_image_url": "https://example.com/publisher2.jpg",
				"likes_count": 200,
				"quote_source": "Source 2",
				"foreignKey:ArticleID": null
			}]`

			articleUsecase.EXPECT().SearchInArticleTitle("%example%", 10).Return(expectedArticles, nil)

			if assert.NoError(t, articleHandler.SearchInArticleTitle(ctx)) {
				assert.Equal(t, http.StatusOK, rec.Code)
				assert.JSONEq(t, expectedJSON, rec.Body.String())
			}
		},
	)
}
