package model

import (
	"database/sql"
	"time"
	db "treads/db/sqlc"
)

type PostRespose struct {
	ID        int32     `json:"id"`
	UserID    int32     `json:"user_id"`
	Content   string    `json:"content"`
	ImageUrl  string    `json:"image_url"`
	Likes     int32     `json:"likes"`
	Shares    int32     `json:"shares"`
	CreatedAt time.Time `json:"created_at"`
	LastLogin time.Time `json:"last_login"`
}

type PostCreateDto struct {
	UserID   int32  `json:"user_id"`
	Content  string `json:"content"`
	ImageUrl string `json:"image_url"`
	Likes    int32  `json:"likes"`
	Shares   int32  `json:"shares"`
}

type PostUpdateDto struct {
	ID       int32  `json:"id"`
	Content  string `json:"content"`
	ImageUrl string `json:"image_url"`
}

type PostDeleteDto struct {
	ID int32 `json:"id"`
}

func (p *PostCreateDto) ParseCreateToPost() db.CreatePostParams {
	arg := db.CreatePostParams{
		UserID: sql.NullInt32{
			Int32: p.UserID,
			Valid: true,
		},
		Content: p.Content,
		ImageUrl: sql.NullString{
			String: p.ImageUrl,
			Valid:  true,
		},
		Likes: sql.NullInt32{
			Int32: p.Likes,
			Valid: true,
		},
		Shares: sql.NullInt32{
			Int32: p.Shares,
			Valid: true,
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

func (p *PostRespose) ParseFromPostObject(result db.Post) {
	p.ID = result.ID
	p.UserID = result.UserID.Int32
	p.Content = result.Content
	p.ImageUrl = result.ImageUrl.String
	p.Likes = result.Likes.Int32
	p.Shares = result.Shares.Int32
	p.CreatedAt = time.Now()
}
