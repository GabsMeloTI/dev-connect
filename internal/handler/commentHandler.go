package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"treads/internal/model"
	"treads/internal/service"
	"treads/validation"
)

type Comment struct {
	CommentInterface service.CommentInterface
}

func NewComment(CommentInterface service.CommentInterface) *Comment {
	return &Comment{CommentInterface}
}

func (h *Comment) CreateComment(c echo.Context) error {
	var request model.CommentCreateDto
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := h.CommentInterface.CreateComment(c.Request().Context(), request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func (h *Comment) UpdateComment(c echo.Context) error {
	var request model.CommentUpdateDto
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := h.CommentInterface.UpdateComment(c.Request().Context(), request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func (h *Comment) DeleteComment(c echo.Context) error {
	var request model.CommentDeleteDto
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := h.CommentInterface.DeleteComment(c.Request().Context(), request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Comment deleted successfully")
}

func (h *Comment) GetAllComments(c echo.Context) error {
	postIDStr := c.QueryParam("post_id")
	postID, err := validation.ParseStringToInt64(postIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	limitStr := c.QueryParam("limit")
	offsetStr := c.QueryParam("offset")
	limit, err := validation.ParseStringToInt32(limitStr)
	if err != nil {
		limit = 10 // default limit
	}
	offset, err := validation.ParseStringToInt32(offsetStr)
	if err != nil {
		offset = 0 // default offset
	}

	result, err := h.CommentInterface.GetAllComments(c.Request().Context(), postID, limit, offset)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func (h *Comment) IncrementLikes(c echo.Context) error {
	idStr := c.Param("id")
	id, err := validation.ParseStringToInt64(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := h.CommentInterface.IncrementLikes(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func (h *Comment) DecrementLikes(c echo.Context) error {
	idStr := c.Param("id")
	id, err := validation.ParseStringToInt64(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := h.CommentInterface.DecrementLikes(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
