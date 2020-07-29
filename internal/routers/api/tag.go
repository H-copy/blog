package api

import (
	"github.com/gin-gonic/gin"
)

type TagRouter struct {}

func NewTagRouter() TagRouter{
	return TagRouter{}
}

func (t TagRouter) Get(c *gin.Context){}
func (t TagRouter) List(c *gin.Context){}
func (t TagRouter) Create(c *gin.Context){}
func (t TagRouter) Update(c *gin.Context){}
func (t TagRouter) Delete(c *gin.Context){}
