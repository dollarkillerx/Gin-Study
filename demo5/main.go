/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-10
* Time: 上午9:56
* */
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type User struct {
	User string `form:"user"`
	Password string `form:"password"`
}

func main() {
	app := gin.Default()

	app.GET("/test",getlog)
	app.POST("/logo",logo)


	app.Run(":8085")
}
func logo(ctx *gin.Context) {
	user := &User{}
	ctx.ShouldBind(user)
	ctx.JSON(200,user)
}

// 注意不推荐用git传输秘密这样是非常不安全的!
func getlog(ctx *gin.Context)  {
	var user User
	if ctx.ShouldBindQuery(&user) == nil {
		log.Println(user.User)
		fmt.Println(ctx.Query("user"))
		log.Println(user.Password)
	}
	ctx.JSON(200,user)
}
