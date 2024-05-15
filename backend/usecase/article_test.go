package usecase

import (
	"backend/domain/model"
	mock "backend/testutils/mock/domain/repository"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestArticleUsecase_ArticlesPerPage(t *testing.T) {
	t.Run(
		"異常系： データが取得できなかった場合",
		func(t *testing.T) {
			articles := []model.Article{}
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			articleRepository := mock.NewMockArticleRepository(ctrl)
			articleUsecase := NewArticleUsecase(articleRepository)

			err := errors.New("error")
			// 関数の振る舞いを定義
			articleRepository.EXPECT().ArticlesPerPages(1).Return(articles, err)

			_, err = articleUsecase.ArticlesPerPage(1)

			assert.Error(t, err)
		},
	)
	t.Run(
		"正常系： データが取得できた場合",
		func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			layout := "2006-01-02 15:04:06"
			createdAtExample := "2006-01-02 15:04:06"
			time, _ := time.Parse(layout, createdAtExample)

			// 正常にデータを返すようにモックを設定
			mockArticles := []model.Article{
				{
					ID: "1", Title: "test", Url: "http://example.com/1", OgpImageUrl: "http://image.com/1",
					CreatedAt: time, UpdatedAt: time, PublisherId: "Pub1", PublisherName: "Publisher 1",
					PublisherImageURL: "http://pubimage.com/1", LikesCount: 100, QuoteSource: "Source 1",
				},
			}

			articleRepository := mock.NewMockArticleRepository(ctrl)
			articleUsecase := NewArticleUsecase(articleRepository)

			// 関数の振る舞いを定義
			articleRepository.EXPECT().ArticlesPerPages(1).Return(mockArticles, nil)

			res, err := articleUsecase.ArticlesPerPage(1)

			assert.NoError(t, err)
			assert.Equal(t, len(mockArticles), len(res))
			assert.Equal(t, mockArticles[0], res[0])
		},
	)
}

func TestArticleUsecase_AllArticles(t *testing.T) {
	t.Run(
		"異常系： データが取得できなかった場合",
		func(t *testing.T) {
			articles := []model.Article{}
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			articleRepository := mock.NewMockArticleRepository(ctrl)
			articleUsecase := NewArticleUsecase(articleRepository)

			err := errors.New("error")
			// 関数の振る舞いを定義
			articleRepository.EXPECT().AllArticles().Return(articles, err)

			_, err = articleUsecase.AllArticles()

			assert.Error(t, err)
		},
	)
	t.Run(
		"正常系： データが取得できた場合",
		func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			layout := "2006-01-02 15:04:06"
			createdAtExample := "2006-01-02 15:04:06"
			time, _ := time.Parse(layout, createdAtExample)

			// 正常にデータを返すようにモックを設定
			mockArticles := []model.Article{
				{
					ID: "1", Title: "test", Url: "http://example.com/1", OgpImageUrl: "http://image.com/1",
					CreatedAt: time, UpdatedAt: time, PublisherId: "Pub1", PublisherName: "Publisher 1",
					PublisherImageURL: "http://pubimage.com/1", LikesCount: 100, QuoteSource: "Source 1",
				},
			}

			articleRepository := mock.NewMockArticleRepository(ctrl)
			articleUsecase := NewArticleUsecase(articleRepository)

			// 関数の振る舞いを定義
			articleRepository.EXPECT().AllArticles().Return(mockArticles, nil)

			res, err := articleUsecase.AllArticles()

			assert.NoError(t, err)
			assert.Equal(t, len(mockArticles), len(res))
			assert.Equal(t, mockArticles[0], res[0])
		},
	)
}

func TestArticleUsecase_SearchInArticleTitle(t *testing.T) {
	t.Run(
		"異常系： データが取得できなかった場合",
		func(t *testing.T) {
			articles := []model.Article{}
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			articleRepository := mock.NewMockArticleRepository(ctrl)
			articleUsecase := NewArticleUsecase(articleRepository)

			err := errors.New("error")
			// 関数の振る舞いを定義
			articleRepository.EXPECT().SearchInArticleTitle("test", 1).Return(articles, err)

			_, err = articleUsecase.SearchInArticleTitle("test", 1)

			assert.Error(t, err)
		},
	)
	t.Run(
		"正常系： データが取得できた場合",
		func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			layout := "2006-01-02 15:04:06"
			createdAtExample := "2006-01-02 15:04:06"
			time, _ := time.Parse(layout, createdAtExample)

			// 正常にデータを返すようにモックを設定
			mockArticles := []model.Article{
				{
					ID: "1", Title: "test", Url: "http://example.com/1", OgpImageUrl: "http://image.com/1",
					CreatedAt: time, UpdatedAt: time, PublisherId: "Pub1", PublisherName: "Publisher 1",
					PublisherImageURL: "http://pubimage.com/1", LikesCount: 100, QuoteSource: "Source 1",
				},
			}

			articleRepository := mock.NewMockArticleRepository(ctrl)
			articleUsecase := NewArticleUsecase(articleRepository)

			// 関数の振る舞いを定義
			articleRepository.EXPECT().SearchInArticleTitle("test", 1).Return(mockArticles, nil)

			res, err := articleUsecase.SearchInArticleTitle("test", 1)

			assert.NoError(t, err)
			assert.Equal(t, len(mockArticles), len(res))
			assert.Equal(t, mockArticles[0], res[0])
		},
	)
}
