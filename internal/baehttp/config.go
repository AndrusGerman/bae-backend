package baehttp

type Config struct {
	Pattern    string
	Method     string
	Middleware []IMiddleware
}
