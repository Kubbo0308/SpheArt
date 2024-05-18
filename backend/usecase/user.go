package usecase

import (
	"backend/domain/model"
	"backend/domain/repository"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserUsecase interface {
	SignUp(user model.User) (*model.UserResponse, error)
	SignIn(user model.User) (string, error)
}

type userUsecase struct {
	ur repository.UserRepository
}

func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	return &userUsecase{ur}
}

func (uu *userUsecase) SignUp(user model.User) (*model.UserResponse, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return nil, err
	}
	// 既に同じメールアドレスを持つユーザーが存在するかチェック
	existingUser := model.User{}
	if err := uu.ur.UserByEmail(&existingUser, user.Email); err == nil {
		return nil, model.ErrUserAlreadyExists
	}
	newUser := model.User{Email: user.Email, Password: string(hash)}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return nil, err
	}
	resUser := model.UserResponse{
		ID:    newUser.ID,
		Email: newUser.Email,
	}
	return &resUser, nil
}

func (uu *userUsecase) SignIn(user model.User) (string, error) {
	checkUser := model.User{}
	if err := uu.ur.UserByEmail(&checkUser, user.Email); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// ユーザーが見つからない場合の一般的な認証失敗エラー
			return "", model.ErrAuthenticationFailure
		}
		// その他のデータベースエラーなど
		return "", err
	}
	err := bcrypt.CompareHashAndPassword([]byte(checkUser.Password), []byte(user.Password))
	if err != nil {
		// パスワードが間違っている場合も認証失敗エラー
		return "", model.ErrAuthenticationFailure
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": checkUser.ID,
		"exp":     time.Now().Add(time.Hour * 3).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		// JWT生成失敗の抽象化したエラーメッセージ
		return "", errors.New("failed to generate token")
	}
	return tokenString, err
}
