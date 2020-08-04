package api

import (	
	
	"blog/global"
	"blog/pkg/errcode"
	"blog/internal/model"

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
	
	tags := []model.Tag{}

	global.DBEngine.Find(&tags)

	print(tags)
	
	c.JSON(
		statusCode,
		gin.H{
			"serverConf": &global.ServerSetting, 
			"msg": msg,
			"dataList": tags,
		},
	)
	
}
func (t TagRouter) Create(c *gin.Context){}
func (t TagRouter) Update(c *gin.Context){}
func (t TagRouter) Delete(c *gin.Context){}
