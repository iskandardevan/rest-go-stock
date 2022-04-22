package request

import "rest-go-stock/businesses/products"

type ProductInRequest struct {
	ProductType_ID uint    `json:"product_type_id"`
	Name           string  `json:"name"`
	Price          float64 `json:"price"`
	TotalPrice  		float64 `json:"total_price"`
	Description    string  `json:"description" `
	Quantity	   int     `json:"quantity"`
}

func (Product *ProductInRequest) ToDomain() *products.Domain {
	return &products.Domain{
		ProductType_ID: Product.ProductType_ID,
		Name:           Product.Name,
		Price:          Product.Price,
		TotalPrice:     Product.TotalPrice,
		Description:    Product.Description,
		Quantity:       Product.Quantity,
	}
}