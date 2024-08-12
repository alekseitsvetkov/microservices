package services

import (
	"context"

	"github.com/alekseytsvetkov/microservices/apps/product/internal/dto"
	"github.com/alekseytsvetkov/microservices/apps/product/internal/model"
	"github.com/alekseytsvetkov/microservices/apps/product/internal/repositories"
	pb "github.com/alekseytsvetkov/microservices/proto/product"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Service interface {
	Create(context.Context, string, *dto.CreateInput) error
	GetAll(context.Context, string) ([]*pb.Product, error)
	Update(context.Context, string, string, *dto.UpdateInput) error
	Delete(context.Context, string, string) error
}

type service struct {
	repository repositories.Repository
}

func NewService(repository repositories.Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Create(ctx context.Context, userID string, input *dto.CreateInput) error {
	if input.Title == "" {
		return ErrTitleIsRequired
	}

	product := &model.Product{
		Title:       input.Title,
		Description: input.Description,
	}

	return s.repository.Create(ctx, userID, product)
}

func (s *service) GetAll(ctx context.Context, userID string) ([]*pb.Product, error) {
	products, err := s.repository.GetAll(ctx, userID)
	if err != nil {
		return nil, err
	}

	var output []*pb.Product
	for _, product := range products {
		output = append(
			output,
			&pb.Product{
				Id:          product.ID,
				Title:       product.Title,
				Description: product.Description,
				CreatedAt:   timestamppb.New(product.CreatedAt),
			},
		)
	}

	return output, nil
}

func (s *service) Update(ctx context.Context, id string, userID string, input *dto.UpdateInput) error {
	updatedFields := make(map[string]interface{})

	if input.Title != nil && *input.Title != "" {
		updatedFields["title"] = *input.Title
	}

	if input.Description != nil {
		updatedFields["description"] = *input.Description
	}

	if len(updatedFields) == 0 {
		return ErrNoFieldsToUpdate
	}

	return s.repository.Update(ctx, id, userID, updatedFields)
}

func (s *service) Delete(ctx context.Context, id string, userID string) error {
	return s.repository.Delete(ctx, id, userID)
}
