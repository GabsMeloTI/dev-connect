package service

import (
	"context"
	"database/sql"
	"go.uber.org/mock/gomock"
	"testing"
	"time"
	db "treads/db/sqlc"
	"treads/internal/model"
	"treads/internal/service"
	"treads/tests/mocks"
)

//func TestGetAllPosts(t *testing.T) {
//	control := gomock.NewController(t)
//	defer control.Finish()
//
//	ctx := context.Background()
//
//	postResponse := model.PostRespose{
//		ID:        1,
//		UserID:    1,
//		Content:   "Teste unitário",
//		ImageUrl:  "teste.png",
//		Likes:     0,
//		Shares:    0,
//		CreatedAt: time.Now(),
//	}
//
//	expectedDbPost := db.Post{
//		ID: 1,
//		UserID: sql.NullInt32{
//			Int32: 1,
//			Valid: true,
//		},
//		Content: postResponse.Content,
//		ImageUrl: sql.NullString{
//			String: postResponse.ImageUrl,
//			Valid:  true,
//		},
//		Likes: sql.NullInt32{
//			Int32: 0,
//			Valid: true,
//		},
//		Shares: sql.NullInt32{
//			Int32: 0,
//			Valid: true,
//		},
//	}
//
//	expectResponse := model.PostRespose{}
//	expectResponse.ParseFromPostObject(expectedDbPost)
//
//	repo := mocks.NewMockPostInterface(control)
//	service := service.NewPost(repo)
//
//	repo.EXPECT().CreatePost(ctx, expectedDbPost).Return(expectedDbPost, nil)
//
//	result, err := service.GetAllPost(ctx)
//
//	if err != nil {
//		t.Errorf("expected no error, got %v", err)
//	}
//
//
//}

func TestCreatePost(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	ctx := context.Background()

	postCreateDto := model.PostCreateDto{
		UserID:   1,
		Content:  "Teste unitário",
		ImageUrl: "teste.png",
	}

	expectedCreateParams := db.CreatePostParams{
		UserID: sql.NullInt32{
			Int32: 1,
			Valid: true,
		},
		Content: postCreateDto.Content,
		ImageUrl: sql.NullString{
			String: postCreateDto.ImageUrl,
			Valid:  true,
		},
		Likes: sql.NullInt32{
			Int32: 0,
			Valid: true,
		},
		Shares: sql.NullInt32{
			Int32: 0,
			Valid: true,
		},
	}

	createdAtTime := time.Now()

	expectedDbPost := db.Post{
		ID: 1,
		UserID: sql.NullInt32{
			Int32: 1,
			Valid: true,
		},
		Content: postCreateDto.Content,
		ImageUrl: sql.NullString{
			String: postCreateDto.ImageUrl,
			Valid:  true,
		},
		Likes: sql.NullInt32{
			Int32: 0,
			Valid: true,
		},
		Shares: sql.NullInt32{
			Int32: 0,
			Valid: true,
		},
		CreatedAt: createdAtTime,
	}

	expectedResponse := model.PostRespose{}
	expectedResponse.ParseFromPostObject(expectedDbPost)

	repo := mocks.NewMockPostInterface(control)
	service := service.NewPost(repo)

	repo.EXPECT().CreatePost(ctx, expectedCreateParams).Return(expectedDbPost, nil)

	result, err := service.CreatePost(ctx, postCreateDto)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if result.ID != expectedResponse.ID ||
		result.UserID != expectedResponse.UserID ||
		result.Content != expectedResponse.Content ||
		result.ImageUrl != expectedResponse.ImageUrl ||
		result.Likes != expectedResponse.Likes ||
		result.Shares != expectedResponse.Shares {
		t.Errorf("expected %v, got %v", expectedResponse, result)
	}

	if !isTimeApproxEqual(result.CreatedAt, expectedResponse.CreatedAt, time.Millisecond) {
		t.Errorf("expected CreatedAt to be %v, but got %v", expectedResponse.CreatedAt, result.CreatedAt)
	}
}

func isTimeApproxEqual(t1, t2 time.Time, tolerance time.Duration) bool {
	diff := t1.Sub(t2)
	return diff < tolerance && diff > -tolerance
}

func TestPutPost(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	ctx := context.Background()

	postUpdateDto := model.PostUpdateDto{
		ID:       1,
		Content:  "Teste unitário",
		ImageUrl: "teste.png",
	}

	expectedUpdateParams := db.UpdatePostParams{
		ID:      postUpdateDto.ID,
		Content: postUpdateDto.Content,
		ImageUrl: sql.NullString{
			String: postUpdateDto.ImageUrl,
			Valid:  true,
		},
	}

	expectedDbPost := db.Post{
		ID: 1,
		UserID: sql.NullInt32{
			Int32: 1,
			Valid: true,
		},
		Content: postUpdateDto.Content,
		ImageUrl: sql.NullString{
			String: postUpdateDto.ImageUrl,
			Valid:  true,
		},
		Likes: sql.NullInt32{
			Int32: 0,
			Valid: true,
		},
		Shares: sql.NullInt32{
			Int32: 0,
			Valid: true,
		},
		CreatedAt: time.Now(),
	}

	repo := mocks.NewMockPostInterface(control)
	service := service.NewPost(repo)

	repo.EXPECT().UpdatePost(ctx, expectedUpdateParams).Return(expectedDbPost, nil)

	result, err := service.UpdatePost(ctx, postUpdateDto)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	expectedResponse := model.PostRespose{}
	expectedResponse.ParseFromPostObject(expectedDbPost)

	if result.ID != expectedResponse.ID ||
		result.UserID != expectedResponse.UserID ||
		result.Content != expectedResponse.Content ||
		result.ImageUrl != expectedResponse.ImageUrl ||
		result.Likes != expectedResponse.Likes ||
		result.Shares != expectedResponse.Shares {
		t.Errorf("expected %v, got %v", expectedResponse, result)
	}
}

func TestDeletePost(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	ctx := context.Background()
	id := int32(1)

	repo := mocks.NewMockPostInterface(control)
	service := service.NewPost(repo)

	repo.EXPECT().DeletePost(ctx, id).Return(nil)

	err := service.DeletePost(ctx, id)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}
