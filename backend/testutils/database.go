package testutils

import (
	"backend/domain/utils"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDBMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		return nil, mock, err
	}

	db, err := gorm.Open(mysql.New(
		mysql.Config{
			Conn:                      mockDB,
			SkipInitializeWithVersion: true,
		}),
		&gorm.Config{
			// テスト時にutils.SetNow()で現在時刻を設定するためモックDBのNowFuncを上書き
			NowFunc: func() time.Time {
				return utils.Now()
			},
		},
	)

	return db, mock, err
}
