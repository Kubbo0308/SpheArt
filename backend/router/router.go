package router

import (
	"backend/interface/handler"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(ah handler.ArticleHandler, uh handler.UserHandler, bh handler.BookmarkHandler) *echo.Echo {
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
	b := e.Group("/bookmark")
	// bookmarkエンドポイントにミドルウェア追加
	b.Use(echojwt.WithConfig(echojwt.Config{
		// jwtを生成した時と同じSECRET_KEYを指定
		SigningKey: []byte(os.Getenv("SECRET")),
		// Clientから送られてくるjwtトークンの置き場所を指定
		TokenLookup: "cookie:token",
	}))
	b.GET("/", bh.AllBookmark)
	b.POST("/:articleId", bh.PostBookmark)
	b.DELETE("/:articleId", bh.DeleteBookmark)
	return e
}
