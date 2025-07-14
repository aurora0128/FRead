package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ppeua/FRead/internal/service"
	"ppeua/FRead/model"
)

func GetArticles(c *gin.Context) {
	//没有request,直接走获取数据库
	articles, err := service.GetArticles()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "GetArticles: " + err.Error(),
		})
		return
	}
	//数据转换

	resp := model.GetArticleListRes{
		Success: true,
		Data:    *articles,
		Pagination: struct {
			Page  int `json:"page"`
			Limit int `json:"limit"`
			Total int `json:"total"`
			Pages int `json:"pages"`
		}{
			Page:  1,
			Limit: 20,
			Total: 50,
			Pages: 3,
		},
	}
	c.JSON(http.StatusOK, resp)
}
