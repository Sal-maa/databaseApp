package handler

import (
	"designpattern/entity"
	"designpattern/helper"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *userHandler) AuthController(c echo.Context) error {
	login := entity.Login{}

	if err := c.Bind(&login); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("input must be string"))
	}
	// fmt.Printf("%#v\n", login)

	user, err := h.userService.LoginUserService(login)
	if err != nil {
		fmt.Println(err)
		fmt.Println(user)
		return c.JSON(http.StatusBadRequest, helper.FailedResponses("failed to login data"))
	}

	// membuat token
	token, err := h.authService.GenerateToken(user.Name)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, helper.FailedResponses("cannot create token"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponses("login success", token))

}
