package products

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	Id        	uint		`gorm:"primaryKey"`
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	DeletedAt 	gorm.DeletedAt `gorm:"index"`
	Code     			int
	ProductType_Id 		int
	Name  				string
	Description     	string
	Price  				float64
	Quantity			int
}

type ProductUsecaseInterface interface {
	GetByID(id uint, ctx context.Context) (Domain, error)
	AddProduct(ctx context.Context, domain Domain) (Domain, error)
}

type ProductRepoInterface interface {
	GetByID(id uint, ctx context.Context) (Domain, error)
	AddProduct(ctx context.Context, domain Domain) (Domain, error)
}

