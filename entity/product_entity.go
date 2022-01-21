package entity

type Product struct {
	Id     int    `json:"id" form:"id"`
	UserId int    `json:"user_id" form:"user_id"`
	Name   string `json:"name" form:"name"`
	Price  int    `json:"price" form:"price"`
	Stock  int    `json:"stock" form:"stock"`
}

type CreateProductRequest struct {
	UserId int    `json:"user_id" form:"user_id"`
	Name   string `json:"name" form:"name"`
	Price  int    `json:"price" form:"price"`
	Stock  int    `json:"stock" form:"stock"`
}

type EditProductRequest struct {
	UserId int    `json:"user_id" form:"user_id"`
	Name   string `json:"name" form:"name"`
	Price  int    `json:"price" form:"price"`
	Stock  int    `json:"stock" form:"stock"`
}

type ProductResponse struct {
	Id     int    `json:"id" form:"id"`
	UserId int    `json:"user_id" form:"user_id"`
	Name   string `json:"name" form:"name"`
	Price  int    `json:"price" form:"price"`
	Stock  int    `json:"stock" form:"stock"`
}

func FormatProductResponse(product Product) ProductResponse {
	return ProductResponse{
		Id:     product.Id,
		UserId: product.UserId,
		Name:   product.Name,
		Price:  product.Price,
		Stock:  product.Stock,
	}
}
