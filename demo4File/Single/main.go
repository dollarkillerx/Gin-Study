/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-10
* Time: 上午9:08
* */
package Single

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	app := gin.Default()

	app.POST("/upload", func(context *gin.Context) {
		app.MaxMultipartMemory = 100 << 20 // 设置最大上传大小为100M
		header, _ := context.FormFile("file")

		path := "./file/" + header.Filename // 上传存储到的地址
		fmt.Println(path)
		err := context.SaveUploadedFile(header, path)
		if err != nil {
			fmt.Println(err.Error())
		}

		log.Println(header.Filename," ",header.Size," ",header.Header)
		context.JSON(200,gin.H{
			"fileName":header.Filename,
		})
	})

	app.Run(":8085")
}
