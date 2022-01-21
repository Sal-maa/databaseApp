package handler

import (
	"designpattern/entity"
	"designpattern/helper"
	"designpattern/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type productHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) *productHandler {
	return &productHandler{productService}
}

// get all product
func (h *productHandler) GetProductsController(c echo.Context) error {
	products, err := h.productService.GetProductsService()

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed to read data"))
	}

	var data []entity.ProductResponse
	for i := 0; i < len(products); i++ {
		formatRes := entity.FormatProductResponse(products[i])
		data = append(data, formatRes)
	}
	return c.JSON(http.StatusOK, helper.SuccessResponses("success to read data", data))
}

// create new product
func (h *productHandler) CreateProductController(c echo.Context) error {
	productCreate := entity.CreateProductRequest{}
	if err := c.Bind(&productCreate); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed to bind data"))
	}
	userId := c.Get("currentUser").(entity.User)
	productCreate.UserId = userId.Id
	_, err := h.productService.CreateProductService(productCreate)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed to insert data"))
	}

	return c.JSON(http.StatusOK, helper.SuccessWithoutDataResponses("success insert data"))
}

//get product by id
func (h *productHandler) GetProductController(c echo.Context) error {
	idParam, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("data not found"))
	}

	product, err := h.productService.GetProductByIdService(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed to fetch data"))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponses("success to read data", product))
}

// update product by id
func (h *productHandler) UpdateProductController(c echo.Context) error {
	productUpdate := entity.EditProductRequest{}
	idParam, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("data not found"))
	}

	if err := c.Bind(&productUpdate); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed to bind data"))
	}

	userId := c.Get("currentUser").(entity.User)
	if productUpdate.UserId != userId.Id {
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("you dont have permission"))
	}

	productUp, err := h.productService.UpdateProductService(idParam, productUpdate)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed to update data"))
	}

	formatRes := entity.FormatProductResponse(productUp)
	return c.JSON(http.StatusOK, helper.SuccessResponses("success update data", formatRes))
}

// delete user by id
func (h *productHandler) DeleteProductController(c echo.Context) error {
	product := entity.Product{}
	idParam, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed convert id"))
	}

	userId := c.Get("currentUser").(entity.User)
	if product.UserId != userId.Id {
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("you dont have permission"))
	}

	_, err1 := h.productService.DeleteProductService(idParam)

	if err1 != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed delete data"))
	}

	return c.JSON(http.StatusOK, helper.SuccessWithoutDataResponses("Success delete data"))
}
