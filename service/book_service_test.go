package service_test

import (
	"database/sql"
	"designpattern/entity"
	"designpattern/repository"
	"designpattern/router"
	"designpattern/service"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func InitEchoTestAPIBook() (*echo.Echo, *sql.DB) {
	jwtSecret := os.Getenv("JWT_SECRET")
	connectionString := os.Getenv("DB_CONNECTION_STRING")
	db, err := entity.InitDB(connectionString)
	if err != nil {
		panic(err)
	}
	e := echo.New()
	router.UserRouter(e, db, jwtSecret)
	return e, db
}

func TestCreateBook(t *testing.T) {
	e, db := InitEchoTestAPIBook()
	defer db.Close()

	bookRepository := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepository)

	t.Run("Test Create Success", func(t *testing.T) {

		input := entity.CreateBooksRequest{
			Title:       "harry potter",
			Author:      "jk rowling",
			PublishedAt: "saat hujan",
		}

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/books")

		result, err := bookService.CreateBookService(input)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
}

func TestGetBook(t *testing.T) {
	// setting controller
	e, db := InitEchoTestAPIBook()
	defer db.Close()
	bookRepository := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepository)

	t.Run("TestGetAll", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/books")

		result, err := bookService.GetBooksService()
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
	t.Run("TestGetbyIdSuccess", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/books/:id")
		context.SetParamNames("id")

		result, err := bookService.GetBookByIdService(2)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
	t.Run("TestGetbyIdError", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/books/:id")
		context.SetParamNames("id")

		result, err := bookService.GetBookByIdService(1)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
}

func TestUpdateBook(t *testing.T) {
	// setting controller
	e, db := InitEchoTestAPIBook()

	bookRepository := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepository)

	t.Run("TestUpdateSuccess", func(t *testing.T) {
		input := entity.EditBooksRequest{
			Title:       "sam and friens",
			Author:      "lucas graham",
			PublishedAt: "12-03-2021",
		}
		req := httptest.NewRequest(http.MethodPut, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/books/:id")
		context.SetParamNames("id")

		result, err := bookService.UpdateBookService(13, input)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
	t.Run("TestUpdateError", func(t *testing.T) {
		input := entity.EditBooksRequest{
			Title:       "sam and friens",
			Author:      "lucas graham",
			PublishedAt: "12-03-2021",
		}
		req := httptest.NewRequest(http.MethodPut, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/books/:id")
		context.SetParamNames("id")

		result, err := bookService.UpdateBookService(6, input)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
}

func TestDeleteBook(t *testing.T) {
	// setting controller
	e, db := InitEchoTestAPIBook()

	bookRepository := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepository)

	t.Run("TestDeleteSuccess", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/books/:id")
		context.SetParamNames("id")

		result, err := bookService.DeleteBookService(14)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
	t.Run("TestDeleteError", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/books/:id")
		context.SetParamNames("id")

		result, err := bookService.DeleteBookService(7)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
}
