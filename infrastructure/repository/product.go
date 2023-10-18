package repository

import (
	"fc01-grpc-product/model"
	"github.com/jinzhu/gorm"
)

type ProductRepositoryDb struct {
	Db *gorm.DB
}

func (p *ProductRepositoryDb) ListProducts() (*model.Products, error) {
	var products model.Products
	if err := p.Db.Find(&products.Product).Error; err != nil {
		return nil, err
	}

	if len(products.Product) == 0 {
		return &products, nil
	}

	return &products, nil
}

func (p *ProductRepositoryDb) CreateProduct(product *model.Product) (*model.Product, error) {
	err := p.Db.Save(product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}
