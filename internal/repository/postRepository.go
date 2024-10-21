package repository

import (
	"context"
	"database/sql"
	db "treads/db/sqlc"
)

type PostInterface interface {
	GetAllPosts(context.Context) ([]db.Post, error)
	CreatePost(context.Context, db.CreatePostParams) (db.Post, error)
	UpdatePost(context.Context, db.UpdatePostParams) (db.Post, error)
	DeletePost(context.Context, int32) error
}

type Post struct {
	DBtx    db.DBTX
	Queries *db.Queries
	SqlConn *sql.DB
}

func NewPost(sqlDB *sql.DB) *Post {
	q := db.New(sqlDB)
	return &Post{
		DBtx:    sqlDB,
		Queries: q,
		SqlConn: sqlDB,
	}
}

func (r *Post) GetAllPosts(ctx context.Context) ([]db.Post, error) {
	return r.Queries.GetAllPosts(ctx)
}

func (r *Post) CreatePost(ctx context.Context, arg db.CreatePostParams) (db.Post, error) {
	return r.Queries.CreatePost(ctx, arg)
}

func (r *Post) UpdatePost(ctx context.Context, arg db.UpdatePostParams) (db.Post, error) {
	return r.Queries.UpdatePost(ctx, arg)
}

func (r *Post) DeletePost(ctx context.Context, id int32) error {
	return r.Queries.DeletePost(ctx, id)
}
