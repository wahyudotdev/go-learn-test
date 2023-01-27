package models

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	Id          string          `json:"id,omitempty" gorm:"primaryKey,column:id"`
	Name        string          `json:"name,omitempty" gorm:"column:name"`
	Description string          `json:"description"`
	DeletedAt   *gorm.DeletedAt `json:"-"`
	CreatedAt   *time.Time      `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   *time.Time      `json:"updated_at"`
}

func (r Product) TableName() string {
	return "products"
}

func (r Product) BeforeUpdate(tx *gorm.DB) error {
	t := time.Now().UTC()
	r.UpdatedAt = &t
	return nil
}

type GetProductRequest struct {
	Page  int `json:"page" query:"page" validate:"min=1"`
	Limit int `json:"limit" query:"limit" validate:"min=1,max=100"`
}

type CreateProductRequest struct {
	Name        string `json:"name" form:"name" validate:"required"`
	Description string `form:"description"`
}

type UpdateProductRequest struct {
	Name        string `form:"name"`
	Description string `form:"description"`
}
