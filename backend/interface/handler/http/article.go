package handler

import (
	"backend/usecase"
	"encoding/json"
	"net/http"
	"strconv"
)

type ArticleSearchHandler interface {
	SearchInArticleTitle(w http.ResponseWriter, r *http.Request)
}

type articleSearchHandler struct {
	au usecase.ArticleUsecase
}

func NewArticleSearchHandler(au usecase.ArticleUsecase) ArticleSearchHandler {
	return &articleSearchHandler{au}
}

func (ah *articleSearchHandler) SearchInArticleTitle(w http.ResponseWriter, r *http.Request) {
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
