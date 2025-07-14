package handler

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/articles", AddArticle)
		api.GET("/articles", GetArticles)
		api.GET("/articles/:id", GetArticle)
		api.DELETE("/articles/:id", DeleteArticle)
	}
}
