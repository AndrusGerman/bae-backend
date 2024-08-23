package baehttp

type HandlerConfig interface {
	GetPattern() string
	GetMethod() string
}

var _ HandlerConfig = (*HandlerBasicConfig)(nil)

type HandlerBasicConfig struct {
	Pattern     string
	Method      string
	Middlewares []Middleware
}

func (hc *HandlerBasicConfig) GetPattern() string {
	return hc.Pattern
}
func (hc *HandlerBasicConfig) GetMethod() string {
	return hc.Method
}

func NewHandlerConfig(method string, pattern string) HandlerConfig {
	return &HandlerBasicConfig{
		Method:  method,
		Pattern: pattern,
	}
}
