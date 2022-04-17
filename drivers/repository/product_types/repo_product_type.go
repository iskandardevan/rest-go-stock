package product_types

import (
	"context"
	"errors"
	"rest-go-stock/businesses/product_types"

	"gorm.io/gorm"
)

type productTypeRepo struct {
	DB *gorm.DB
}

func NewProductTypeRepo(DB *gorm.DB) *productTypeRepo {
	return &productTypeRepo{DB: DB}
}

func (Repo *productTypeRepo) GetByID(id uint, ctx context.Context) (product_types.Domain, error) {
	var productType ProductType
	err := Repo.DB.Find(&productType, "id=?", id)
	if err.Error != nil {
		return product_types.Domain{}, errors.New("product type not found")
	}
	return productType.ToDomain(), nil
}

func (Repo *productTypeRepo) AddProductType(ctx context.Context, domain product_types.Domain) (product_types.Domain, error) {
	productType := FromDomain(domain)
	err := Repo.DB.Create(&productType)
	if err.Error != nil {
		return product_types.Domain{}, err.Error
	}
	return productType.ToDomain(), nil
}