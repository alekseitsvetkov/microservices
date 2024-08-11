package services

import (
	"context"

	"example.com/microservices/apps/product/internal/dto"
	"example.com/microservices/apps/product/internal/model"
	"example.com/microservices/apps/product/internal/repositories"
	pb "example.com/microservices/libs/grpc/product"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Service interface {
	Create(context.Context, *dto.CreateInput) error
	GetAll(context.Context) ([]*pb.Product, error)
	Update(context.Context, string, *dto.UpdateInput) error
	Delete(context.Context, string) error
}

type service struct {
	repository repositories.Repository
}

func NewService(repository repositories.Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Create(ctx context.Context, input *dto.CreateInput) error {
	if input.Title == "" {
		return repositories.ErrTitleRequired
	}

	product := &model.Product{
		Title: input.Title,
	}

	return s.repository.Create(ctx, product)
}

func (s *service) GetAll(ctx context.Context) ([]*pb.Product, error) {
	products, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var output []*pb.Product
	for _, product := range products {
		output = append(output,
			&pb.Product{
				Id:        product.ID,
				Title:     product.Title,
				CreatedAt: timestamppb.New(product.CreatedAt),
			},
		)
	}

	return output, nil
}

func (s *service) Update(ctx context.Context, id string, input *dto.UpdateInput) error {
	product, err := s.repository.Get(ctx, id)
	if err != nil {
		return err
	}

	if input.Title == "" {
		return repositories.ErrTitleRequired
	}

	product.Title = input.Title

	return s.repository.Update(ctx, id, product)
}

func (s *service) Delete(ctx context.Context, id string) error {
	return s.repository.Delete(ctx, id)
}
