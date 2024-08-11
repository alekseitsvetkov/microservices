package app

import (
	"example.com/microservices/apps/product/internal/config"
	"example.com/microservices/apps/product/internal/database"
	"example.com/microservices/apps/product/internal/grpc"
	"example.com/microservices/apps/product/internal/repositories"
	"example.com/microservices/apps/product/internal/services"
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
		fx.Invoke(database.Run, grpc.Run),
	)
}
