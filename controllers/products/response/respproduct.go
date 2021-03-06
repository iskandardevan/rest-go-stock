package response

import (
	"rest-go-stock/businesses/products"
	respProType "rest-go-stock/controllers/product_types/response"
	"time"

	"gorm.io/gorm"
)

type ProductResponse struct {
	ID uint `json:"id"`
	ProductType respProType.ProductTypeResponse `json:"product_type"`
	ProductType_ID uint    `json:"product_type_id"`
	Name           string  `json:"name"`
	Price          float64 `json:"price"`
	TotalPrice  		float64 `json:"total_price"`
	Description    string  `json:"description" `
	Quantity       int     `json:"quantity"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

func FromDomainProduct(domain products.Domain) ProductResponse {
	return ProductResponse{
		ID: domain.Id,
		ProductType: respProType.FromDomainProductType(domain.ProductType),
		ProductType_ID: domain.ProductType_ID,
		Name: domain.Name,
		Price: domain.Price,
		TotalPrice: domain.TotalPrice,
		Description: domain.Description,
		Quantity: domain.Quantity,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
	}
}

func FromDomainProductArray(data []products.Domain) []ProductResponse {
	var res []ProductResponse
	for _, val := range data {
		res = append(res, FromDomainProduct(val))
	}
	return res
}