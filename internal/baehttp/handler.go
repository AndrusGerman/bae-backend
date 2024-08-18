package baehttp

type Handler interface {
	Handler(ctx Context) error
	Config() HandlerConfig
}
