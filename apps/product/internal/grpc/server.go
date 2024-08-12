package grpc

import (
	"context"
	"errors"

	"github.com/alekseytsvetkov/microservices/apps/product/internal/dto"
	"github.com/alekseytsvetkov/microservices/apps/product/internal/repositories"
	"github.com/alekseytsvetkov/microservices/apps/product/internal/services"
	pb "github.com/alekseytsvetkov/microservices/proto/product"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	pb.UnimplementedProductServiceServer

	service services.Service
}

func NewServer(service services.Service) *Server {
	return &Server{
		service: service,
	}
}

func (s *Server) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*emptypb.Empty, error) {
	input := &dto.CreateInput{
		Title:       req.Title,
		Description: req.Description,
	}

	if err := s.service.Create(ctx, req.UserId, input); err != nil {
		if errors.Is(err, repositories.ErrAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) ListProducts(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	products, err := s.service.GetAll(ctx, req.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.ListProductsResponse{
		Products: products,
	}, nil
}

func (s *Server) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*emptypb.Empty, error) {
	input := &dto.UpdateInput{
		Title:       req.Title,
		Description: req.Description,
	}

	if err := s.service.Update(ctx, req.Id, req.UserId, input); err != nil {
		if errors.Is(err, services.ErrNoFieldsToUpdate) {
			return nil, status.Error(codes.Internal, err.Error())
		}

		if errors.Is(err, repositories.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		if errors.Is(err, repositories.ErrAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*emptypb.Empty, error) {
	if err := s.service.Delete(ctx, req.Id, req.UserId); err != nil {
		if errors.Is(err, repositories.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
