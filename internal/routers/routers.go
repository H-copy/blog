package routers

import (
	"github.com/gin-gonic/gin"
	"blog/internal/routers/api"
)

func NewRouter() *gin.Engine{
	r := gin.New()
	
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	
	Tag := api.NewTagRouter()
	Article := api.NewArticleRouter()

	apiV1 := r.Group("/v1")
	{
		apiV1.POST("/tags", Tag.Create)
		apiV1.DELETE("/tags", Tag.Delete)
		apiV1.PUT("/tags/:id", Tag.Update)
		apiV1.GET("/tags", Tag.List)

		apiV1.POST("/articles", Article.Create)
		apiV1.DELETE("/articles", Article.Delete)
		apiV1.PUT("/articles/:id", Article.Update)
		apiV1.GET("/articles", Article.List)
	}


	return r
	
	
}