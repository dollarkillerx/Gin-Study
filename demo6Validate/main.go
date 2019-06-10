/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-10
* Time: 上午10:21
* */
package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Login struct {
	User string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	app := gin.Default()

	app.POST("/login", func(ctx *gin.Context) {
		login := &Login{}
		if err := ctx.ShouldBind(login);err != nil {
			log.Println(err.Error())
			ctx.String(400,"you input error")
			return
		}
		ctx.JSON(200,login)
	})

	app.Run(":8085")
}
