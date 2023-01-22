package test

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mock_repository "learn-mock/mocks"
	"learn-mock/models"
	"learn-mock/repository"
	"testing"
	"time"
)

var reqBody = &models.CreateProductRequest{
	Name: "product A",
}

func CreateProduct(repo repository.ProductRepository) (*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	data, err := repo.CreateProduct(ctx, reqBody)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func TestProductRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mock_repository.NewMockProductRepository(ctrl)
	mockRepo.EXPECT().CreateProduct(gomock.Any(), reqBody).Return(&models.Product{Name: "shampo"}, nil)
	data, err := CreateProduct(mockRepo)
	if assert.NoError(t, err) {
		assert.Equal(t, "shampo", data.Name)
	}
}
