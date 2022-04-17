package products

import (
	"context"
	"rest-go-stock/businesses/products"

	"gorm.io/gorm"
)

type productRepo struct {
	DB *gorm.DB
}

func NewProductRepo(DB *gorm.DB) *productRepo {
	return &productRepo{DB: DB}
}

func (Repo *productRepo) Add(ctx context.Context, domain products.Domain) (products.Domain, error) {
	product := Product{
		Name: domain.Name,
		ProductType_ID: domain.ProductType_ID,
		Description: domain.Description,
		Price: domain.Price,
		Quantity: domain.Quantity,
	}
	err := Repo.DB.Create(&product)
	if err.Error != nil {
		return products.Domain{}, err.Error
	}
	return product.ToDomain(), nil
}