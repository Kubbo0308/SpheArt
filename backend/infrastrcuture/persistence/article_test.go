package persistence

import (
	"backend/testutils"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestArticlePersistence_ArticlesPerPages(t *testing.T) {
	t.Run(
		"異常系： コネクションエラー",
		func(t *testing.T) {
			db, _, err := testutils.NewDBMock()
			assert.NoError(t, err)

			ap := NewArticlePersistence(db)
			_, error := ap.ArticlesPerPages(1)

			assert.Error(t, error)
		},
	)
	t.Run(
		"正常系： データが正常に取得できた場合",
		func(t *testing.T) {
			db, mock, err := testutils.NewDBMock()
			assert.NoError(t, err)

			// モックのレスポンスで返すカラムとデータを設定
			rows := sqlmock.NewRows([]string{
				"id", "title", "url", "ogp_image_url", "created_at", "updated_at",
				"publisher_id", "publisher_name", "publisher_image_url", "likes_count", "quote_source",
			}).AddRow(
				"1", "test", "http://example.com", "http://example.com/image.jpg",
				time.Now(), time.Now(), "123", "Publisher Name", "http://example.com/pub_img.jpg",
				10, "Sample Source",
			)

			mock.ExpectQuery(
				regexp.QuoteMeta(
					"SELECT * FROM `articles` LIMIT ?",
				)).WithArgs(20).WillReturnRows(rows)

			ap := NewArticlePersistence(db)
			res, error := ap.ArticlesPerPages(1)

			assert.NoError(t, error)
			assert.Len(t, res, 1)
			assert.Equal(t, res[0].ID, "1")
			assert.Equal(t, res[0].Title, "test")
			assert.NoError(t, mock.ExpectationsWereMet())
		},
	)
}

func TestArticlePersistence_AllArticles(t *testing.T) {
	t.Run(
		"異常系： コネクションエラー",
		func(t *testing.T) {
			db, _, err := testutils.NewDBMock()
			assert.NoError(t, err)

			ap := NewArticlePersistence(db)
			_, error := ap.AllArticles()

			assert.Error(t, error)
		},
	)
	t.Run(
		"正常系： データが正常に取得できた場合",
		func(t *testing.T) {
			db, mock, err := testutils.NewDBMock()
			assert.NoError(t, err)

			// モックのレスポンスで返すカラムとデータを設定
			rows := sqlmock.NewRows([]string{
				"id", "title", "url", "ogp_image_url", "created_at", "updated_at",
				"publisher_id", "publisher_name", "publisher_image_url", "likes_count", "quote_source",
			}).AddRow(
				"1", "test", "http://example.com", "http://example.com/image.jpg",
				time.Now(), time.Now(), "123", "Publisher Name", "http://example.com/pub_img.jpg",
				10, "Sample Source",
			)

			mock.ExpectQuery(
				regexp.QuoteMeta(
					"SELECT * FROM `articles`",
				)).WithArgs().WillReturnRows(rows)

			ap := NewArticlePersistence(db)
			res, error := ap.AllArticles()

			assert.NoError(t, error)
			assert.Len(t, res, 1)
			assert.Equal(t, res[0].ID, "1")
			assert.Equal(t, res[0].Title, "test")
			assert.NoError(t, mock.ExpectationsWereMet())
		},
	)
}

func TestArticlePersistence_SearchInArticleTitle(t *testing.T) {
	t.Run(
		"異常系： コネクションエラー",
		func(t *testing.T) {
			db, _, err := testutils.NewDBMock()
			assert.NoError(t, err)

			ap := NewArticlePersistence(db)
			_, error := ap.SearchInArticleTitle("", 1)

			assert.Error(t, error)
		},
	)
	t.Run(
		"正常系： データが正常に取得できた場合",
		func(t *testing.T) {
			db, mock, err := testutils.NewDBMock()
			assert.NoError(t, err)

			// モックのレスポンスで返すカラムとデータを設定
			rows := sqlmock.NewRows([]string{
				"id", "title", "url", "ogp_image_url", "created_at", "updated_at",
				"publisher_id", "publisher_name", "publisher_image_url", "likes_count", "quote_source",
			}).AddRow(
				"1", "test", "http://example.com", "http://example.com/image.jpg",
				time.Now(), time.Now(), "123", "Publisher Name", "http://example.com/pub_img.jpg",
				10, "Sample Source",
			)

			mock.ExpectQuery(
				regexp.QuoteMeta(
					"SELECT * FROM `articles` WHERE title LIKE ? LIMIT ?",
				)).WithArgs("test", 20).WillReturnRows(rows)

			ap := NewArticlePersistence(db)
			res, error := ap.SearchInArticleTitle("test", 1)

			assert.NoError(t, error)
			assert.Len(t, res, 1)
			assert.Equal(t, res[0].ID, "1")
			assert.Equal(t, res[0].Title, "test")
			assert.NoError(t, mock.ExpectationsWereMet())
		},
	)
}
