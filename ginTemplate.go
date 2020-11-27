package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("html/index.html")
	r.GET("/html", func(c *gin.Context) {
		c.HTML(200, "index.html", "flysnow_org")
	})
	r.Run(":8080")
}
