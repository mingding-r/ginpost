package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ping",handler )
	r.POST("/ping2",handler2 )
	r.Run() // listen and serve on 0.0.0.0:8080
}

func handler(c *gin.Context) {
	user :=c.DefaultQuery("user","mingding")
	psw :=c.Query("psw")
	c.JSON(200, gin.H{
		"user":user,
		"psw":psw,

	})
}

func handler2(c *gin.Context) {

		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")
		/*c.JSON(200, gin.H{
			"user":user,
			"psw":psw,

		})*/
		c.String(http.StatusOK,"id: %s; page: %s; name: %s; message: %s", id, page, name, message)
		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
		//fmt.Printf("name: %s; message: %s",  name, message)
}
