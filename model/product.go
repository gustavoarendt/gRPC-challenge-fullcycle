package model

import (
	"github.com/asaskevich/govalidator"
	"math/rand"
)

type ProductRepositoryInterface interface {
	ListProducts() (*Products, error)
	CreateProduct(product *Product) (*Product, error)
}

type Products struct {
	Product []Product
}

type Product struct {
	ID          int32   `json:"id" gorm:"type:integer;primary_key" valid:"notnull"`
	Name        string  `json:"name" gorm:"type:varchar(255)" valid:"notnull"`
	Description string  `json:"description" gorm:"type:text" valid:"notnull"`
	Price       float32 `json:"price" gorm:"type:float" valid:"notnull"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func (p *Product) isValid() error {
	_, err := govalidator.ValidateStruct(p)

	if err != nil {
		return err
	}

	return nil
}

func NewProduct(name string, description string, price float32) (*Product, error) {
	product := Product{
		Name:        name,
		Description: description,
		Price:       price,
	}

	product.ID = int32(rand.Int())
	err := product.isValid()
	if err != nil {
		return nil, err
	}

	return &product, nil
}
