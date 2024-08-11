package http

import (
	"context"
	"fmt"
	"net/http"

	"example.com/microservices/apps/gateway/internal/config"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"go.uber.org/fx"
)

func Run(lc fx.Lifecycle, cfg *config.Config, server *handler.Server, handler *Handler) {
	mux := http.NewServeMux()

	mux.Handle("/", server)
	mux.Handle("/playground", playground.Handler("GraphQL Playground", "/"))

	s := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.HTTP.Host, cfg.HTTP.Port),
		Handler: mux,
	}

	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go s.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return s.Shutdown(ctx)
		},
	})
}
