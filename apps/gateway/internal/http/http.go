package http

import (
	"context"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/alekseytsvetkov/microservices/apps/gateway/internal/config"
	"go.uber.org/fx"
)

func Run(lc fx.Lifecycle, cfg *config.Config, server *handler.Server, middleware *Middleware) {
	mux := http.NewServeMux()

	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", middleware.Auth(server))

	svr := &http.Server{
		Addr:    cfg.HTTP.Address,
		Handler: mux,
	}

	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go svr.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return svr.Shutdown(ctx)
		},
	})
}
