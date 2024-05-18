package usecase

import (
	"backend/domain/model"
	mock "backend/testutils/mock/domain/repository"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func TestUserUsecase_SignUp_Failure(t *testing.T) {
	layout := "2006-01-02 15:04:06"
	createdAtExample := "2006-01-02 15:04:06"
	time, _ := time.Parse(layout, createdAtExample)

	user := model.User{
		Email:     "test@example.com",
		Password:  "securepassword123",
		CreatedAt: time,
		UpdatedAt: time,
	}

	tests := []struct {
		testName string
		input    model.User
		mockFn   func(
			patches *gomonkey.Patches,
			ur *mock.MockUserRepository,
		)
	}{
		{
			testName: "異常系： パスワードの生成に失敗した場合",
			input:    model.User{},
			mockFn: func(
				patches *gomonkey.Patches,
				ur *mock.MockUserRepository,
			) {
				patches.ApplyFunc(bcrypt.GenerateFromPassword, func(password []byte, cost int) ([]byte, error) {
					return nil, errors.New("error")
				})
			},
		},
		{
			testName: "準正常系： ユーザが既に存在している場合",
			input:    user,
			mockFn: func(
				patches *gomonkey.Patches,
				ur *mock.MockUserRepository,
			) {
				patches.ApplyFunc(bcrypt.GenerateFromPassword, func(password []byte, cost int) ([]byte, error) {
					return []byte{}, nil
				})
				ur.EXPECT().UserByEmail(gomock.AssignableToTypeOf(&model.User{}), "test@example.com").DoAndReturn(
					func(u *model.User, email string) error {
						*u = user
						return nil
					},
				).Times(1)
			},
		},
		{
			testName: "異常系： ユーザの作成に失敗した場合",
			input:    user,
			mockFn: func(
				patches *gomonkey.Patches,
				ur *mock.MockUserRepository,
			) {
				patches.ApplyFunc(bcrypt.GenerateFromPassword, func(password []byte, cost int) ([]byte, error) {
					return []byte{}, nil
				})
				ur.EXPECT().UserByEmail(gomock.AssignableToTypeOf(&model.User{}), "test@example.com").Return(errors.New("error")).Times(1)
				ur.EXPECT().CreateUser(gomock.Any()).Return(errors.New("error")).Times(1)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			userRepository := mock.NewMockUserRepository(ctrl)
			patches := gomonkey.NewPatches()
			defer patches.Reset()

			tt.mockFn(
				patches,
				userRepository,
			)
			userUsecase := NewUserUsecase(userRepository)
			res, err := userUsecase.SignUp(tt.input)
			assert.Error(t, err)
			assert.Nil(t, res)
		})
	}
}

func TestUserUsecase_SignUp(t *testing.T) {
	t.Run(
		"正常系: ユーザ登録が正常に行えた場合",
		func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUser := model.User{}

			userRepository := mock.NewMockUserRepository(ctrl)
			userUsecase := NewUserUsecase(userRepository)

			patches := gomonkey.NewPatches()
			defer patches.Reset()

			patches.ApplyFunc(bcrypt.GenerateFromPassword, func(password []byte, cost int) ([]byte, error) {
				return []byte{}, nil
			})

			// 関数の振る舞いを定義
			userRepository.EXPECT().UserByEmail(&mockUser, "").Return(errors.New("error"))
			userRepository.EXPECT().CreateUser(&mockUser).Return(nil)

			_, err := userUsecase.SignUp(mockUser)

			assert.NoError(t, err)
		},
	)
}

