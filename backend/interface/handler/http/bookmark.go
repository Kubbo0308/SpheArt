package handler

import (
	"backend/usecase"
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
)

type BookmarkHandler interface {
	BookmarkPerPage(w http.ResponseWriter, r *http.Request)
	AllBookmark(w http.ResponseWriter, r *http.Request)
	PostBookmark(w http.ResponseWriter, r *http.Request)
}

type bookmarkHandler struct {
	bu usecase.BookmarkUsecase
}

func NewBookmarkHandler(bu usecase.BookmarkUsecase) BookmarkHandler {
	return &bookmarkHandler{bu}
}

func (bh *bookmarkHandler) BookmarkPerPage(w http.ResponseWriter, r *http.Request) {
	user, err := jwt.Parse(r.Header.Get("Authorization"), func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		http.Error(w, "Unauthorized!!!", http.StatusUnauthorized)
		return
	}
	claims, ok := user.Claims.(jwt.MapClaims)
	if !ok || !user.Valid {
		http.Error(w, "Unauthorized!", http.StatusUnauthorized)
		return
	}
	userId := uint(claims["user_id"].(float64))

	queryParam := r.URL.Query().Get("per_page")
	perPage, err := strconv.Atoi(queryParam)
	if err != nil {
		http.Error(w, "Invalid per_page parameter", http.StatusBadRequest)
		return
	}

	res, err := bh.bu.BookmarkedArticlePerPage(userId, perPage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (bh *bookmarkHandler) AllBookmark(w http.ResponseWriter, r *http.Request) {
	user, err := jwt.Parse(r.Header.Get("Authorization"), func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	claims, ok := user.Claims.(jwt.MapClaims)
	if !ok || !user.Valid {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	userId := uint(claims["user_id"].(float64))

	res, err := bh.bu.AllBookmarkedArticle(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
	}
}

func (bh *bookmarkHandler) PostBookmark(w http.ResponseWriter, r *http.Request) {
	user, err := jwt.Parse(r.Header.Get("Authorization"), func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	claims, ok := user.Claims.(jwt.MapClaims)
	if !ok || !user.Valid {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	userId := uint(claims["user_id"].(float64))

	articleId := r.URL.Query().Get("articleId")

	res, err := bh.bu.PostBookmark(userId, articleId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
	}
}
