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
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":      "craigslist-global",
			"categories": model.CategoryMapKeys(),
		})
	})
	r.POST("/", func(c *gin.Context) {
		query := c.PostForm("search")
		category := c.PostForm("category")

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":      "craigslist-global",
			"categories": model.CategoryMapKeys(),
			"links":      controller.ScrapeCL(model.CategoryMap[category], controller.CleanForQuery(query)),
			"selected":   category,
		})
	})
	r.Run(":8095") // 0.0.0.0:8095
}
