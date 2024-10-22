# ginGenRouter
Auto generate routers from gin controller comments and generate router code to file
# Example
```go
	
    // new a gin app
    app := gin.New()
    autoRouter := r.Group("/")
	// params:
    // controllersPath: the path of the controllers package directory
    // routersPath: the path of the routers package directory
    generate.GenRouters("", "")
    autoRouter := app.Group("/")
	// register routers
    //RegisterRoutes(autoRouter)
    gin_app.Run(":8080")

```
