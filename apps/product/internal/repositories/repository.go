package repositories

import (
	"context"

	"github.com/alekseytsvetkov/microservices/apps/product/internal/model"
)

type Repository interface {
	Create(context.Context, string, *model.Product) error
	GetAll(context.Context, string) ([]*model.Product, error)
	Update(context.Context, string, string, map[string]interface{}) error
	Delete(context.Context, string, string) error
}
