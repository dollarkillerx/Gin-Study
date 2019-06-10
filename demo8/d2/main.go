/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-10
* Time: 上午10:56
* */
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	app := gin.Default()

	app.LoadHTMLGlob("temp/**/*")

	app.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK,"posts/index.html",gin.H{
			"Title":"Jell",
		})
	})

	app.Run(":8085")
}
