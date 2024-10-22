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
// @Router /api/v1/GetApiList/ [post]
// @Param   id   path   int  true  "Account ID" "primary key"
// @Param   参数  查询方式   参数类型  是否必须  "默认值" "备注"
func (a *Api) GetApiList(c *gin.Context) {

	c.JSON(200, gin.H{"message": "get api"})

}
