/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-9
* Time: 下午9:46
* */
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func main() {
	app := gin.Default()
	app.Use(MiddleMain)
	//admin := app.Group("/admin")
	//admin.Use()
	app.GET("/",MiddleAuth,container,MiddleAfter)
	app.GET("/middle",before,content,after)

	app.Run(":8085")
}
func MiddleMain(ctx *gin.Context)  {
	fmt.Println("Main Middleware")
}

func before(ctx *gin.Context) {
	fmt.Println("before")
	ctx.Next()
	fmt.Println("before...........")
}

func after(ctx *gin.Context) {
	fmt.Println("after")
}

func content(ctx *gin.Context) {
	fmt.Println("content")
}


// 鉴权
func MiddleAuth(ctx *gin.Context) {
	fmt.Println("Middleware Before")
	token := ctx.GetHeader("token")
	if strings.EqualFold(token,"token") {
		ctx.Next()
		fmt.Println("Ok")
		return
	}
	ctx.JSON(401,gin.H{
		"msg":"401",
	})
	return
}

func container(ctx *gin.Context) {
	ctx.String(200,"container")
	ctx.Next()
}

func MiddleAfter(ctx *gin.Context) {
	fmt.Println("Middleware After")
}