package service

import (
	"treads/internal/repository"
)

type CommentInterface interface {
}

type Comment struct {
	CommentInterface repository.CommentInterface
}

func NewComment(CommentInterface repository.CommentInterface) *Comment {
	return &Comment{CommentInterface: CommentInterface}
}
