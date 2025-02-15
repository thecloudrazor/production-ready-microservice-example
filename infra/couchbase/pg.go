package couchbase

import (
	"context"
	"database/sql"
	"microservicetest/domain"
)

type PgRepository struct {
	db *sql.DB
}

func NewPgRepository() *PgRepository {
	return &PgRepository{}
}

func (r *PgRepository) GetProduct(ctx context.Context, id string) (*domain.Product, error) {
	return nil, nil
}

func (r *PgRepository) CreateProduct(ctx context.Context, product *domain.Product) error {
	return nil
}
