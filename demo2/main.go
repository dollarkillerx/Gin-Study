/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-9
* Time: 下午9:36
* */
package main

import "github.com/gin-gonic/gin"

func main() {
	app := gin.Default()

 	// 版本v1
	v1 := app.Group("/v1")
	{
		v1.GET("/hello",v1Fun)
	}

	v2 := app.Group("/v2")
	{
		v2.GET("/hello",v2Fun)
		admin := v2.Group("/admin")
		{
			admin.GET("/home",adminHome)
		}
	}

	app.Run(":8085")
}

func v1Fun(ctx *gin.Context)  {
	ctx.JSON(200,gin.H{
		"hello":"world",
	})
}

func v2Fun(ctx *gin.Context) {
	ctx.String(200,"Hello World")
}

func adminHome(ctx *gin.Context) {
	ctx.String(200,"this is home page")
}