package baehttp

import "github.com/gin-gonic/gin"

type RouterGroup struct {
	baeHttp        *Bae
	ginRouterGroup *gin.RouterGroup
}

func NewRouterGroup(baeHttp *Bae, relativePath string, handlers ...HandlerFunc) *RouterGroup {
	return &RouterGroup{
		baeHttp:        baeHttp,
		ginRouterGroup: baeHttp.core.Group(relativePath, baeHandlersToGin(baeHttp, handlers...)...), // handlers...),
	}
}

func (rg *RouterGroup) Group(relativePath string, handlers ...HandlerFunc) *RouterGroup {
	return &RouterGroup{
		baeHttp:        rg.baeHttp,
		ginRouterGroup: rg.ginRouterGroup.Group(relativePath, baeHandlersToGin(rg.baeHttp, handlers...)...),
	}
}

func (rg *RouterGroup) POST(relativePath string, handlers ...HandlerFunc) {
	rg.ginRouterGroup.POST(relativePath, baeHandlersToGin(rg.baeHttp, handlers...)...)
}

func (rg *RouterGroup) GET(relativePath string, handlers ...HandlerFunc) {
	rg.ginRouterGroup.GET(relativePath, baeHandlersToGin(rg.baeHttp, handlers...)...)
}
