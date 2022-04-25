package products

import (
	"context"
	"errors"
	"rest-go-stock/businesses/products"

	"gorm.io/gorm"
)

type productRepo struct {
	DB *gorm.DB
}

func NewProductRepo(DB *gorm.DB) *productRepo {
	return &productRepo{DB: DB}
}

// func (Repo *productRepo) Add(ctx context.Context, domain products.Domain) (products.Domain, error) {
// 	product := Product{
// 		Name: domain.Name,
// 		ProductTypeID: domain.ProductType_ID,
// 		Description: domain.Description,
// 		Price: domain.Price,
// 		Quantity: domain.Quantity,
// 	}
// 	err := Repo.DB.Create(&product)
// 	if err.Error != nil {
// 		return products.Domain{}, err.Error
// 	}
// 	return product.ToDomain(), nil
// }

func (Repo *productRepo) ProductIn(ctx context.Context, domain products.Domain) (products.Domain, error) {
	product := Product{}
	err := Repo.DB.Where("name=?", domain.Name).First(&product).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		product.ProductTypeID = domain.ProductType_ID
		product.Description = domain.Description
		product.Name = domain.Name
		product.Price = domain.Price
		product.Quantity = domain.Quantity
	} else if product.Name == domain.Name {
		product.ProductTypeID = domain.ProductType_ID
		product.Name = domain.Name
		product.Description = domain.Description
		product.Price = domain.Price
		product.Quantity += domain.Quantity
	}  else {
		return products.Domain{}, err
	}
	err = Repo.DB.Save(&product).Error
	if err != nil {
		return products.Domain{}, err
	}
	return product.ToDomain(), nil
}

func (Repo *productRepo) ProductOut(ctx context.Context, domain products.Domain) (products.Domain, error) {
	product := Product{}
	err := Repo.DB.Where("name=?", domain.Name).First(&product).Error
 
	if errors.Is(err, gorm.ErrRecordNotFound) {
		product.ProductTypeID = domain.ProductType_ID
		product.Description = domain.Description
		product.Price = domain.Price
		product.TotalPrice = (product.Price * float64(domain.Quantity))
		product.Quantity = domain.Quantity
	} else if product.Name == domain.Name && product.Quantity >= domain.Quantity {
		product.Quantity -= domain.Quantity
		product.TotalPrice = (product.Price * float64(domain.Quantity)) + (product.Price * float64(domain.Quantity) * 2/10)
	}  else {
		return products.Domain{}, err
	}
	err = Repo.DB.Save(&product).Error
	if err != nil {
		return products.Domain{}, err
	}
	return product.ToDomain(), nil
}

// (harga barang * stok keluar) + (harga barang * stok keluar * 0.2)

func (Repo *productRepo) GetAll(ctx context.Context) ([]products.Domain, error) {
	var product []Product
	// err := Repo.DB.Find(&product)
	err := Repo.DB.Preload("ProductType").Find(&product)
	if err.Error != nil {
		return []products.Domain{}, err.Error
	}
	return GetAllProducts(product), nil
}

func (Repo *productRepo) GetByID(id uint, ctx context.Context ) (products.Domain, error){
	var product Product
	err := Repo.DB.Find(&product, "id=?", id)
	if err.Error != nil {
		return products.Domain{}, err.Error
	}
	return product.ToDomain(), nil
}