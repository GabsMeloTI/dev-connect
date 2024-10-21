package service

import (
	"context"
	"treads/internal/model"
	"treads/internal/repository"
)

type PostInterface interface {
	CreatePost(context.Context, model.PostCreateDto) (model.PostRespose, error)
	UpdatePost(context.Context, model.PostUpdateDto) (model.PostRespose, error)
	DeletePost(context.Context, int32) error
	GetAllPost(context.Context) ([]model.PostRespose, error)
}

type Post struct {
	PostInterface repository.PostInterface
}

func NewPost(PostInterface repository.PostInterface) *Post {
	return &Post{PostInterface: PostInterface}
}

func (s *Post) GetAllPost(ctx context.Context) ([]model.PostRespose, error) {
	results, err := s.PostInterface.GetAllPosts(ctx)
	if err != nil {
		return nil, err
	}

	getAllPost := model.PostRespose{}
	var postsResponse []model.PostRespose
	for _, result := range results {
		getAllPost.ParseFromPostObject(result)
		postsResponse = append(postsResponse, getAllPost)
	}

	return postsResponse, nil
}

func (s *Post) CreatePost(ctx context.Context, data model.PostCreateDto) (model.PostRespose, error) {
	arg := data.ParseCreateToPost()
	result, err := s.PostInterface.CreatePost(ctx, arg)
	if err != nil {
		return model.PostRespose{}, err
	}

	createPostService := model.PostRespose{}
	createPostService.ParseFromPostObject(result)

	return createPostService, nil
}

func (s *Post) UpdatePost(ctx context.Context, data model.PostUpdateDto) (model.PostRespose, error) {
	arg := data.ParseUpdateToPost()
	result, err := s.PostInterface.UpdatePost(ctx, arg)
	if err != nil {
		return model.PostRespose{}, err
	}

	updatePostService := model.PostRespose{}
	updatePostService.ParseFromPostObject(result)

	return updatePostService, nil
}

func (s *Post) DeletePost(ctx context.Context, id int32) error {
	err := s.PostInterface.DeletePost(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
