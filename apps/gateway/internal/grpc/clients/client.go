package clients

import (
	productpb "example.com/microservices/libs/grpc/product"
)

type Client struct {
	product productpb.ProductServiceClient
}

func NewClient(
	product productpb.ProductServiceClient,
) *Client {
	return &Client{
		product: product,
	}
}
