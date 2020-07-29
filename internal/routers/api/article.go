package api

import (
	"github.com/gin-gonic/gin"
)

type ArticleRouter struct {}

func NewArticleRouter() ArticleRouter{
	return ArticleRouter{}
}

func (t ArticleRouter) Get(c *gin.Context){}
func (t ArticleRouter) List(c *gin.Context){}
func (t ArticleRouter) Create(c *gin.Context){}
func (t ArticleRouter) Update(c *gin.Context){}
func (t ArticleRouter) Delete(c *gin.Context){}
