package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ppeua/FRead/internal/service"
)

func DeleteArticle(c *gin.Context) {
	//解析id
	id := c.Param("id")
	err := service.DeleteArticle(id)
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
		"message": "文章删除成功",
	})
	return
}
