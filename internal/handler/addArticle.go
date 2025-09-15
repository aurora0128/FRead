package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ppeua/FRead/internal/service"
	"ppeua/FRead/model"
)

/*
	对文章的增删改查 以及获取所有文章
*/

func AddArticle(c *gin.Context) {
	//处理请求
	var req model.AddArticleReq

	//将网络参数绑定在req请求中
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "ShouldBind err: " + err.Error(),
		})
		return
	}
	//下一层处理
	article, err := service.AddArticle(req.URL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    *article,
	})
}
