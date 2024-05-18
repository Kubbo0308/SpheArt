package handler

import (
	"backend/domain/model"
	"backend/usecase"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	SignUp(ctx echo.Context) error
	SignIn(ctx echo.Context) error
	SignOut(ctx echo.Context) error
}

type userHandler struct {
	uu usecase.UserUsecase
}

func NewUserHandler(uu usecase.UserUsecase) UserHandler {
	return &userHandler{uu}
}

func (uh *userHandler) SignUp(ctx echo.Context) error {
	user := model.User{}
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	res, err := uh.uu.SignUp(user)
	if err != nil {
		// 既に存在するユーザーである場合のエラー
		if err == model.ErrUserAlreadyExists { // ErrUserAlreadyExists はユーザーが既に存在するときのエラー
			return ctx.JSON(http.StatusConflict, map[string]string{"error": err.Error()})
		}
		// その他のエラー
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusCreated, res)
}

func (uh *userHandler) SignIn(ctx echo.Context) error {
	user := model.User{}
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	tokenString, err := uh.uu.SignIn(user)
	if err != nil {
		// 認証失敗エラー
		if err == model.ErrAuthenticationFailure {
			return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
		}
		// その他の内部エラー
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(3 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	ctx.SetCookie(cookie)
	return ctx.JSON(http.StatusOK, tokenString)
}

func (uh *userHandler) SignOut(ctx echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	ctx.SetCookie(cookie)
	return ctx.NoContent(http.StatusOK)
}
