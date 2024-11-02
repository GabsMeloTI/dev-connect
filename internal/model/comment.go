package model

import (
	"database/sql"
	"time"
	db "treads/db/sqlc"
)

type CommentResponse struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	PostID    int64     `json:"post_id"`
	Content   string    `json:"content"`
	Likes     int64     `json:"likes"`
	CreatedAt time.Time `json:"created_at"`
}

type CommentCreateDto struct {
	UserID  int64  `json:"user_id"`
	PostID  int64  `json:"post_id"`
	Content string `json:"content"`
}

type CommentUpdateDto struct {
	ID      int64  `json:"id"`
	Content string `json:"content"`
}

type CommentUpdateLikesDto struct {
	ID    int64 `json:"id"`
	Likes int64 `json:"likes"`
}

func (p *CommentCreateDto) ParseCreateToComment() db.CreateCommentParams {
	arg := db.CreateCommentParams{
		UserID:  p.UserID,
		PostID:  p.PostID,
		Content: p.Content,
	}
	return arg
}

func (p *CommentUpdateDto) ParseUpdateToComment() db.UpdateCommentParams {
	arg := db.UpdateCommentParams{
		ID:      p.ID,
		Content: p.Content,
	}
	return arg
}

func (p *CommentUpdateLikesDto) ParseUpdateToCommentLike() db.UpdateCommentLikesParams {
	arg := db.UpdateCommentLikesParams{
		ID: p.ID,
		Likes: sql.NullInt64{
			Int64: p.Likes,
			Valid: true,
		},
	}
	return arg
}

func (p *CommentResponse) ParseFromCommentObject(result db.Comment) {
	p.ID = result.ID
	p.UserID = result.UserID
	p.PostID = result.PostID
	p.Content = result.Content
	p.Likes = result.Likes.Int64
	p.CreatedAt = time.Now()
}
