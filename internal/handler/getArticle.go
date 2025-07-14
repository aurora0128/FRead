package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ppeua/FRead/internal/service"
)

func GetArticle(c *gin.Context) {
	//解析id
	id := c.Param("id")
	article, err := service.GetArticle(id)
	if err != nil {
		//todo:详细处理文章没有找到的情况
		if err.Error() == "article not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"error":   err.Error(),
			})
			return
		}
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
	return
}
