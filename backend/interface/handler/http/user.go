package handler

import (
	"backend/domain/model"
	"backend/usecase"
	"encoding/json"
	"net/http"
	"os"
	"time"
)

type UserHandler interface {
	SignUp(w http.ResponseWriter, r *http.Request)
	SignIn(w http.ResponseWriter, r *http.Request)
	SignOut(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	uu usecase.UserUsecase
}

func NewUserHandler(uu usecase.UserUsecase) UserHandler {
	return &userHandler{uu}
}

func decodeUser(r *http.Request) (model.User, error) {
	var user model.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	return user, err
}

func (uh *userHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	user, err := decodeUser(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := uh.uu.SignUp(user)
	if err != nil {
		if err == model.ErrUserAlreadyExists {
			http.Error(w, err.Error(), http.StatusConflict)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
	}
}

func (uh *userHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	user, err := decodeUser(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tokenString, err := uh.uu.SignIn(user)
	if err != nil {
		if err == model.ErrAuthenticationFailure {
			http.Error(w, err.Error(), http.StatusUnauthorized)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	cookie := &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  time.Now().Add(3 * time.Hour),
		Path:     "/",
		Domain:   os.Getenv("API_DOMAIN"),
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	}
	http.SetCookie(w, cookie)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tokenString)
	if err != nil {
		http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
	}
}

func (uh *userHandler) SignOut(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now(),
		Path:     "/",
		Domain:   os.Getenv("API_DOMAIN"),
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	}
	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
}
