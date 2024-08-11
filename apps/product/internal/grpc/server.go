package grpc

import (
	"context"
	"errors"

	"example.com/microservices/apps/product/internal/dto"
	"example.com/microservices/apps/product/internal/repositories"
	"example.com/microservices/apps/product/internal/services"
	pb "example.com/microservices/libs/grpc/product"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	pb.UnimplementedProductServiceServer

	service services.Service
}

func NewServer(service services.Service) *server {
	return &server{
		service: service,
	}
}

func (s *server) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*emptypb.Empty, error) {
	input := &dto.CreateInput{
		Title: req.Title,
	}

	if err := s.service.Create(ctx, input); err != nil {
		if errors.Is(err, repositories.ErrAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *server) GetAllProducts(ctx context.Context, _ *emptypb.Empty) (*pb.Products, error) {
	products, err := s.service.GetAll(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.Products{
		Products: products,
	}, nil
}

func (s *server) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*emptypb.Empty, error) {
	input := &dto.UpdateInput{
		Title: req.Title,
	}

	if err := s.service.Update(ctx, req.Id, input); err != nil {
		if errors.Is(err, repositories.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *server) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*emptypb.Empty, error) {
	if err := s.service.Delete(ctx, req.Id); err != nil {
		if errors.Is(err, repositories.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
