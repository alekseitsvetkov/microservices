package app

import (
	"example.com/microservices/apps/gateway/internal/config"
	"example.com/microservices/apps/gateway/internal/gql"
	"example.com/microservices/apps/gateway/internal/gql/resolvers"
	"example.com/microservices/apps/gateway/internal/grpc"
	"example.com/microservices/apps/gateway/internal/grpc/clients"
	"example.com/microservices/apps/gateway/internal/http"
	"go.uber.org/fx"
)

func New() *fx.App {
	return fx.New(
		fx.Provide(
			config.New,
			grpc.NewProduct,
			clients.NewClient,
			resolvers.NewResolver,
			gql.New,
			http.NewHandler,
		),
		fx.Invoke(http.Run),
	)
}
