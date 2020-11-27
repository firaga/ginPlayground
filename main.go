package main

import "github.com/gin-gonic/gin"

type User struct {
	ID   uint64
	Name string
}

func main() {
	//users := []User{{ID: 123, Name: "zhangsan"}, {ID: 345, Name: "lisi"}}
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Blog":   "www.flysnow.org",
			"wechat": "flysnow_org",
		})
	})
	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(200, "The user id is  %s", id)
	})

	r.Run(":8080")
}
