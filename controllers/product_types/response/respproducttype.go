package response

import (
	"rest-go-stock/businesses/product_types"
	"time"

	"gorm.io/gorm"
)

type ProductTypeResponse struct {
	ID        uint            `json:"id"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

func FromDomainProductType(domain product_types.Domain) ProductTypeResponse {
	return ProductTypeResponse{
		ID:        domain.Id,
		Name:      domain.Name,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
	}
}