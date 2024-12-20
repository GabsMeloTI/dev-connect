package cmd

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	echoSwagger "github.com/swaggo/echo-swagger"
	"time"
	"treads/infra"
)

func StartAPI(ctx context.Context, container *infra.ContainerDI) {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:3000",
			"https://devconnect-z3dw.vercel.app",
			"https://rzsxmvcdsg.us-east-1.awsapprunner.com",
		},
		AllowMethods: []string{
			echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS,
		},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
			"X-Requested-With",
			"Access-Control-Allow-Origin",
		},
		ExposeHeaders: []string{
			"Content-Length",
			"X-Request-ID",
		},
		AllowCredentials: true,
	}))

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
	post.DELETE("/delete/:id", container.HandlerPost.DeletePost)
	post.GET("/list-all/:id", container.HandlerPost.GetAllPostsByUser)

	user := e.Group("/user")
	user.GET("/list-all", container.HandlerUser.GetAllUsers)
	user.POST("/login", container.HandlerUser.LoginUser)
	user.POST("/create", container.HandlerUser.CreateUser)
	user.PUT("/update", container.HandlerUser.UpdateUser)
	user.PUT("/update/password", container.HandlerUser.UpdatePassword)
	user.PUT("/disable/:id", container.HandlerUser.DisableUser)
	user.DELETE("/delete/:id", container.HandlerUser.DeleteUser)

	comment := e.Group("/comment")
	comment.GET("/list-all", container.HandlerComment.GetAllComments)
	comment.POST("/create", container.HandlerComment.CreateComment)
	comment.PUT("/update/:id", container.HandlerComment.UpdateComment)
	comment.DELETE("/delete/:id", container.HandlerComment.DeleteComment)
	comment.PUT("/increment-like/:id", container.HandlerComment.IncrementLikes)
	comment.PUT("/decrement-like/:id", container.HandlerComment.DecrementLikes)

	e.Logger.Fatal(e.Start(container.Config.ServerPort))
}
