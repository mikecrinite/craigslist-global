package main

import (
	"net/http"

	"github.com/mikecrinite/craigslist-go/controller"
	"github.com/mikecrinite/craigslist-go/model"

	"github.com/gin-gonic/gin"
)

// Run the application
func main() {
	r := gin.Default()
	r.Static("/css", "./web/css")
	r.LoadHTMLGlob("web/templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":      "craigslist-global",
			"categories": model.CategoryMapKeys(),
		})
	})
	r.POST("/", func(c *gin.Context) {
		query := c.PostForm("search")
		category := c.PostForm("category")
		posts := controller.ScrapeCL(model.CategoryMap[category], controller.CleanForQuery(query))

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":      "craigslist-global",
			"categories": model.CategoryMapKeys(),
			"links":      posts,
			"selected":   category,
		})
	})
	r.Run(":8095") // 0.0.0.0:8095
}
