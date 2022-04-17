package request

import "rest-go-stock/businesses/product_types"

type ProductTypeRequest struct {
	Name        string  `json:"name"`
}

func (ProductType *ProductTypeRequest) ToDomain() *product_types.Domain {
	return &product_types.Domain{
		Name:        ProductType.Name,
	}
}