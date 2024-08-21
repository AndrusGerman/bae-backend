package baehttp

type HandlerConfig interface {
	GetPattern() string
	GetMethod() string
	GetMiddlewares() []any
}

var _ HandlerConfig = (*HandlerBasicConfig)(nil)

type HandlerBasicConfig struct {
	Pattern     string
	Method      string
	Middlewares []any
}

func (hc *HandlerBasicConfig) GetPattern() string {
	return hc.Pattern
}
func (hc *HandlerBasicConfig) GetMethod() string {
	return hc.Method
}

func (hc *HandlerBasicConfig) GetMiddlewares() []any {
	return hc.Middlewares
}

func NewHandlerConfig(method string, pattern string, Middlewares ...any) HandlerConfig {
	return &HandlerBasicConfig{
		Method:      method,
		Pattern:     pattern,
		Middlewares: Middlewares,
	}
}
