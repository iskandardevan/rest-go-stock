package products

import (
	"rest-go-stock/businesses/products"
	"rest-go-stock/drivers/repository/product_types"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	Id        uint    `gorm:"primaryKey"`
	Name      string `gorm:"unique"`
	ProductType_ID uint
	ProductType product_types.ProductType `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Description string
	Price float64
	Quantity int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (product *Product) ToDomain() products.Domain{
	res := products.Domain{
		Id: product.Id,
		Name: product.Name,
		ProductType_ID: product.ProductType_ID,
		ProductType: product.ProductType.ToDomain(),
		Description: product.Description,
		Price: product.Price,
		Quantity: product.Quantity,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
		DeletedAt:   gorm.DeletedAt{},
	}
	return res
}

func FromDomain(domain products.Domain) Product{
	return Product{
		Id: domain.Id,
		Name: domain.Name,
		ProductType_ID: domain.ProductType_ID,
		ProductType: product_types.FromDomain(domain.ProductType),
		Description: domain.Description,
		Price: domain.Price,
		Quantity: domain.Quantity,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}



func GetAllProducts(data []Product) []products.Domain{
	res := []products.Domain{}
	for _, val := range data{
		res = append(res, val.ToDomain())
	}
	return res
}

