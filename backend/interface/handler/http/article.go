package handler

import (
	"backend/usecase"
	"encoding/json"
	"net/http"
	"strconv"
)

type ArticleHandler interface {
	ArticlesPerPage(w http.ResponseWriter, r *http.Request)
	AllArticles(w http.ResponseWriter, r *http.Request)
	SearchInArticleTitle(w http.ResponseWriter, r *http.Request)
}

type articleHandler struct {
	au usecase.ArticleUsecase
}

func NewArticleHandler(au usecase.ArticleUsecase) ArticleHandler {
	return &articleHandler{au}
}

func (ah *articleHandler) ArticlesPerPage(w http.ResponseWriter, r *http.Request) {
	queryParam := r.URL.Query().Get("per_page")
	perPage, err := strconv.Atoi(queryParam)
	if err != nil {
		http.Error(w, "Invalid per_page parameter", http.StatusBadRequest)
		return
	}

	res, err := ah.au.ArticlesPerPage(perPage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (ah *articleHandler) AllArticles(w http.ResponseWriter, r *http.Request) {
	res, err := ah.au.AllArticles()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (ah *articleHandler) SearchInArticleTitle(w http.ResponseWriter, r *http.Request) {
	queryParamTitle := r.URL.Query().Get("title")
	searchTitle := "%" + queryParamTitle + "%"

	queryParamPerPage := r.URL.Query().Get("per_page")
	perPage, err := strconv.Atoi(queryParamPerPage)
	if err != nil {
		http.Error(w, "Invalid per_page parameter", http.StatusBadRequest)
		return
	}

	res, err := ah.au.SearchInArticleTitle(searchTitle, perPage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
