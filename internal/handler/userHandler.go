package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"treads/internal/model"
	"treads/internal/service"
	"treads/validation"
)

type User struct {
	UserInterface service.UserInterface
}

func NewUser(UserInterface service.UserInterface) *User {
	return &User{UserInterface}
}

func (h *User) CreateUser(c echo.Context) error {
	var request model.UserCreateDto
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := validation.Validate(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := h.UserInterface.CreateUser(c.Request().Context(), request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func (h *User) UpdateUser(c echo.Context) error {
	var request model.UserUpdateDto
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	data := model.UserUpdateDto{
		ID:        request.ID,
		Name:      request.Name,
		Username:  request.Username,
		Email:     request.Email,
		Bio:       request.Bio,
		AvatarUrl: request.AvatarUrl,
	}

	result, err := h.UserInterface.UpdateUser(c.Request().Context(), data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func (h *User) UpdatePassword(c echo.Context) error {
	var request model.UserRequestUpdatePasswordByUser
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := validation.Validate(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = h.UserInterface.UpdatePassword(c.Request().Context(), request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Success")
}

func (h *User) DisableUser(c echo.Context) error {
	idStr := c.Param("id")
	id, err := validation.ParseStringToInt64(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = h.UserInterface.DisableUser(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Usuário desabilitado com sucesso!")
}

func (h *User) DeleteUser(c echo.Context) error {
	idStr := c.Param("id")
	id, err := validation.ParseStringToInt64(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = h.UserInterface.DeleteUser(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Usuário deletado com sucesso!")
}

func (h *User) GetAllUsers(c echo.Context) error {
	usernameParams := c.Param("search")

	var request model.UserResponse
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := h.UserInterface.GetAllUsers(c.Request().Context(), usernameParams)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func (h *User) LoginUser(c echo.Context) error {
	var request model.LoginUserRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := h.UserInterface.UserLogin(c.Request().Context(), request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
