# ginGenRouter
- Auto generate routers from gin controller comments and generate router code to file
- ginAutoRouter is a tool to generate gin router code from gin controller comments and generate router code to file.
# Go语言的Gin框架自动路由组件
#### 1. 通过解析控制器注释自动生成路由
#### 2. 通过生成的路由代码文件方式,路由更加清晰不易出错
#### 3. Gin的路由自动注册，不用每次都手动添加
# Example
```go

package main

import (
	"demo/routers"

	"github.com/gin-gonic/gin"
	"github.com/mefuwei/ginGenRouter"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	ginGenRouter.GenRouters("", "")
	// 创建路由组并注册路由,首次需要注销掉
	//autoRouter := r.Group("/")
	// routers.RegisterRoutes(autoRouter)
	//
	r.Run() // listen and serve on 0.0.0.0:8080
}


```
