package persistence

import (
	"backend/domain/model"
	"backend/testutils"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestUserPersistence_CreateUser(t *testing.T) {
	t.Run(
		"異常系：コネクションエラー",
		func(t *testing.T) {
			db, _, err := testutils.NewDBMock()
			assert.NoError(t, err)

			up := NewUserPersistence(db)
			user := model.User{}
			err = up.CreateUser(&user)

			assert.Error(t, err)
		},
	)
	t.Run(
		"正常系： データ登録が正常に行えた場合",
		func(t *testing.T) {
			db, mock, err := testutils.NewDBMock()
			assert.NoError(t, err)

			layout := "2006-01-02 15:04:06"
			createdAtExample := "2006-01-02 15:04:06"
			time, _ := time.Parse(layout, createdAtExample)

			user := model.User{
				Email:     "test@example.com",
				Password:  "securepassword123",
				CreatedAt: time,
				UpdatedAt: time,
			}

			mock.ExpectBegin()
			mock.ExpectExec(
				regexp.QuoteMeta(
					"INSERT INTO `users` (`email`,`password`,`created_at`,`updated_at`) VALUES (?,?,?,?)",
				)).WithArgs(user.Email, user.Password, time, time).WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()

			up := NewUserPersistence(db)
			err = up.CreateUser(&user)

			assert.NoError(t, err)
			assert.NoError(t, mock.ExpectationsWereMet())
		},
	)
}

func TestUserPersistence_UserByEmail(t *testing.T) {
	t.Run(
		"異常系：コネクションエラー",
		func(t *testing.T) {
			db, _, err := testutils.NewDBMock()
			assert.NoError(t, err)

			email := "test@example.com"

			up := NewUserPersistence(db)
			user := model.User{}
			err = up.UserByEmail(&user, email)

			assert.Error(t, err)
		},
	)
	t.Run(
		"正常系：ユーザーを正常に取得できた場合",
		func(t *testing.T) {
			db, mock, err := testutils.NewDBMock()
			assert.NoError(t, err)

			layout := "2006-01-02 15:04:06"
			createdAtExample := "2006-01-02 15:04:06"
			time, _ := time.Parse(layout, createdAtExample)

			email := "test@example.com"
			mockUser := model.User{
				ID:        1,
				Email:     email,
				Password:  "securepassword123",
				CreatedAt: time,
				UpdatedAt: time,
			}

			// データベースからの期待されるクエリと戻り値の設定
			rows := sqlmock.NewRows([]string{"id", "email", "password", "created_at", "updated_at"}).
				AddRow(mockUser.ID, mockUser.Email, mockUser.Password, mockUser.CreatedAt, mockUser.UpdatedAt)

			mock.ExpectQuery(
				regexp.QuoteMeta("SELECT * FROM `users` WHERE email=? ORDER BY `users`.`id` LIMIT ?"),
			).WithArgs(email, 1).WillReturnRows(rows)

			up := NewUserPersistence(db)
			user := &model.User{}
			err = up.UserByEmail(user, email)

			assert.NoError(t, err)
			assert.Equal(t, mockUser.ID, user.ID)
			assert.Equal(t, mockUser.Email, user.Email)
			assert.NoError(t, mock.ExpectationsWereMet())
		},
	)
}
