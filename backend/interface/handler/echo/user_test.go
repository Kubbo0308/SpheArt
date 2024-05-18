package handler

import (
	"backend/domain/model"
	mock "backend/testutils/mock/usecase"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestUserHandler_SignUp(t *testing.T) {
	t.Run(
		"準正常系: 不正なパラメータが渡された場合",
		func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userUsecase := mock.NewMockUserUsecase(ctrl)
			userHandler := NewUserHandler(userUsecase)

			e := echo.New()

			req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewReader([]byte(`invalid json`)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			if assert.NoError(t, userHandler.SignUp(ctx)) {
				assert.Equal(t, http.StatusBadRequest, rec.Code)
				assert.JSONEq(t, `{"error":"code=400, message=Syntax error: offset=1, error=invalid character 'i' looking for beginning of value, internal=invalid character 'i' looking for beginning of value"}`, rec.Body.String())
			}
		},
	)
	t.Run(
		"準正常系: ユーザーが既に存在する場合",
		func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userUsecase := mock.NewMockUserUsecase(ctrl)
			userHandler := NewUserHandler(userUsecase)

			e := echo.New()

			user := model.User{
				Email:    "test@example.com",
				Password: "securepassword123",
			}
			userJSON, _ := json.Marshal(user)

			req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewReader(userJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			userUsecase.EXPECT().SignUp(user).Return(nil, model.ErrUserAlreadyExists)

			if assert.NoError(t, userHandler.SignUp(ctx)) {
				assert.Equal(t, http.StatusConflict, rec.Code)
				assert.JSONEq(t, `{"error":"user already exist"}`, rec.Body.String())
			}
		},
	)
	t.Run(
		"準正常系: サーバーエラーが発生した場合",
		func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userUsecase := mock.NewMockUserUsecase(ctrl)
			userHandler := NewUserHandler(userUsecase)

			e := echo.New()

			user := model.User{
				Email:    "test@example.com",
				Password: "securepassword123",
			}
			userJSON, _ := json.Marshal(user)

			req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewReader(userJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			userUsecase.EXPECT().SignUp(user).Return(nil, assert.AnError)

			if assert.NoError(t, userHandler.SignUp(ctx)) {
				assert.Equal(t, http.StatusInternalServerError, rec.Code)
				assert.JSONEq(t, `{"error":"assert.AnError general error for testing"}`, rec.Body.String())
			}
		},
	)
	t.Run(
		"正常系: ユーザ登録が正常に行えた場合",
		func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userUsecase := mock.NewMockUserUsecase(ctrl)
			userHandler := NewUserHandler(userUsecase)

			e := echo.New()

			user := model.User{
				Email:    "test@example.com",
				Password: "securepassword123",
			}

			resUser := model.UserResponse{
				ID:    1,
				Email: "test@example.com",
			}
			userJSON, _ := json.Marshal(user)

			req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewReader(userJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			userUsecase.EXPECT().SignUp(user).Return(&resUser, nil)

			if assert.NoError(t, userHandler.SignUp(ctx)) {
				assert.Equal(t, http.StatusCreated, rec.Code)
				expectedJSON, _ := json.Marshal(resUser)
				assert.JSONEq(t, string(expectedJSON), rec.Body.String())
			}
		},
	)
}

func TestUserHandler_SignIn(t *testing.T) {
	t.Run(
		"準正常系: 不正なパラメータが渡された場合",
		func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userUsecase := mock.NewMockUserUsecase(ctrl)
			userHandler := NewUserHandler(userUsecase)

			e := echo.New()

			req := httptest.NewRequest(http.MethodPost, "/signin", bytes.NewReader([]byte(`invalid json`)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			if assert.NoError(t, userHandler.SignIn(ctx)) {
				assert.Equal(t, http.StatusBadRequest, rec.Code)
				assert.JSONEq(t, `{"error":"code=400, message=Syntax error: offset=1, error=invalid character 'i' looking for beginning of value, internal=invalid character 'i' looking for beginning of value"}`, rec.Body.String())
			}
		},
	)
	t.Run(
		"異常系: 認証が失敗した場合",
		func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userUsecase := mock.NewMockUserUsecase(ctrl)
			userHandler := NewUserHandler(userUsecase)

			e := echo.New()

			user := model.User{
				Email:    "test@example.com",
				Password: "wrongpassword",
			}
			userJSON, _ := json.Marshal(user)

			req := httptest.NewRequest(http.MethodPost, "/signin", bytes.NewReader(userJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			userUsecase.EXPECT().SignIn(user).Return("", model.ErrAuthenticationFailure)

			if assert.NoError(t, userHandler.SignIn(ctx)) {
				assert.Equal(t, http.StatusUnauthorized, rec.Code)
				assert.JSONEq(t, `{"error":"authentication failed"}`, rec.Body.String())
			}
		},
	)
	t.Run(
		"準正常系: サーバーエラーが発生した場合",
		func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userUsecase := mock.NewMockUserUsecase(ctrl)
			userHandler := NewUserHandler(userUsecase)

			e := echo.New()

			user := model.User{
				Email:    "test@example.com",
				Password: "wrongpassword",
			}
			userJSON, _ := json.Marshal(user)

			req := httptest.NewRequest(http.MethodPost, "/signin", bytes.NewReader(userJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			userUsecase.EXPECT().SignIn(user).Return("", assert.AnError)

			if assert.NoError(t, userHandler.SignIn(ctx)) {
				assert.Equal(t, http.StatusInternalServerError, rec.Code)
				assert.JSONEq(t, `{"error":"assert.AnError general error for testing"}`, rec.Body.String())
			}
		},
	)
	t.Run(
		"正常系: サインインが正常に行えた場合",
		func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userUsecase := mock.NewMockUserUsecase(ctrl)
			userHandler := NewUserHandler(userUsecase)

			e := echo.New()

			user := model.User{
				Email:    "test@example.com",
				Password: "securepassword123",
			}
			userJSON, _ := json.Marshal(user)
			tokenString := "valid.token.string"

			req := httptest.NewRequest(http.MethodPost, "/signin", bytes.NewReader(userJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			userUsecase.EXPECT().SignIn(user).Return(tokenString, nil)

			if assert.NoError(t, userHandler.SignIn(ctx)) {
				assert.Equal(t, http.StatusOK, rec.Code)
				assert.JSONEq(t, `"`+tokenString+`"`, rec.Body.String())

				cookie := rec.Result().Cookies()[0]
				assert.Equal(t, "token", cookie.Name)
				assert.Equal(t, tokenString, cookie.Value)
				assert.Equal(t, "/", cookie.Path)
				assert.Equal(t, true, cookie.HttpOnly)
				assert.Equal(t, http.SameSiteNoneMode, cookie.SameSite)
			}
		},
	)
}

func TestUserHandler_SignOut(t *testing.T) {
	t.Run(
		"正常系: サインインが正常に行えた場合",
		func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userHandler := NewUserHandler(nil)

			e := echo.New()

			req := httptest.NewRequest(http.MethodPost, "/signout", nil)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			os.Setenv("API_DOMAIN", "example.com")
			defer os.Unsetenv("API_DOMAIN")

			if assert.NoError(t, userHandler.SignOut(ctx)) {
				assert.Equal(t, http.StatusOK, rec.Code)

				cookie := rec.Result().Cookies()[0]
				assert.Equal(t, "token", cookie.Name)
				assert.Equal(t, "", cookie.Value)
				assert.WithinDuration(t, time.Now(), cookie.Expires, time.Second)
				assert.Equal(t, "/", cookie.Path)
				assert.Equal(t, "example.com", cookie.Domain)
				assert.Equal(t, true, cookie.Secure)
				assert.Equal(t, true, cookie.HttpOnly)
				assert.Equal(t, http.SameSiteNoneMode, cookie.SameSite)
			}
		},
	)
}
