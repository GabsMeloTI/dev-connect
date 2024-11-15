package model

import (
	"database/sql"
	"time"
	db "treads/db/sqlc"
)

type PostResponse struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Content   string    `json:"content"`
	ImageUrl  string    `json:"image_url"`
	Likes     int32     `json:"likes"`
	Shares    int32     `json:"shares"`
	CreatedAt time.Time `json:"created_at"`
}

type PostData struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Content   string    `json:"content"`
	ImageUrl  string    `json:"image_url"`
	Likes     int32     `json:"likes"`
	Shares    int32     `json:"shares"`
	CreatedAt time.Time `json:"created_at"`
	NameUser  string    `json:"name_user"`
	AvatarUrl string    `json:"avatar_url"`
}

type PostCreateDto struct {
	UserID   int64  `json:"user_id"`
	Content  string `json:"content"`
	ImageUrl string `json:"image_url"`
	Likes    int32  `json:"likes"`
	Shares   int32  `json:"shares"`
}

type PostUpdateDto struct {
	ID       int64  `json:"id"`
	Content  string `json:"content"`
	ImageUrl string `json:"image_url"`
}

type PostDeleteDto struct {
	ID int64 `json:"id"`
}

func (p *PostCreateDto) ParseCreateToPost() db.CreatePostParams {
	arg := db.CreatePostParams{
		UserID:  p.UserID,
		Content: p.Content,
		ImageUrl: sql.NullString{
			String: p.ImageUrl,
			Valid:  true,
		},
	}
	return arg
}

func (p *PostUpdateDto) ParseUpdateToPost() db.UpdatePostParams {
	arg := db.UpdatePostParams{
		Content: p.Content,
		ImageUrl: sql.NullString{
			String: p.ImageUrl,
			Valid:  true,
		},
		ID: p.ID,
	}
	return arg
}

func (p *PostResponse) ParseFromPostObject(result db.Post) {
	p.ID = result.ID
	p.UserID = result.UserID
	p.Content = result.Content
	p.ImageUrl = result.ImageUrl.String
	p.Likes = result.Likes.Int32
	p.Shares = result.Shares.Int32
	p.CreatedAt = time.Now()
}

func (p *PostData) ParseFromPostData(result db.GetAllPostsRow) {
	p.ID = result.ID
	p.UserID = result.UserID
	p.Content = result.Content
	p.ImageUrl = result.ImageUrl.String
	p.Likes = result.Likes.Int32
	p.Shares = result.Shares.Int32
	p.CreatedAt = time.Now()
	p.NameUser = result.Name
	p.AvatarUrl = result.AvatarUrl.String
}
