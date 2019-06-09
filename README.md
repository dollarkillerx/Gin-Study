# Gin-Study
Gin Go Web Framework Study
文档地址:https://www.kancloud.cn/adapa/gingolang

### 课程环境:
- Ubuntu18.4
- Go1.12.5
- Ide: Goland
- vgo

### 老规矩先来一个HelloWorld
`vgo get -u github.com/gin-gonic/gin`
``` 
package main

import "github.com/gin-gonic/gin"

func main() {
	app := gin.Default()
	
	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{
			"message":"Hello World",
		})
	})
	
	app.Run(":8085")
}
```