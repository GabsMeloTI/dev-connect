package repository

import (
	"context"
	"database/sql"
	db "treads/db/sqlc"
)

type CommentInterface interface {
	GetAllComments(context.Context, db.GetAllCommentsParams) ([]db.Comment, error)
	CreateComment(context.Context, db.CreateCommentParams) (db.Comment, error)
	UpdateComment(context.Context, db.UpdateCommentParams) (db.Comment, error)
	DeleteComment(context.Context, db.DeleteCommentParams) error
	IncrementLikes(ctx context.Context, commentID int64) (db.Comment, error)
	DecrementLikes(ctx context.Context, commentID int64) (db.Comment, error)
	GetCommentByID(ctx context.Context, commentID int64) (db.Comment, error)
}

type Comment struct {
	DBtx    db.DBTX
	Queries *db.Queries
	SqlConn *sql.DB
}

func NewComment(sqlDB *sql.DB) *Comment {
	q := db.New(sqlDB)
	return &Comment{
		DBtx:    sqlDB,
		Queries: q,
		SqlConn: sqlDB,
	}
}

func (r *Comment) GetAllComments(ctx context.Context, arg db.GetAllCommentsParams) ([]db.Comment, error) {
	return r.Queries.GetAllComments(ctx, arg)
}

func (r *Comment) CreateComment(ctx context.Context, arg db.CreateCommentParams) (db.Comment, error) {
	return r.Queries.CreateComment(ctx, arg)
}

func (r *Comment) UpdateComment(ctx context.Context, arg db.UpdateCommentParams) (db.Comment, error) {
	return r.Queries.UpdateComment(ctx, arg)
}

func (r *Comment) DeleteComment(ctx context.Context, arg db.DeleteCommentParams) error {
	return r.Queries.DeleteComment(ctx, arg)
}

func (r *Comment) IncrementLikes(ctx context.Context, commentID int64) (db.Comment, error) {
	return r.Queries.IncrementCommentLikes(ctx, commentID)
}

func (r *Comment) DecrementLikes(ctx context.Context, commentID int64) (db.Comment, error) {
	return r.Queries.DecrementCommentLikes(ctx, commentID)
}

func (r *Comment) GetCommentByID(ctx context.Context, commentID int64) (db.Comment, error) {
	return r.Queries.GetCommentByID(ctx, commentID)
}
