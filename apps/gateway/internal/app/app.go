package app

import (
	"github.com/alekseytsvetkov/microservices/apps/gateway/internal/config"
	"github.com/alekseytsvetkov/microservices/apps/gateway/internal/gql"
	"github.com/alekseytsvetkov/microservices/apps/gateway/internal/gql/directives"
	"github.com/alekseytsvetkov/microservices/apps/gateway/internal/gql/resolvers"
	"github.com/alekseytsvetkov/microservices/apps/gateway/internal/grpc"
	"github.com/alekseytsvetkov/microservices/apps/gateway/internal/http"
	"go.uber.org/fx"
)

func New() *fx.App {
	return fx.New(
		fx.Provide(
			config.New,
			grpc.NewUser,
			grpc.NewProduct,
			grpc.NewAuth,
			directives.NewDirective,
			resolvers.NewResolver,
			gql.New,
			http.NewMiddleware,
		),
		fx.Invoke(http.Run),
	)
}
