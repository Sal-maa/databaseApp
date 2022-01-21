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

type bookHandler struct {
	bookService service.BookService
}

func NewBookHandler(bookService service.BookService) *bookHandler {
	return &bookHandler{bookService}
}

// get all book
func (h *bookHandler) GetBooksController(c echo.Context) error {
	books, err := h.bookService.GetBooksService()

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed to read data"))
	}

	var data []entity.BooksResponse
	for i := 0; i < len(books); i++ {
		formatRes := entity.FormatBookResponse(books[i])
		data = append(data, formatRes)
	}
	return c.JSON(http.StatusOK, helper.SuccessResponses("success to read data", data))
}

// create new book
func (h *bookHandler) CreateBookController(c echo.Context) error {
	bookCreate := entity.CreateBooksRequest{}
	if err := c.Bind(&bookCreate); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed to bind data"))
	}
	_, err := h.bookService.CreateBookService(bookCreate)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed to insert data"))
	}

	return c.JSON(http.StatusOK, helper.SuccessWithoutDataResponses("success insert data"))
}

//get book by id
func (h *bookHandler) GetBookController(c echo.Context) error {
	idParam, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("data not found"))
	}

	book, err := h.bookService.GetBookByIdService(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed to fetch data"))
	}

	formatRes := entity.FormatBookResponse(book)
	return c.JSON(http.StatusOK, helper.SuccessResponses("success to read data", formatRes))
}

// update user by id
func (h *bookHandler) UpdateBookController(c echo.Context) error {
	bookUpdate := entity.EditBooksRequest{}
	idParam, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("data not found"))
	}

	if err := c.Bind(&bookUpdate); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed to bind data"))
	}

	bookUp, err := h.bookService.UpdateBookService(idParam, bookUpdate)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed to update data"))
	}

	formatRes := entity.BooksResponse(bookUp)
	return c.JSON(http.StatusOK, helper.SuccessResponses("success update data", formatRes))
}

// delete user by id
func (h *bookHandler) DeleteBookController(c echo.Context) error {
	idParam, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed convert id"))
	}

	book, err1 := h.bookService.DeleteBookService(idParam)
	if err1 != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed delete data"))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponses("Success delete data", book))
}
