package admin

import (
	"vvblog/api"
	"vvblog/middleware"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
	g := r.Group("/api/admin")
	{
		g.POST("/login", api.User.Login)

		userG := g.Group("/user")
		{
			userG.Use(middleware.Auth())
			userG.POST("/logout", api.User.Logout)
		}

		postG := g.Group("/posts")
		{
			// postG.Use(middleware.Auth())
			postG.POST("/upload", api.Post.Upload)
			postG.POST("", api.Post.Create)
			postG.GET("", api.Post.List)
			postG.GET("/:id", api.Post.Show)
			postG.PUT("/:id", api.Post.Update)
		}

		postCategoryG := g.Group("/post-categories")
		{
			postCategoryG.Use(middleware.Auth())
			postCategoryG.GET("/options", api.PostCategory.Options)
			postCategoryG.POST("", api.PostCategory.Create)
			postCategoryG.PUT("/:id", api.PostCategory.Update)
			postCategoryG.GET("", api.PostCategory.Index)
		}
	}
}
