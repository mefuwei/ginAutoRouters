// Package routers is a package for routing the requests to the appropriate handler functions.
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
