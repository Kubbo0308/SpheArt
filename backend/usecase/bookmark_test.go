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

func TestBookmarkUsecase_BookmarkedArticlePerPage(t *testing.T) {
	t.Run(
		"異常系： データが取得できなかった場合",
		func(t *testing.T) {
			articles := []model.Article{}
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			bookmarkRepository := mock.NewMockBookmarkRepository(ctrl)
			bookmarkUsecase := NewBookmarkUsecase(bookmarkRepository)

			// 関数の振る舞いを定義
			bookmarkRepository.EXPECT().BookmarkedArticlesPerPages(uint(1), 1).Return(articles, errors.New("error"))

			_, err := bookmarkUsecase.BookmarkedArticlePerPage(1, 1)

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
			articles := []model.Article{
				{
					ID: "1", Title: "test", Url: "http://example.com/1", OgpImageUrl: "http://image.com/1",
					CreatedAt: time, UpdatedAt: time, PublisherId: "Pub1", PublisherName: "Publisher 1",
					PublisherImageURL: "http://pubimage.com/1", LikesCount: 100, QuoteSource: "Source 1",
				},
			}

			bookmarkRepository := mock.NewMockBookmarkRepository(ctrl)
			bookmarkUsecase := NewBookmarkUsecase(bookmarkRepository)

			// 関数の振る舞いを定義
			bookmarkRepository.EXPECT().BookmarkedArticlesPerPages(uint(1), 1).Return(articles, nil)

			res, err := bookmarkUsecase.BookmarkedArticlePerPage(uint(1), 1)

			assert.NoError(t, err)
			assert.Equal(t, len(articles), len(res))
			assert.Equal(t, articles[0], res[0])
		},
	)
}

func TestBookmarkUsecase_AllBookmarkedArticle(t *testing.T) {
	t.Run(
		"異常系： データが取得できなかった場合",
		func(t *testing.T) {
			articles := []model.Article{}
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			bookmarkRepository := mock.NewMockBookmarkRepository(ctrl)
			bookmarkUsecase := NewBookmarkUsecase(bookmarkRepository)

			// 関数の振る舞いを定義
			bookmarkRepository.EXPECT().AllBookmarkedArticleByUserId(uint(1)).Return(articles, errors.New("error"))

			_, err := bookmarkUsecase.AllBookmarkedArticle(1)

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
			articles := []model.Article{
				{
					ID: "1", Title: "test", Url: "http://example.com/1", OgpImageUrl: "http://image.com/1",
					CreatedAt: time, UpdatedAt: time, PublisherId: "Pub1", PublisherName: "Publisher 1",
					PublisherImageURL: "http://pubimage.com/1", LikesCount: 100, QuoteSource: "Source 1",
				},
			}

			bookmarkRepository := mock.NewMockBookmarkRepository(ctrl)
			bookmarkUsecase := NewBookmarkUsecase(bookmarkRepository)

			// 関数の振る舞いを定義
			bookmarkRepository.EXPECT().AllBookmarkedArticleByUserId(uint(1)).Return(articles, nil)

			res, err := bookmarkUsecase.AllBookmarkedArticle(uint(1))

			assert.NoError(t, err)
			assert.Equal(t, len(articles), len(res))
			assert.Equal(t, articles[0], res[0])
		},
	)
}

func TestBookmarkUsecase_PostBookmark(t *testing.T) {
	t.Run(
		"異常系： ブックマークの作成に失敗した場合",
		func(t *testing.T) {
			bookmark := model.Bookmark{
				UserID:    uint(1),
				ArticleID: "1",
			}

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			bookmarkRepository := mock.NewMockBookmarkRepository(ctrl)
			bookmarkUsecase := NewBookmarkUsecase(bookmarkRepository)

			// 関数の振る舞いを定義
			bookmarkRepository.EXPECT().PostBookmark(&bookmark).Return(errors.New("error"))

			res, err := bookmarkUsecase.PostBookmark(uint(1), "1")

			assert.Error(t, err)
			assert.Equal(t, model.Bookmark{}, res)
		},
	)
	t.Run(
		"正常系： ブックマークの作成が正常にできた場合",
		func(t *testing.T) {
			bookmark := model.Bookmark{
				UserID:    uint(1),
				ArticleID: "1",
			}

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			bookmarkRepository := mock.NewMockBookmarkRepository(ctrl)
			bookmarkUsecase := NewBookmarkUsecase(bookmarkRepository)

			// 関数の振る舞いを定義
			bookmarkRepository.EXPECT().PostBookmark(&bookmark).Return(nil)

			res, err := bookmarkUsecase.PostBookmark(uint(1), "1")

			assert.NoError(t, err)
			assert.Equal(t, bookmark, res)
		},
	)
}
