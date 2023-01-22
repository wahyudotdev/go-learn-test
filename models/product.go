package models

type Product struct {
	Id   string `json:"id,omitempty" gorm:"primaryKey,column:id"`
	Name string `json:"name,omitempty" gorm:"column:name"`
}

func (r Product) TableName() string {
	return "products"
}

type GetProductRequest struct {
	Page  int `json:"page" query:"page" validate:"min=1"`
	Limit int `json:"limit" query:"limit" validate:"min=1,max=100"`
}

type CreateProductRequest struct {
	Name string `json:"name" form:"name" validate:"required"`
}
