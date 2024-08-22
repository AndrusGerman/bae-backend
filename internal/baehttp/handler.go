package baehttp

type Handler interface {
	HandlerBase
	IHandlerConfig
}

type HandlerBase interface {
	Handler(ctx Context) error
}
type IHandlerConfig interface {
	Config() HandlerConfig
}

type IHandlerAdd interface {
	GetHandler() Handler
	GetMiddlewares() []IMiddleware
}

type HandlerAdd struct {
	Handler     Handler
	middlewares []IMiddleware
}

func NewHandlerAdd(handler Handler, middlewares ...IMiddleware) IHandlerAdd {
	return &HandlerAdd{
		Handler:     handler,
		middlewares: middlewares,
	}
}

func (ha *HandlerAdd) GetMiddlewares() []IMiddleware {
	return ha.middlewares
}

func (ha *HandlerAdd) GetHandler() Handler {
	return ha.Handler
}
