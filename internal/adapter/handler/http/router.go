package http

import (
	"bae-backend/internal/baehttp"
	"bae-backend/internal/core/domain"

	"go.uber.org/fx"
)

type RouterDto struct {
	fx.In
	CoreBae  *baehttp.Bae
	Handlers []baehttp.Handler `group:"routes"`
}

// AddRouter creates a new HTTP router
func AddRouter(rdto RouterDto) {
	var bae = rdto.CoreBae

	bae.Use(
		baehttp.Cors(baehttp.CorsConfig{AllowAllOrigins: true}),
		baehttp.Recovery(),
	).ErrorStatusMap(domain.ErrorStatusMap)

	for _, handler := range rdto.Handlers {
		bae.Add(handler)
	}
}

func AsRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(baehttp.Handler)),
		fx.ResultTags(`group:"routes"`),
	)
}
