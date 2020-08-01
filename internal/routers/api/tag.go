package api

import (	
	
	"blog/global"
	"blog/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type TagRouter struct {}

func NewTagRouter() TagRouter{
	return TagRouter{}
}

func (t TagRouter) Get(c *gin.Context){}
func (t TagRouter) List(c *gin.Context){

	statusCode := errcode.Success.StatusCode()
	msg := errcode.Success.Msg()
	
	c.JSON(
		statusCode,
		gin.H{
			"serverConf": &global.ServerSetting, 
			"appConf": &global.AppSetting, 
			"databaseConf": &global.DatabaseSetting, 
			"msg": msg,
		},
	)
	
}
func (t TagRouter) Create(c *gin.Context){}
func (t TagRouter) Update(c *gin.Context){}
func (t TagRouter) Delete(c *gin.Context){}
