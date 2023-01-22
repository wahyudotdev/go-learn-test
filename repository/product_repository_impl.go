package repository

import (
	"context"
	"github.com/gofiber/fiber/v2/utils"
	"learn-mock/database"
	"learn-mock/models"
)

type ProductRepositoryImpl struct {
	db *database.Database
}

func (p ProductRepositoryImpl) CreateProduct(ctx context.Context, req *models.CreateProductRequest) (*models.Product, error) {
	id := utils.UUID()
	data := models.Product{
		Id:   id,
		Name: req.Name,
	}
	tx := p.db.WithContext(ctx).Model(models.Product{}).Create(&data)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &data, nil
}

func (p ProductRepositoryImpl) GetAll(ctx context.Context, req *models.GetProductRequest) (*[]models.Product, error) {
	products := make([]models.Product, 0)
	offset := (req.Page - 1) * req.Limit
	tx := p.db.WithContext(ctx).Raw("SELECT * FROM products LIMIT ? OFFSET ?", req.Limit, offset).Scan(&products)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &products, nil
}

func NewProductRepositoryImpl(db *database.Database) ProductRepository {
	return ProductRepositoryImpl{
		db: db,
	}
}
