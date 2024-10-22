# ginGenRouter
- Auto generate routers from gin controller comments and generate router code to file
- ginAutoRouter is a tool to generate gin router code from gin controller comments and generate router code to file.
# Go语言的Gin框架自动路由组件
#### 1. 通过解析控制器注释自动生成路由支持路由组
#### 2. 通过生成的路由代码文件方式,路由更加清晰不易出错
#### 3. Gin的路由自动注册，不用每次都手动添加
#### 4. 自动生成路由代码autoGenRouter.go 与autoGenRouter.json 两个文件

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
# Controller example Code
```go

package controllers

import "github.com/gin-gonic/gin"

type Api struct{}

// GetApis 测试API 描述
// @Summary GetApi  Summary
// @Router /api/v1/get/api/ [get]
// @Param  abc 	path  string
func (a *Api) GetApis(c *gin.Context) {

	c.JSON(200, gin.H{"message": "get api"})

}

// GetApiList param not less than 3
// @Summary GetApiList Summary
// @Router /api/v1/GetApiList/ [post,get]
// @Param   id   path   int  true  "Account ID" "primary key"
// @Param   参数  查询方式   参数类型  是否必须  "默认值" "备注"
func (a *Api) GetApiList(c *gin.Context) {

	c.JSON(200, gin.H{"message": "GetApiList"})

}


// GetUserDetail 
// @Summary GetApiList Summary
// @Router /api/v1/getUserList/:id [get]  
// @Param   id   path   int  true  "Account ID" "primary key"
func (a *Api) GetUserDetail(c *gin.Context) {

	c.JSON(200, gin.H{"message": "getUserDetail"})

}
    
```
# autoGenRouter.go , auto gen router code
```go
package routers

import (
	controllers "example/controllers"
	gin "github.com/gin-gonic/gin"
)

func RegisterRoutes(routerGroup *gin.RouterGroup) {
	api := controllers.Api{}
	// GetApi  Summary
	routerGroup.GET("/api/v1/get/api/", api.GetApis)
	// GetApiList Summary
	routerGroup.POST("/api/v1/GetApiList/", api.GetApiList)
	// GetApiList Summary
	routerGroup.GET("/api/v1/GetApiList/", api.GetApiList)
}
    
    ```