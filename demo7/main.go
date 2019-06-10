/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-10
* Time: 上午10:29
* */
package main

import "github.com/gin-gonic/gin"

func main() {
	app := gin.Default()

	app.GET("/getString",returnString,returnJson,returnXml,returnYaml)
	app.GET("/getJson",returnJson)
	app.GET("/getXml",returnXml)
	app.GET("/getYaml",returnYaml)

	app.Run(":8085")
}

func returnString(ctx *gin.Context) {
	ctx.String(200,"return string")
}

func returnJson(ctx *gin.Context) {
	ctx.JSON(200,gin.H{
		"msg":"data",
	})
}

func returnXml(ctx *gin.Context)  {
	ctx.XML(200,gin.H{
		"msg":"data",
	})
}

func returnYaml(ctx *gin.Context) {
	ctx.YAML(200,gin.H{
		"MSG":"DATA",
	})
}
