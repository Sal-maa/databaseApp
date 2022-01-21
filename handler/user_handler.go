package handler

import (
	"designpattern/entity"
	"designpattern/helper"
	"designpattern/middleware"
	"designpattern/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	authService middleware.JWTService
	userService service.UserService
}

func NewUserHandler(authService middleware.JWTService, userService service.UserService) *userHandler {
	return &userHandler{authService, userService}
}

// get all user
func (h *userHandler) GetUsersController(c echo.Context) error {
	users, err := h.userService.GetUsersService()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed to read data"))
	}

	var data []entity.UserResponse
	for i := 0; i < len(users); i++ {
		formatRes := entity.FormatUserResponse(users[i])
		data = append(data, formatRes)
	}
	return c.JSON(http.StatusOK, helper.SuccessResponses("success to read data", data))
}

// create new user
func (h *userHandler) CreateUserController(c echo.Context) error {
	userCreate := entity.CreateUserRequest{}
	if err := c.Bind(&userCreate); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed to bind data"))
	}
	_, err := h.userService.CreateUserService(userCreate)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed to insert data"))
	}
	return c.JSON(http.StatusCreated, helper.SuccessWithoutDataResponses("success insert data"))
}

//get user by id
func (h *userHandler) GetUserController(c echo.Context) error {
	idParam, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("data not found"))
	}

	user, err := h.userService.GetUserByIdService(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed to fetch data"))
	}

	formatRes := entity.FormatUserResponse(user)
	return c.JSON(http.StatusOK, helper.SuccessResponses("success to read data", formatRes))
}

// update user by id
func (h *userHandler) UpdateUserController(c echo.Context) error {
	userUpdate := entity.EditUserRequest{}
	idParam, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("data not found"))
	}

	if err := c.Bind(&userUpdate); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed to bind data"))
	}

	userUp, err := h.userService.UpdateUserService(idParam, userUpdate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed to fetch data"))
	}
	formatRes := entity.FormatUserResponse(userUp)
	return c.JSON(http.StatusOK, helper.SuccessResponses("success update data", formatRes))
}

// delete user by id
func (h *userHandler) DeleteUserController(c echo.Context) error {
	idParam, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed convert id"))
	}

	user, err1 := h.userService.DeleteUserService(idParam)
	if err1 != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed delete data"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponses("Success delete data", user))

}
