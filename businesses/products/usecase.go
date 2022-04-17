package products

import (
	"context"
	"errors" 
	"time"
)

type ProductUseCase struct {
	repo ProductRepoInterface
	ctx time.Duration 
	
}



func NewUseCase(productRepo ProductRepoInterface, ctx time.Duration) *ProductUseCase {
	return &ProductUseCase{
		repo: 		productRepo,
		ctx:		ctx,
	}
}


func (usecase *ProductUseCase) Add(ctx context.Context, domain Domain) (Domain, error){
	if domain.Name == "" {
		return Domain{}, errors.New("nama belum di isi")
	}
	if domain.Description == "" {
		return Domain{}, errors.New("deskripsi belum di isi")
	}
	// if domain.ProductType_Id == 0 {
	// 	return Domain{}, errors.New("ProductType Id belum di isi")
	// }
	// if domain.Code == 0 {
	// 	return Domain{}, errors.New("code belum di isi")
	// }
	if domain.Price == 0 {
		return Domain{}, errors.New("price belum di isi")
	}
	if domain.Quantity == 0 {
		return Domain{}, errors.New("quantity belum di isi")
	}
	

	product, err := usecase.repo.Add(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	
	return product, nil
}