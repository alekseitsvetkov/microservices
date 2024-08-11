package repositories

import (
	"context"

	"example.com/microservices/apps/product/internal/model"
)

type Repository interface {
	Create(context.Context, *model.Product) error
	GetAll(context.Context) ([]*model.Product, error)
	Get(context.Context, string) (*model.Product, error)
	Update(context.Context, string, *model.Product) error
	Delete(context.Context, string) error
}
