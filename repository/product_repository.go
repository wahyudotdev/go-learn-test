package repository

import (
	"context"
	"learn-mock/models"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, req *models.CreateProductRequest) (*models.Product, error)
	GetAll(ctx context.Context, req *models.GetProductRequest) (*[]models.Product, error)
}
