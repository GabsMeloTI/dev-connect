package repository

import (
	"context"
	"database/sql"
	db "treads/db/sqlc"
)

type CommentInterface interface {
	GetAllComments(context.Context) ([]db.Comment, error)
	CreateComment(context.Context, db.CreateCommentParams) (db.Comment, error)
	UpdateComment(context.Context, db.UpdateCommentParams) (db.Comment, error)
	DeleteComment(context.Context, int64) error
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

func (r *Comment) GetAllComments(ctx context.Context) ([]db.Comment, error) {
	return r.Queries.GetAllComments(ctx)
}

func (r *Comment) CreateComment(ctx context.Context, arg db.CreateCommentParams) (db.Comment, error) {
	return r.Queries.CreateComment(ctx, arg)
}

func (r *Comment) UpdateComment(ctx context.Context, arg db.UpdateCommentParams) (db.Comment, error) {
	return r.Queries.UpdateComment(ctx, arg)
}

func (r *Comment) DeleteComment(ctx context.Context, id int64) error {
	return r.Queries.DeleteComment(ctx, id)
}
