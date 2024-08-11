package clients

import (
	"context"

	"example.com/microservices/apps/gateway/internal/grpc/models"
	"example.com/microservices/libs/grpc"
	productpb "example.com/microservices/libs/grpc/product"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (c *Client) CreateProduct(ctx context.Context, req *models.CreateProductRequest) error {
	if _, err := c.product.CreateProduct(
		ctx,
		&productpb.CreateProductRequest{
			Title: req.Title,
		},
	); err != nil {
		grpcErr := grpc.ParseError(err)

		if grpcErr.Code() == codes.AlreadyExists {
			return grpcErr.Error()
		}

		return grpcErr.Error()
	}

	return nil
}

func (c *Client) GetAllProducts(ctx context.Context) ([]*productpb.Product, error) {
	res, err := c.product.GetAllProducts(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, grpc.ParseError(err).Error()
	}

	return res.Products, nil
}

func (c *Client) UpdateProduct(ctx context.Context, id uuid.UUID, req *models.UpdateProductRequest) error {
	if _, err := c.product.UpdateProduct(
		ctx,
		&productpb.UpdateProductRequest{
			Id:    id.String(),
			Title: req.Title,
		},
	); err != nil {
		grpcErr := grpc.ParseError(err)

		if grpcErr.Code() == codes.NotFound {
			return grpcErr.Error()
		}

		return grpcErr.Error()
	}

	return nil
}

func (c *Client) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	if _, err := c.product.DeleteProduct(
		ctx,
		&productpb.DeleteProductRequest{
			Id: id.String(),
		},
	); err != nil {
		grpcErr := grpc.ParseError(err)

		if grpcErr.Code() == codes.NotFound {
			return grpcErr.Error()
		}

		return grpcErr.Error()
	}

	return nil
}
