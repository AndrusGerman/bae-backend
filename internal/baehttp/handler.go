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
	GetMiddlewares() []Middleware
}

type HandlerAdd struct {
	Handler     Handler
	middlewares []Middleware
}

func NewHandlerAdd(handler Handler, middlewares ...Middleware) IHandlerAdd {
	return &HandlerAdd{
		Handler:     handler,
		middlewares: middlewares,
	}
}

func (ha *HandlerAdd) GetMiddlewares() []Middleware {
	return ha.middlewares
}

func (ha *HandlerAdd) GetHandler() Handler {
	return ha.Handler
}
