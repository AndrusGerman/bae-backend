package baehttp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewContext(ctx *gin.Context, baeHttp *Bae) *Context {
	return &Context{ginCtx: ctx, baeHttp: baeHttp}
}

type Context struct {
	ginCtx  *gin.Context
	baeHttp *Bae
}

func (ctx *Context) BindJSON(obj any) error {
	return ctx.ginCtx.ShouldBindJSON(obj)
}

func (ctx *Context) HandleSuccess(data any) {
	rsp := newResponse(true, "Success", data)
	ctx.ginCtx.JSON(http.StatusOK, rsp)
}

func (ctx *Context) HandleError(err error) {
	statusCode, ok := ctx.baeHttp.errorStatusMap[err]
	if !ok {
		statusCode = http.StatusInternalServerError
	}

	errMsg := err.Error()
	errRsp := newErrorResponse([]string{errMsg})
	ctx.ginCtx.JSON(statusCode, errRsp)
}
