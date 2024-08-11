package repositories

import (
	"context"
	"database/sql"
	"errors"

	"example.com/microservices/apps/product/internal/model"
	"github.com/lib/pq"
)

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) Repository {
	return &postgresRepository{
		db: db,
	}
}

func (r *postgresRepository) Create(ctx context.Context, product *model.Product) error {
	query := "INSERT INTO products(title) VALUES($1)"

	if _, err := r.db.ExecContext(ctx, query, product.Title); err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code.Name() == "unique_violation" {
			return ErrAlreadyExists
		}
		return err
	}

	return nil
}

func (r *postgresRepository) GetAll(ctx context.Context) ([]*model.Product, error) {
	query := "SELECT id, title, created_at FROM products"

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*model.Product
	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.ID, &product.Title, &product.CreatedAt); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *postgresRepository) Get(ctx context.Context, id string) (*model.Product, error) {
	query := "SELECT id, title, created_at FROM products WHERE id = $1 LIMIT 1"

	var product model.Product
	if err := r.db.QueryRowContext(ctx, query, id).Scan(&product.ID, &product.Title, &product.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &product, nil
}

func (r *postgresRepository) Update(ctx context.Context, id string, product *model.Product) error {
	query := "UPDATE products SET title = $1 WHERE id = $2"

	result, err := r.db.ExecContext(ctx, query, product.Title, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}

func (r *postgresRepository) Delete(ctx context.Context, id string) error {
	query := "DELETE FROM products WHERE id = $1"

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}
