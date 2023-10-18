package grpc

import (
	"context"
	"fc01-grpc-product/application/grpc/pb"
	"fc01-grpc-product/application/usecase"
	"fmt"
)

type ProductGrpcService struct {
	ProductUseCase usecase.ProductUseCase
	pb.UnimplementedProductServiceServer
}

func (p *ProductGrpcService) CreateProduct(ctx context.Context, in *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	product, err := p.ProductUseCase.CreateProduct(in.Name, in.Description, in.Price)
	if err != nil {
		return nil, err
	}

	return &pb.CreateProductResponse{
		Product: &pb.Product{
			Id:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
		},
	}, nil
}

func (p *ProductGrpcService) FindProducts(ctx context.Context, in *pb.FindProductsRequest) (*pb.FindProductsResponse, error) {
	products, err := p.ProductUseCase.ListProducts()
	if err != nil {
		return nil, err
	}

	var grpcProducts []*pb.Product
	for _, modelProduct := range products.Product {
		grpcProduct := &pb.Product{
			Id:          modelProduct.ID,
			Name:        modelProduct.Name,
			Description: modelProduct.Description,
			Price:       modelProduct.Price,
		}
		grpcProducts = append(grpcProducts, grpcProduct)
	}

	fmt.Println(&products)

	return &pb.FindProductsResponse{
		Products: grpcProducts,
	}, nil
}

func NewProductGrpcService(usecase usecase.ProductUseCase) *ProductGrpcService {
	return &ProductGrpcService{
		ProductUseCase: usecase,
	}
}
