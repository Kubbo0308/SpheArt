package router

import (
	"backend/interface/handler"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(ah handler.ArticleHandler, uh handler.UserHandler) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")},
		AllowHeaders: []string{
			echo.HeaderOrigin, echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders,
			echo.HeaderXCSRFToken,
		},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))

	e.GET("/articles", ah.GetAllArticles)
	e.GET("/articles/search", ah.SearchInArticleTitle)
	e.POST("/signup", uh.SignUp)
	e.POST("/signin", uh.SignIn)
	e.POST("/signout", uh.SignOut)
	return e
}
