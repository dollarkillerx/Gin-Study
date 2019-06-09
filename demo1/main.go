/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-9
* Time: 下午4:16
* */
package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func main() {
	app := gin.New()
	
	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{
			"message":"Hello World",
		})
	})

	app.GET("/pps",pps) // /pps?name=?&password=?
	app.GET("/name/:name/:age",hello)
	app.POST("/pps",ppos)
	app.POST("/json",jsson)
	app.POST("/jsonBan",jsBan)
	app.GET("/test",test)

	app.Run(":8085")
}

func test(ctx *gin.Context) {
	ctx.String(200,"Test")
}

func jsBan(ctx *gin.Context) {
	type user struct {
		User string `json:"user"`
		Password string `json:"password"`
	}
	u := &user{}
	err := ctx.Bind(u)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(u.User)
	fmt.Println(u.Password)
	ctx.JSON(http.StatusOK,u)
}

func jsson(ctx *gin.Context) {
	body := ctx.Request.Body
	defer body.Close()
	type User struct {
		User string `json:"user"`
		Password string `json:password`
	}

	bytes, _ := ioutil.ReadAll(body)
	user := &User{}
	json.Unmarshal(bytes,user)
	ctx.JSON(200,user)
}

func ppos(ctx *gin.Context) {
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")
	ctx.JSON(200,gin.H{
		"name":name,
		"password":password,
	})
}

func pps(ctx *gin.Context) {
	name := ctx.Query("name")
	pwd := ctx.Query("password")

	ctx.JSON(200,gin.H{
		"name":name,
		"pwd":pwd,
	})
}

func hello(ctx *gin.Context) {
	name := ctx.Param("name")
	age := ctx.Param("age")
	ctx.JSON(200,gin.H{
		"name":name,
		"age":age,
	})
}