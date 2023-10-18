package usecase

import (
	"errors"
	"fc01-grpc-product/model"
)

type ProductUseCase struct {
	ProductRepository model.ProductRepositoryInterface
}

func (p *ProductUseCase) CreateProduct(name string, description string, price float32) (*model.Product, error) {
	product, err := model.NewProduct(name, description, price)
	if err != nil {
		return nil, err
	}

	_, _ = p.ProductRepository.CreateProduct(product)
	if product.ID > 0 {
		return product, nil
	}

	return nil, errors.New("unable to save this product")
}

func (p *ProductUseCase) ListProducts() (*model.Products, error) {
	products, err := p.ProductRepository.ListProducts()
	if err != nil {
		return nil, err
	}

	return products, nil
}
