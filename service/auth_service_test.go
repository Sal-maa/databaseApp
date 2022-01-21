package service_test

import (
	"database/sql"
	"designpattern/entity"
	"designpattern/repository"
	"designpattern/router"
	"designpattern/service"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func InitEchoTestAPIUser() (*echo.Echo, *sql.DB) {
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

func TestLogin(t *testing.T) {
	e, db := InitEchoTestAPIUser()
	defer db.Close()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)

	t.Run("Test Login Success", func(t *testing.T) {

		input := entity.Login{
			Name:     "sandra bahagia",
			Password: "123",
		}
		fmt.Println(input)
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")

		result, err := userService.LoginUserService(input)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
	t.Run("Test Login Error", func(t *testing.T) {

		input := entity.Login{
			Name:     "sandra bahagia",
			Password: "543",
		}
		fmt.Println(input)
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")

		result, err := userService.LoginUserService(input)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
}

func TestCreate(t *testing.T) {
	e, db := InitEchoTestAPIUser()
	defer db.Close()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)

	t.Run("Test Create Success", func(t *testing.T) {

		input := entity.CreateUserRequest{
			Name:     "suzana",
			Email:    "suz@na",
			Password: "123sus",
			Address:  "dirumah",
		}

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users")

		result, err := userService.CreateUserService(input)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
}

func TestGet(t *testing.T) {
	// setting controller
	e, db := InitEchoTestAPIUser()
	defer db.Close()
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)

	t.Run("TestGetAll", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users")

		result, err := userService.GetUsersService()
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
	t.Run("TestGetbyIdSuccess", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")
		context.SetParamNames("id")

		result, err := userService.GetUserByIdService(23)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
	t.Run("TestGetbyIdError", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")
		context.SetParamNames("id")

		result, err := userService.GetUserByIdService(20)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
	t.Run("TestGetbyNameSuccess", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")
		context.SetParamNames("id")

		result, err := userService.GetUserByNameService("suzana")
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
	t.Run("TestGetbyNameError", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")
		context.SetParamNames("id")

		result, err := userService.GetUserByNameService("aku")
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
}

func TestUpdate(t *testing.T) {
	// setting controller
	e, db := InitEchoTestAPIUser()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)

	t.Run("TestUpdateSuccess", func(t *testing.T) {
		input := entity.EditUserRequest{
			Name:     "suzini",
			Email:    "suz@ni",
			Password: "123susi",
			Address:  "dirumah aja",
		}
		req := httptest.NewRequest(http.MethodPut, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")
		context.SetParamNames("id")

		result, err := userService.UpdateUserService(39, input)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
	t.Run("TestUpdateError", func(t *testing.T) {
		input := entity.EditUserRequest{
			Name:     "suzene",
			Email:    "suz@ne",
			Password: "123suse",
			Address:  "dirumah ye",
		}
		req := httptest.NewRequest(http.MethodPut, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")
		context.SetParamNames("id")

		result, err := userService.UpdateUserService(20, input)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
}

func TestDelete(t *testing.T) {
	// setting controller
	e, db := InitEchoTestAPIUser()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)

	t.Run("TestDeleteSuccess", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")
		context.SetParamNames("id")

		result, err := userService.DeleteUserService(26)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
	t.Run("TestDeleteError", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")
		context.SetParamNames("id")

		result, err := userService.DeleteUserService(20)
		if err != nil {
			json.Marshal(err)
		}
		json.Marshal(result)
	})
}
