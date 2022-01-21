package service

import (
	"designpattern/entity"
	"designpattern/repository"
	"fmt"
)

type ProductService interface {
	GetProductsService() ([]entity.Product, error)
	CreateProductService(productCreate entity.CreateProductRequest) (entity.Product, error)
	GetProductByIdService(id int) (entity.Product, error)
	UpdateProductService(id int, productUpdate entity.EditProductRequest) (entity.Product, error)
	DeleteProductService(id int) (entity.Product, error)
}

type productService struct {
	repository repository.ProductRepo
}

func NewProductService(repository repository.ProductRepo) *productService {
	return &productService{repository}
}

func (s *productService) GetProductsService() ([]entity.Product, error) {
	products, err := s.repository.GetAllProduct()
	return products, err
}

func (s *productService) CreateProductService(productCreate entity.CreateProductRequest) (entity.Product, error) {
	product := entity.Product{}
	product.UserId = productCreate.UserId
	product.Name = productCreate.Name
	product.Price = productCreate.Price
	product.Stock = productCreate.Stock

	createProduct, err := s.repository.CreateProduct(product)
	return createProduct, err
}

func (s *productService) GetProductByIdService(id int) (entity.Product, error) {
	product, err := s.repository.GetProduct(id)
	if err != nil {
		fmt.Println(err)
		return product, err
	}
	return product, nil
}

func (s *productService) UpdateProductService(id int, productUpdate entity.EditProductRequest) (entity.Product, error) {
	product, err := s.repository.GetProduct(id)
	if err != nil {
		return product, err
	}

	product.Name = productUpdate.Name
	product.Price = productUpdate.Price
	product.Stock = productUpdate.Stock

	updateProduct, err := s.repository.UpdateProduct(product)
	return updateProduct, err
}

func (s *productService) DeleteProductService(id int) (entity.Product, error) {
	productID, err := s.GetProductByIdService(id)
	if err != nil {
		return productID, err
	}

	deleteProduct, err := s.repository.DeleteProduct(productID)

	return deleteProduct, err

}
