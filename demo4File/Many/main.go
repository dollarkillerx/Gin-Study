/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-10
* Time: 上午9:45
* */
package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	app := gin.Default()

	app.POST("/uploads", func(ctx *gin.Context) {
		form, _ := ctx.MultipartForm()
		files := form.File["upload[]"]
		for _,file := range files {
			log.Println(file.Filename)
			ctx.SaveUploadedFile(file,"./file/"+file.Filename) // 文件夹需要向创建
		}
	})

	app.Run(":8085")
}
