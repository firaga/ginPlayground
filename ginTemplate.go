package main

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"html/template"
)

func main() {
	r := gin.Default()

	r.SetFuncMap(template.FuncMap{
		"md5": MD5,
	})
	r.LoadHTMLFiles("html/index.html")
	r.GET("/html", func(c *gin.Context) {
		c.HTML(200, "index.html", "flysnow_org")
	})
	r.Run(":8080")
}
func MD5(in string) (string, error) {
	hash := md5.Sum([]byte(in))
	return hex.EncodeToString(hash[:]), nil
}
