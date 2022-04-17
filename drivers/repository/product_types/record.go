package product_types

import (
	"time"
	"rest-go-stock/businesses/product_types"
	"gorm.io/gorm"
)

type ProductType struct {
	Id        uint    `gorm:"primaryKey"`
	Name      string `gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (productType *ProductType) ToDomain() product_types.Domain {
	return product_types.Domain{
		Id:        productType.Id,
		Name:      productType.Name,
		CreatedAt: productType.CreatedAt,
		UpdatedAt: productType.UpdatedAt,
	}
}

func FromDomain(domain product_types.Domain) ProductType  {
	return ProductType{
		Id: 	  domain.Id,
		Name:     domain.Name,
		CreatedAt 	:domain.CreatedAt,
		UpdatedAt 	:domain.UpdatedAt,
	}

}

func GetAllProductTypes(data []ProductType) []product_types.Domain{
	res := []product_types.Domain{}
	for _, val := range data{
		res = append(res, val.ToDomain())
	}
	return res
}

