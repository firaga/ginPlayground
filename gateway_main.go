package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/*filepath", Gateway("http://www.baidu.com", false))
}
