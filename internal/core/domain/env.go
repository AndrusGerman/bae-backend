package domain

type Env string

const (
	EnvDevelopment Env = "development"
	EnvRelease     Env = "release"
	EnvTest        Env = "test"
)
