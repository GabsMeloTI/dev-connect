package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"treads/internal/model"
	"treads/internal/service"
)

type User struct {
	UserInterface service.UserInterface
}

func NewUser(UserInterface service.UserInterface) *User {
	return &User{UserInterface}
}

func (h *User) GetAllUsers(c echo.Context) error {
	var request model.UserResponse
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := h.UserInterface.GetAllUsers(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
