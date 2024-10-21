package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"treads/internal/model"
	"treads/internal/service"
)

type Post struct {
	PostInterface service.PostInterface
}

func NewPost(PostInterface service.PostInterface) *Post {
	return &Post{PostInterface}
}

func (h *Post) GetAllPosts(c echo.Context) error {
	var request model.PostRespose
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := h.PostInterface.GetAllPost(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func (h *Post) CreatePost(c echo.Context) error {
	var request model.PostCreateDto
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := h.PostInterface.CreatePost(c.Request().Context(), request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func (h *Post) UpdatePost(c echo.Context) error {
	var request model.PostUpdateDto
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := h.PostInterface.UpdatePost(c.Request().Context(), request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func (h *Post) DeletePost(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result := h.PostInterface.DeletePost(c.Request().Context(), int32(id))

	return c.JSON(http.StatusOK, result)
}
