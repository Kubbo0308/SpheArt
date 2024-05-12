package persistence

import (
	"backend/domain/model"
	"backend/testutils"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestBookmarkPersistence_BookmarkedArticlePerPages(t *testing.T) {
	t.Run(
		"異常系：コネクションエラー",
		func(t *testing.T) {
			db, _, err := testutils.NewDBMock()
			assert.NoError(t, err)

			bp := NewBookmarkPersistence(db)
			_, err = bp.BookmarkedArticlesPerPages(1, 1)

			assert.Error(t, err)
		},
	)
	t.Run(
		"正常系： データが正常に取得できた場合",
		func(t *testing.T) {
			db, mock, err := testutils.NewDBMock()
			assert.NoError(t, err)

			userId := uint(1)
			pageNum := 1
			pageSize := 30

			layout := "2006-01-02 15:04:06"
			createdAtExample := "2006-01-02 15:04:06"
			time, _ := time.Parse(layout, createdAtExample)

			// 正常にデータを返すようにモックを設定
			mockArticles := []model.Article{
				{
					ID: "1", Title: "Article 1", Url: "http://example.com/1", OgpImageUrl: "http://image.com/1", CreatedAt: time, UpdatedAt: time, PublisherId: "Pub1", PublisherName: "Publisher 1", PublisherImageURL: "http://pubimage.com/1", LikesCount: 100, QuoteSource: "Source 1",
				},
			}

			rows := sqlmock.NewRows([]string{"id", "title", "url", "ogp_image_url", "created_at", "updated_at", "publisher_id", "publisher_name", "publisher_image_url", "likes_count", "quote_source"}).
				AddRow(mockArticles[0].ID, mockArticles[0].Title, mockArticles[0].Url, mockArticles[0].OgpImageUrl, mockArticles[0].CreatedAt, mockArticles[0].UpdatedAt, mockArticles[0].PublisherId, mockArticles[0].PublisherName, mockArticles[0].PublisherImageURL, mockArticles[0].LikesCount, mockArticles[0].QuoteSource)

			mock.ExpectQuery(
				regexp.QuoteMeta(
					"SELECT `articles`.`id`,`articles`.`title`,`articles`.`url`,`articles`.`ogp_image_url`,`articles`.`created_at`,`articles`.`updated_at`,`articles`.`publisher_id`,`articles`.`publisher_name`,`articles`.`publisher_image_url`,`articles`.`likes_count`,`articles`.`quote_source` FROM `articles` INNER JOIN bookmarks ON articles.id = bookmarks.article_id WHERE bookmarks.user_id = ? LIMIT ?",
				)).WithArgs(userId, pageSize).WillReturnRows(rows)

			bp := NewBookmarkPersistence(db)
			articles, err := bp.BookmarkedArticlesPerPages(userId, pageNum)

			assert.NoError(t, err)
			assert.Len(t, articles, len(mockArticles))
			assert.Equal(t, mockArticles[0].ID, articles[0].ID)
			assert.Equal(t, mockArticles[0].Title, articles[0].Title)
			assert.NoError(t, mock.ExpectationsWereMet())
		},
	)
}

func TestBookmarkPersistence_AllBookmarkedArticleByUserId(t *testing.T) {
	t.Run(
		"異常系：コネクションエラー",
		func(t *testing.T) {
			db, _, err := testutils.NewDBMock()
			assert.NoError(t, err)

			bp := NewBookmarkPersistence(db)
			_, err = bp.AllBookmarkedArticleByUserId(1)

			assert.Error(t, err)
		},
	)
	t.Run(
		"正常系： データが正常に取得できた場合",
		func(t *testing.T) {
			db, mock, err := testutils.NewDBMock()
			assert.NoError(t, err)

			userId := uint(1)

			layout := "2006-01-02 15:04:06"
			createdAtExample := "2006-01-02 15:04:06"
			time, _ := time.Parse(layout, createdAtExample)

			// 正常にデータを返すようにモックを設定
			mockArticles := []model.Article{
				{
					ID: "1", Title: "Article 1", Url: "http://example.com/1", OgpImageUrl: "http://image.com/1", CreatedAt: time, UpdatedAt: time, PublisherId: "Pub1", PublisherName: "Publisher 1", PublisherImageURL: "http://pubimage.com/1", LikesCount: 100, QuoteSource: "Source 1",
				},
			}

			rows := sqlmock.NewRows([]string{"id", "title", "url", "ogp_image_url", "created_at", "updated_at", "publisher_id", "publisher_name", "publisher_image_url", "likes_count", "quote_source"}).
				AddRow(mockArticles[0].ID, mockArticles[0].Title, mockArticles[0].Url, mockArticles[0].OgpImageUrl, mockArticles[0].CreatedAt, mockArticles[0].UpdatedAt, mockArticles[0].PublisherId, mockArticles[0].PublisherName, mockArticles[0].PublisherImageURL, mockArticles[0].LikesCount, mockArticles[0].QuoteSource)

			mock.ExpectQuery(
				regexp.QuoteMeta(
					"SELECT `articles`.`id`,`articles`.`title`,`articles`.`url`,`articles`.`ogp_image_url`,`articles`.`created_at`,`articles`.`updated_at`,`articles`.`publisher_id`,`articles`.`publisher_name`,`articles`.`publisher_image_url`,`articles`.`likes_count`,`articles`.`quote_source` FROM `articles` INNER JOIN bookmarks ON articles.id = bookmarks.article_id WHERE bookmarks.user_id = ?",
				)).WithArgs(userId).WillReturnRows(rows)

			bp := NewBookmarkPersistence(db)
			articles, err := bp.AllBookmarkedArticleByUserId(userId)

			assert.NoError(t, err)
			assert.Len(t, articles, len(mockArticles))
			assert.Equal(t, mockArticles[0].ID, articles[0].ID)
			assert.Equal(t, mockArticles[0].Title, articles[0].Title)
			assert.NoError(t, mock.ExpectationsWereMet())
		},
	)
}

func TestBookmarkPersistence_PostBookmark(t *testing.T) {
	t.Run(
		"異常系：コネクションエラー",
		func(t *testing.T) {
			db, _, err := testutils.NewDBMock()
			assert.NoError(t, err)

			bp := NewBookmarkPersistence(db)
			bookmark := model.Bookmark{}
			err = bp.PostBookmark(&bookmark)

			assert.Error(t, err)
		},
	)
	t.Run(
		"正常系： ブックマークを正常に作成できた場合",
		func(t *testing.T) {
			db, mock, err := testutils.NewDBMock()
			assert.NoError(t, err)

			bookmark := &model.Bookmark{UserID: 1, ArticleID: "2"}

			// ブックマークが存在しないことを模擬
			mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `bookmarks` WHERE user_id = ? AND article_id = ? ORDER BY `bookmarks`.`id`  LIMIT ?")).
				WithArgs(bookmark.UserID, bookmark.ArticleID, 1).
				WillReturnError(gorm.ErrRecordNotFound)

			// ブックマークの作成を期待
			mock.ExpectBegin()
			mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `bookmarks` (`user_id`,`article_id`) VALUES (?,?)")).
				WithArgs(bookmark.UserID, bookmark.ArticleID).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()

			bp := NewBookmarkPersistence(db)
			err = bp.PostBookmark(bookmark)

			assert.NoError(t, err)
			assert.NoError(t, mock.ExpectationsWereMet())
		},
	)
	t.Run(
		"正常系： ブックマークを正常に削除できた場合",
		func(t *testing.T) {
			db, mock, err := testutils.NewDBMock()
			assert.NoError(t, err)

			bookmark := &model.Bookmark{UserID: 1, ArticleID: "2"}
			rows := sqlmock.NewRows([]string{"id", "user_id", "article_id"}).AddRow(1, bookmark.UserID, bookmark.ArticleID)

			// ブックマークが存在しないことを模擬
			mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `bookmarks` WHERE user_id = ? AND article_id = ? ORDER BY `bookmarks`.`id`  LIMIT ?")).
				WithArgs(bookmark.UserID, bookmark.ArticleID, 1).
				WillReturnRows(rows)

			// ブックマークの作成を期待
			mock.ExpectBegin()
			mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `bookmarks` WHERE `bookmarks`.`id` = ?")).
				WithArgs(1).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()

			bp := NewBookmarkPersistence(db)
			err = bp.PostBookmark(bookmark)

			assert.NoError(t, err)
			assert.NoError(t, mock.ExpectationsWereMet())
		},
	)
}
