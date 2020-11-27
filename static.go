package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	//router.StaticFile("/a.png", "static/t.jpeg")
	router.Static("s", "static")
	router.Run(":8080")
}
