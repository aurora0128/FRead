package main

//todo:完善整个文件结构
import (
	"mime"
	"net/http"
	"ppeua/FRead/internal/config"
	"ppeua/FRead/internal/global"
	"ppeua/FRead/internal/handler"
	"ppeua/FRead/internal/repo"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConfig()
	global.InitGlobal(repo.ReadRepo)

	g := gin.Default()

	//windows bug: http.ServeFile(c.Writer, c.Request, filepath)会调用 mime.TypeByExtension(".css")
	//windows下错误影射了.css = application/x-css返回给浏览器，故不解析css
	mime.AddExtensionType(".css", "text/css")

	g.LoadHTMLFiles("./web/index.html")
	g.Static("/static", "./web/static")
	g.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	//todo：完成404页面的迁移
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 not found")
	})

	handler.RegisterRoutes(g)
	g.Run(":8080")

}
