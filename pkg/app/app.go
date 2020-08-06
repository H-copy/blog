package app


import (
	
	"net/http"
	
	"github.com/gin-gonic/gin"

	"blog/pkg/errcode"
	
)


type Page struct{
	Page int `json:"page"`
	PageSize int `json:"page_size"`
	TotalRows int `json:"total_rows"`
}

type Response struct{
	Ctx *gin.Context
}

func NewResponse(c *gin.Context) *Response{
	return &Response{c}
}

func (r *Response)ToResponse(data interface{}){
	if data == nil{
		data = gin.H{}
	}

	r.Ctx.JSON(http.StatusOK, data)
}


func (r *Response)ToResponseList(list interface{}, totalRows int){
	r.Ctx.JSON(http.StatusOK, gin.H{

		"list": list,
		"pager": Page{
			Page: GetPage(r.Ctx),
			PageSize: GetPageSize(r.Ctx),
			TotalRows: totalRows,
		},
		
	})
}


func (r *Response)ToErrorResponse(err *errcode.Error){

	r.Ctx.JSON(err.StatusCode(), gin.H{
		"code": err.Code(),
		"msg": err.Msg(),
		"detail": err.Details(),
	})
	
}




