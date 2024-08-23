package baehttp

import "bae-backend/internal/core/domain"

type ErrorStatusMap map[error]int

type Config struct {
	Mode           domain.Env
	Middleware     []Middleware
	ErrorStatusMap ErrorStatusMap
	HandlesAdd     []IHandlerAdd
}

func (config *Config) GetMode() domain.Env {
	if config.Mode == "" {
		return domain.EnvDevelopment
	}
	return config.Mode
}

func (config *Config) GetErrorStatusMap() ErrorStatusMap {
	if config.ErrorStatusMap == nil {
		return make(ErrorStatusMap)
	}
	return config.ErrorStatusMap
}
func (config *Config) GetMiddleware() []Middleware {
	return config.Middleware
}

func (config *Config) GetHandlesAdd() []IHandlerAdd {
	return config.HandlesAdd
}
