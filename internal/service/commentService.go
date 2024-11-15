package service

import (
	"context"
	"errors"
	db "treads/db/sqlc"
	"treads/internal/model"
	"treads/internal/repository"
)

type CommentInterface interface {
	CreateComment(context.Context, model.CommentCreateDto) (model.CommentResponse, error)
	UpdateComment(ctx context.Context, data model.CommentUpdateDto) (model.CommentResponse, error)
	DeleteComment(ctx context.Context, data model.CommentDeleteDto) error
	GetAllComments(ctx context.Context, postID int64, limit, offset int32) ([]model.CommentResponse, error)
	IncrementLikes(ctx context.Context, commentID int64) (model.CommentResponse, error)
	DecrementLikes(ctx context.Context, commentID int64) (model.CommentResponse, error)
}

type Comment struct {
	CommentInterface repository.CommentInterface
}

func NewComment(CommentInterface repository.CommentInterface) *Comment {
	return &Comment{CommentInterface: CommentInterface}
}

func (s *Comment) CreateComment(ctx context.Context, data model.CommentCreateDto) (model.CommentResponse, error) {
	arg := data.ParseCreateToComment()
	result, err := s.CommentInterface.CreateComment(ctx, arg)
	if err != nil {
		return model.CommentResponse{}, err
	}

	createCommentService := model.CommentResponse{}
	createCommentService.ParseFromCommentObject(result)

	return createCommentService, nil
}

func (s *Comment) UpdateComment(ctx context.Context, data model.CommentUpdateDto) (model.CommentResponse, error) {
	arg := data.ParseUpdateToComment()
	result, err := s.CommentInterface.UpdateComment(ctx, arg)
	if err != nil {
		return model.CommentResponse{}, err
	}

	updateCommentService := model.CommentResponse{}
	updateCommentService.ParseFromCommentObject(result)

	return updateCommentService, nil
}

func (s *Comment) DeleteComment(ctx context.Context, data model.CommentDeleteDto) error {
	if data.ID <= 0 || data.UserID <= 0 || data.PostID <= 0 {
		return errors.New("invalid input data")
	}

	comment, err := s.CommentInterface.GetCommentByID(ctx, data.ID)
	if err != nil {
		return errors.New("comment not found")
	}

	if comment.UserID != data.UserID {
		return errors.New("user not authorized to delete this comment")
	}

	arg := data.ParseDeleteToComment()
	err = s.CommentInterface.DeleteComment(ctx, arg)
	if err != nil {
		return err
	}

	return nil
}

func (s *Comment) GetAllComments(ctx context.Context, postID int64, limit, offset int32) ([]model.CommentResponse, error) {
	arg := db.GetAllCommentsParams{
		PostID: postID,
		Limit:  limit,
		Offset: offset,
	}
	comments, err := s.CommentInterface.GetAllComments(ctx, arg)
	if err != nil {
		return nil, err
	}

	var response []model.CommentResponse
	for _, comment := range comments {
		var resp model.CommentResponse
		resp.ParseFromCommentObject(comment)
		response = append(response, resp)
	}

	return response, nil
}

func (s *Comment) IncrementLikes(ctx context.Context, commentID int64) (model.CommentResponse, error) {
	if commentID <= 0 {
		return model.CommentResponse{}, errors.New("invalid comment ID")
	}

	result, err := s.CommentInterface.IncrementLikes(ctx, commentID)
	if err != nil {
		return model.CommentResponse{}, err
	}

	response := model.CommentResponse{}
	response.ParseFromCommentObject(result)

	return response, nil
}

func (s *Comment) DecrementLikes(ctx context.Context, commentID int64) (model.CommentResponse, error) {
	if commentID <= 0 {
		return model.CommentResponse{}, errors.New("invalid comment ID")
	}

	comment, err := s.CommentInterface.GetCommentByID(ctx, commentID)
	if err != nil {
		return model.CommentResponse{}, err
	}

	if comment.Likes.Int64 <= 0 {
		return model.CommentResponse{}, errors.New("cannot decrement likes below zero")
	}

	result, err := s.CommentInterface.DecrementLikes(ctx, commentID)
	if err != nil {
		return model.CommentResponse{}, err
	}

	response := model.CommentResponse{}
	response.ParseFromCommentObject(result)

	return response, nil
}
