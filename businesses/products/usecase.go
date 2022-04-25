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


// func (usecase *ProductUseCase) Add(ctx context.Context, domain Domain) (Domain, error){
// 	if domain.Name == "" {
// 		return Domain{}, errors.New("nama belum di isi")
// 	}
// 	if domain.Description == "" {
// 		return Domain{}, errors.New("deskripsi belum di isi")
// 	}
// 	// if domain.ProductType_Id == 0 {
// 	// 	return Domain{}, errors.New("ProductType Id belum di isi")
// 	// }
// 	// if domain.Code == 0 {
// 	// 	return Domain{}, errors.New("code belum di isi")
// 	// }
// 	if domain.Price == 0 {
// 		return Domain{}, errors.New("price belum di isi")
// 	}
// 	if domain.Quantity == 0 {
// 		return Domain{}, errors.New("quantity belum di isi")
// 	}
	

// 	product, err := usecase.repo.Add(ctx, domain)
// 	if err != nil {
// 		return Domain{}, err
// 	}
	
// 	return product, nil
// }

func (usecase *ProductUseCase) ProductIn(ctx context.Context, domain Domain) (Domain, error){
	if domain.Name == "" {
		return Domain{}, errors.New("nama harus di isi")
	}
	if domain.Quantity == 0 {
		return Domain{}, errors.New("quantity harus diisi")
	}
	product, err := usecase.repo.ProductIn(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return product, nil
}

func (usecase *ProductUseCase) ProductOut(ctx context.Context, domain Domain) (Domain, error){
	if domain.Name == "" {
		return Domain{}, errors.New("nama harus di isi")
	}
	if domain.Quantity == 0 {
		return Domain{}, errors.New("quantity harus diisi")
	}
	product, err := usecase.repo.ProductOut(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return product, nil
}

func (usecase *ProductUseCase) GetAll(ctx context.Context) ([]Domain, error){
	product, err := usecase.repo.GetAll(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return product, nil
}

func (usecase *ProductUseCase) GetByID(id uint, ctx context.Context) (Domain, error){
	product, err := usecase.repo.GetByID(id, ctx)
	if err != nil {
		return Domain{}, errors.New("tidak ada product dengan ID tersebut")
	}
	if id == 0 {
		return Domain{}, errors.New("ID harus diisi")
	}
	return product, nil
}