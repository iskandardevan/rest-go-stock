package products

import (
	"context"
	"rest-go-stock/businesses/product_types"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	Id        	uint		`gorm:"primaryKey"`
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	DeletedAt 	gorm.DeletedAt `gorm:"index"`
	// Code     			int
	ProductType_ID uint
	ProductType 	    product_types.Domain
	Name  				string
	Description     	string
	Price  				float64
	Quantity			int		`gorm:"default:0"`
}

type ProductUsecaseInterface interface {
	GetByID(id uint, ctx context.Context) (Domain, error)
	Add(ctx context.Context, domain Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
}

type ProductRepoInterface interface {
	GetByID(id uint, ctx context.Context) (Domain, error)
	Add(ctx context.Context, domain Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
}

