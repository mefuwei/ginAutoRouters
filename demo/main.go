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
	autoRouter := r.Group("/")
	routers.RegisterRoutes(autoRouter)
	r.Run() // listen and serve on 0.0.0.0:8080
}
