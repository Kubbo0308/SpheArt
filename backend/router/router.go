package router

import (
	"backend/interface/handler"

	"github.com/labstack/echo/v4"
)

func NewRouter(ah handler.ArticleHandler, uh handler.UserHandler) *echo.Echo {
	e := echo.New()

	e.GET("/articles", ah.GetAllArticles)
	e.GET("/articles/search", ah.SearchInArticleTitle)
	e.POST("/signup", uh.SignUp)
	e.POST("/signin", uh.SignIn)
	e.POST("/signout", uh.SignOut)
	return e
}
