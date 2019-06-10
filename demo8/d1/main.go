/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-10
* Time: 上午10:51
* */
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Title": "Main website",
		})
	})
	router.Run(":8080")
}
