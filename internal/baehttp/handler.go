package baehttp

type Handler interface {
	Handler(ctx *Context) error
	Config() *Config
}

type HandlerFunc func(*Context) error
