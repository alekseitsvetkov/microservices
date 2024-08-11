package grpc

import (
	"fmt"

	"example.com/microservices/apps/gateway/internal/config"
	productpb "example.com/microservices/libs/grpc/product"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewProduct(cfg *config.Config) (productpb.ProductServiceClient, error) {
	connection, err := grpc.NewClient(
		fmt.Sprintf("%s:%d", cfg.GRPC.Client.Product.Host, cfg.GRPC.Client.Product.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	client := productpb.NewProductServiceClient(connection)

	return client, nil
}
