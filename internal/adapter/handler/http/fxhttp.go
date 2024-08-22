package http

import (
	"bae-backend/internal/adapter/config"
	"bae-backend/internal/baehttp"
	"fmt"
	"log/slog"

	"go.uber.org/fx"
)

type CreateConfigDto struct {
	fx.In
	ErrorStatusMap baehttp.ErrorStatusMap
	Middlewares    []baehttp.IMiddleware `group:"global_middleware"`
	AddHandlers    []baehttp.IHandlerAdd `group:"handlers_add"`
	ConfigHttp     *config.HTTP
}

func NewHttpConfig(dto CreateConfigDto) *baehttp.Config {
	return &baehttp.Config{
		Mode:           dto.ConfigHttp.Env,
		Middleware:     dto.Middlewares,
		ErrorStatusMap: dto.ErrorStatusMap,
		HandlesAdd:     dto.AddHandlers,
	}
}

func RunHttpServer(httpConfig *config.HTTP, baehttp *baehttp.Bae) error {
	listenAddr := fmt.Sprintf("%s:%s", httpConfig.URL, httpConfig.Port)
	// Start server
	slog.Info("Starting the HTTP server", "listen_address", listenAddr)
	var err = baehttp.Serve(listenAddr)
	if err != nil {
		slog.Error("Error starting the HTTP server", "error", err)
	}
	return err
}
