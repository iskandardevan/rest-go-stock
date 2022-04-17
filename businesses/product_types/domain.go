package product_types

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
	Name  				string
}

type ProductTypeUsecaseInterface interface {
	GetByID(id uint, ctx context.Context) (Domain, error)
	AddProductType(ctx context.Context, domain Domain) (Domain, error)
}

type ProductTypeRepoInterface interface {
	GetByID(id uint, ctx context.Context) (Domain, error)
	AddProductType(ctx context.Context, domain Domain) (Domain, error)
}