func TestUserUsecase_SignIn_Failure(t *testing.T) {
	layout := "2006-01-02 15:04:06"
	createdAtExample := "2006-01-02 15:04:06"
	time, _ := time.Parse(layout, createdAtExample)

	user := model.User{
		Email:     "test@example.com",
		Password:  "securepassword123",
		CreatedAt: time,
		UpdatedAt: time,
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	checkUser := model.User{
		Email:     user.Email,
		Password:  string(hashedPassword),
		CreatedAt: time,
		UpdatedAt: time,
	}

	tests := []struct {
		testName string
		input    model.User
		mockFn   func(
			patches *gomonkey.Patches,
			ur *mock.MockUserRepository,
		)
	}{
		{
			testName: "異常系：ユーザーが見つからない場合",
			input:    user,
			mockFn: func(
				patches *gomonkey.Patches,
				ur *mock.MockUserRepository,
			) {
				ur.EXPECT().UserByEmail(gomock.AssignableToTypeOf(&model.User{}), user.Email).Return(gorm.ErrRecordNotFound)
			},
		},
		{
			testName: "異常系：その他のデータベースエラー",
			input:    user,
			mockFn: func(
				patches *gomonkey.Patches,
				ur *mock.MockUserRepository,
			) {
				ur.EXPECT().UserByEmail(gomock.AssignableToTypeOf(&model.User{}), user.Email).Return(errors.New("database error"))
			},
		},
		{
			testName: "異常系：パスワードが間違っている場合",
			input:    user,
			mockFn: func(
				patches *gomonkey.Patches,
				ur *mock.MockUserRepository,
			) {
				ur.EXPECT().UserByEmail(gomock.AssignableToTypeOf(&model.User{}), user.Email).DoAndReturn(
					func(u *model.User, email string) error {
						*u = checkUser
						return nil
					},
				)
				patches.ApplyFunc(bcrypt.CompareHashAndPassword, func(hashedPassword, password []byte) error {
					return errors.New("password mismatch")
				})
			},
		},
		{
			testName: "異常系：トークン生成に失敗した場合",
			input:    user,
			mockFn: func(
				patches *gomonkey.Patches,
				ur *mock.MockUserRepository,
			) {
				ur.EXPECT().UserByEmail(gomock.AssignableToTypeOf(&model.User{}), user.Email).DoAndReturn(
					func(u *model.User, email string) error {
						*u = checkUser
						return nil
					},
				)
				patches.ApplyFunc(jwt.NewWithClaims, func(method jwt.SigningMethod, claims jwt.Claims) *jwt.Token {
					return jwt.NewWithClaims(method, claims)
				})
				patches.ApplyFunc((*jwt.Token).SignedString, func(token *jwt.Token, key interface{}) (string, error) {
					return "", errors.New("failed to generate token")
				})
			},
		},
	}

	os.Setenv("SECRET", "test_secret")

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			userRepository := mock.NewMockUserRepository(ctrl)
			patches := gomonkey.NewPatches()
			defer patches.Reset()
			defer os.Unsetenv("SECRET")

			tt.mockFn(
				patches,
				userRepository,
			)
			userUsecase := NewUserUsecase(userRepository)
			res, err := userUsecase.SignIn(tt.input)
			assert.Error(t, err)
			assert.Equal(t, res, "")
		})
	}
}

func TestUserUsecase_SignIn(t *testing.T) {
	t.Run(
		"正常系: ユーザ認証が正常に行えた場合",
		func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			layout := "2006-01-02 15:04:06"
			createdAtExample := "2006-01-02 15:04:06"
			time, _ := time.Parse(layout, createdAtExample)

			user := model.User{
				Email:     "test@example.com",
				Password:  "securepassword123",
				CreatedAt: time,
				UpdatedAt: time,
			}

			hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
			checkUser := model.User{
				Email:     user.Email,
				Password:  string(hashedPassword),
				CreatedAt: time,
				UpdatedAt: time,
			}

			expectedToken := "valid.token.string"

			userRepository := mock.NewMockUserRepository(ctrl)
			userUsecase := NewUserUsecase(userRepository)

			patches := gomonkey.NewPatches()
			defer patches.Reset()

			os.Setenv("SECRET", "test_secret")
			defer os.Unsetenv("SECRET")

			userRepository.EXPECT().UserByEmail(gomock.AssignableToTypeOf(&model.User{}), user.Email).DoAndReturn(
				func(u *model.User, email string) error {
					*u = checkUser
					return nil
				},
			)
			patches.ApplyFunc(jwt.NewWithClaims, func(method jwt.SigningMethod, claims jwt.Claims) *jwt.Token {
				return jwt.NewWithClaims(method, claims)
			})
			patches.ApplyFunc((*jwt.Token).SignedString, func(token *jwt.Token, key interface{}) (string, error) {
				return "valid.token.string", nil
			})

			res, err := userUsecase.SignIn(user)

			assert.NoError(t, err)
			assert.Equal(t, expectedToken, res)
		},
	)
}
