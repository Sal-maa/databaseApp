package repository

import (
	"database/sql"
	"designpattern/entity"
	"fmt"
)

type ProductRepo interface {
	GetAllProduct() ([]entity.Product, error)
	CreateProduct(product entity.Product) (entity.Product, error)
	GetProduct(idParam int) (entity.Product, error)
	UpdateProduct(product entity.Product) (entity.Product, error)
	DeleteProduct(product entity.Product) (entity.Product, error)
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *productRepository {
	return &productRepository{db}
}

func (r *productRepository) GetAllProduct() ([]entity.Product, error) {
	var products []entity.Product
	result, err := r.db.Query("SELECT id,user_id,name,price,stock FROM products")
	if err != nil {
		fmt.Println(err)
		return products, fmt.Errorf("failed to scan")
	}
	defer result.Close()

	for result.Next() {
		var product entity.Product
		err := result.Scan(&product.Id, &product.UserId, &product.Name, &product.Price, &product.Stock)
		if err != nil {
			return products, fmt.Errorf("failed to scan")
		}
		products = append(products, product)
	}
	return products, err
}

func (r *productRepository) CreateProduct(product entity.Product) (entity.Product, error) {
	_, err := r.db.Exec("INSERT INTO products(user_id, name, price, stock) VALUES(?,?,?,?)", product.UserId, product.Name, product.Price, product.Stock)
	return product, err
}

func (r *productRepository) GetProduct(idParam int) (entity.Product, error) {
	var product entity.Product
	result, err := r.db.Query("SELECT id,user_id,name,price,stock FROM products WHERE id=?", idParam)
	if err != nil {
		return product, fmt.Errorf("failed in query")
	}
	defer result.Close()

	if isExist := result.Next(); !isExist {
		fmt.Println("data not found", err)
	}
	errScan := result.Scan(&product.Id, &product.UserId, &product.Name, &product.Price, &product.Stock)
	fmt.Println(errScan)
	if errScan != nil {

		return product, fmt.Errorf("failed to read data")
	}
	if idParam == product.Id {
		return product, nil
	}
	return product, fmt.Errorf("product not found")
}

func (r *productRepository) UpdateProduct(product entity.Product) (entity.Product, error) {
	result, err := r.db.Exec("UPDATE products SET name=?, price=?, stock=? WHERE id=?", product.Name, product.Price, product.Stock, product.Id)
	if err != nil {
		return product, fmt.Errorf("failed to update data")
	}
	NotAffected, _ := result.RowsAffected()
	if NotAffected == 0 {
		return product, fmt.Errorf("failed to find data id")
	}
	return product, nil
}

func (r *productRepository) DeleteProduct(product entity.Product) (entity.Product, error) {
	_, err := r.db.Exec("DELETE FROM products WHERE id=?", product.Id)
	return product, err
}
