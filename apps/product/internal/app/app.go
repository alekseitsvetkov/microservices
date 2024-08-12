package app

import (
	"github.com/alekseytsvetkov/microservices/apps/product/internal/config"
	"github.com/alekseytsvetkov/microservices/apps/product/internal/database"
	"github.com/alekseytsvetkov/microservices/apps/product/internal/grpc"
	"github.com/alekseytsvetkov/microservices/apps/product/internal/repositories"
	"github.com/alekseytsvetkov/microservices/apps/product/internal/services"
	"go.uber.org/fx"
)

func New() *fx.App {
	return fx.New(
		fx.Provide(
			config.New,
			database.New,
			fx.Annotate(repositories.NewPostgresRepository, fx.As(new(repositories.Repository))),
			fx.Annotate(services.NewService, fx.As(new(services.Service))),
			grpc.NewServer,
		),
		fx.Invoke(grpc.Run),
	)
}
