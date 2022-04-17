package request

import "rest-go-stock/businesses/products"

type ProductRequest struct {
	ProductType_ID uint    `json:"product_type_id"`
	Name           string  `json:"name"`
	Price          float64 `json:"price"`
	Description    string  `json:"description" `
}

func (Product *ProductRequest) ToDomain() *products.Domain {
	return &products.Domain{
		ProductType_ID: Product.ProductType_ID,
		Name:           Product.Name,
		Price:          Product.Price,
		Description:    Product.Description,
	}
}