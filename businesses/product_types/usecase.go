package product_types

import (
	"context"
	"errors"
	"time"
)

type ProductTypeUseCase struct {
	repo ProductTypeRepoInterface
	ctx  time.Duration

}

func NewUseCase(productTypeRepo ProductTypeRepoInterface, ctx time.Duration) *ProductTypeUseCase {
	return &ProductTypeUseCase{
		repo: productTypeRepo,
		ctx:  ctx,
	}
}

func (usecase *ProductTypeUseCase) AddProductType(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Name == "" {
		return Domain{}, errors.New("nama belum di isi")
	}
	productType, err := usecase.repo.AddProductType(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return productType, nil
}

func (usecase *ProductTypeUseCase) GetByID(id uint, ctx context.Context) (Domain, error) {
	product_type, err := usecase.repo.GetByID(id, ctx)
	if err != nil {
		return Domain{}, errors.New("tidak ada product_type dengan ID tersebut")
	}
	if id == 0 {
		return Domain{}, errors.New("ID harus diisi")
	}
	return product_type, nil
}