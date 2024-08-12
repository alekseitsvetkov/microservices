package grpc

import (
	"github.com/alekseytsvetkov/microservices/apps/gateway/internal/config"
	authpb "github.com/alekseytsvetkov/microservices/proto/auth"
	productpb "github.com/alekseytsvetkov/microservices/proto/product"
	userpb "github.com/alekseytsvetkov/microservices/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUser(cfg *config.Config) (userpb.UserServiceClient, error) {
	connection, err := grpc.NewClient(
		cfg.GRPC.Client.User.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	client := userpb.NewUserServiceClient(connection)

	return client, nil
}

func NewProduct(cfg *config.Config) (productpb.ProductServiceClient, error) {
	connection, err := grpc.NewClient(
		cfg.GRPC.Client.Product.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	client := productpb.NewProductServiceClient(connection)

	return client, nil
}

func NewAuth(cfg *config.Config) (authpb.AuthServiceClient, error) {
	connection, err := grpc.NewClient(
		cfg.GRPC.Client.Auth.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	client := authpb.NewAuthServiceClient(connection)

	return client, nil
}
