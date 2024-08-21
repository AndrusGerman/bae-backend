package baehttp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Context interface {
	BindJSON(obj any) error
	HandleSuccess(data any) error
	HandleError(err error) error
	Param(paramName string) Param
}

var _ Context = (*ContextHandler)(nil)

func NewContextHandler(ctx *gin.Context, baeHttp *Bae) Context {
	return &ContextHandler{ginCtx: ctx, baeHttp: baeHttp}
}

type ContextHandler struct {
	ginCtx  *gin.Context
	baeHttp *Bae
}

func (ctx *ContextHandler) BindJSON(obj any) error {
	return ctx.ginCtx.ShouldBindJSON(obj)
}

func (ctx *ContextHandler) HandleSuccess(data any) error {
	rsp := newResponse(true, "Success", data)
	ctx.ginCtx.JSON(http.StatusOK, rsp)
	return nil
}

func (ctx *ContextHandler) HandleError(err error) error {
	statusCode, ok := ctx.baeHttp.errorStatusMap[err]
	if !ok {
		statusCode = http.StatusInternalServerError
	}

	errMsg := err.Error()
	errRsp := newErrorResponse([]string{errMsg})
	ctx.ginCtx.JSON(statusCode, errRsp)
	return nil
}

func (ctx *ContextHandler) Param(paramName string) Param {
	return Param(ctx.ginCtx.Param(paramName))
}
