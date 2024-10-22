package main

import (
	generate "autorouters"
	"autorouters/example/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	autoRouter := r.Group("/")

	generate.GenRouters("", "")
	routers.RegisterRoutes(autoRouter)
	r.Run() // listen and serve on 0.0.0.0:8080
}
