package cmd

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	echoSwagger "github.com/swaggo/echo-swagger"
	"time"
	"treads/infra"
)

func StartAPI(ctx context.Context, container *infra.ContainerDI) {
	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	go func() {
		for {
			select {
			case <-ctx.Done():
				if err := e.Shutdown(ctx); err != nil {
					panic(err)
				}
				return
			default:
				time.Sleep(1 * time.Second)
			}
		}
	}()

	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	post := e.Group("/post")
	post.GET("/list-all", container.HandlerPost.GetAllPosts)
	post.POST("/create", container.HandlerPost.CreatePost)
	post.PUT("/update", container.HandlerPost.UpdatePost)
	post.DELETE("/delete", container.HandlerPost.DeletePost)

	e.Logger.Fatal(e.Start(container.Config.ServerPort))
}
