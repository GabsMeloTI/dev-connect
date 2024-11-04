package service

import (
	"context"
	"treads/internal/model"
	"treads/internal/repository"
)

type PostInterface interface {
	CreatePost(context.Context, model.PostCreateDto) (model.PostResponse, error)
	UpdatePost(context.Context, model.PostUpdateDto) (model.PostResponse, error)
	DeletePost(context.Context, int64) error
	GetAllPost(context.Context) ([]model.PostResponse, error)
	GetAllPostByUser(context.Context, int64) ([]model.PostResponse, error)
}

type Post struct {
	PostInterface repository.PostInterface
}

func NewPost(PostInterface repository.PostInterface) *Post {
	return &Post{PostInterface: PostInterface}
}

func (s *Post) GetAllPost(ctx context.Context) ([]model.PostResponse, error) {
	results, err := s.PostInterface.GetAllPosts(ctx)
	if err != nil {
		return nil, err
	}

	getAllPost := model.PostResponse{}
	var postsResponse []model.PostResponse
	for _, result := range results {
		getAllPost.ParseFromPostObject(result)
		postsResponse = append(postsResponse, getAllPost)
	}

	return postsResponse, nil
}

// TODO: ARRUMAR VALIDAÇÃO PARA NÃO CRIAR COM MESMO E-MAIL OU USER.
func (s *Post) CreatePost(ctx context.Context, data model.PostCreateDto) (model.PostResponse, error) {
	arg := data.ParseCreateToPost()
	result, err := s.PostInterface.CreatePost(ctx, arg)
	if err != nil {
		return model.PostResponse{}, err
	}

	createPostService := model.PostResponse{}
	createPostService.ParseFromPostObject(result)

	return createPostService, nil
}

func (s *Post) UpdatePost(ctx context.Context, data model.PostUpdateDto) (model.PostResponse, error) {
	arg := data.ParseUpdateToPost()
	result, err := s.PostInterface.UpdatePost(ctx, arg)
	if err != nil {
		return model.PostResponse{}, err
	}

	updatePostService := model.PostResponse{}
	updatePostService.ParseFromPostObject(result)

	return updatePostService, nil
}

func (s *Post) DeletePost(ctx context.Context, id int64) error {
	err := s.PostInterface.DeletePost(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *Post) GetAllPostByUser(ctx context.Context, id int64) ([]model.PostResponse, error) {
	results, err := s.PostInterface.GetAllPostsByUser(ctx, id)
	if err != nil {
		return nil, err
	}

	var postsResponse []model.PostResponse
	for _, result := range results {
		var post model.PostResponse
		post.ParseFromPostObject(result)
		postsResponse = append(postsResponse, post)
	}

	return postsResponse, nil
}
