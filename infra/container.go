package infra

import (
	"database/sql"
	"treads/infra/database"
	"treads/infra/database/db_postgresql"
	"treads/internal/handler"
	"treads/internal/repository"
	"treads/internal/service"
	"treads/pkg/aws"
)

type ContainerDI struct {
	Config            Config
	ContainerURL      containerURL
	ConnDB            *sql.DB
	SESClient         *aws.SESClient
	RepositoryPost    *repository.Post
	RepositoryUser    *repository.User
	RepositoryComment *repository.Comment
	ServiceComment    *service.Comment
	HandlerComment    *handler.Comment
	ServicePost       *service.Post
	ServiceUser       *service.User
	HandlerPost       *handler.Post
	HandlerUser       *handler.User
}

type containerURL struct {
	UrlAnticipation              string
	UrlAnticipationMade          string
	UrlLogin                     string
	UrlAnticipationFinancial     string
	UrlAnticipationPurchaseOrder string
}

func NewContainerDI(config Config) *ContainerDI {
	container := &ContainerDI{Config: config}

	container.db()
	container.aws()
	container.buildRepository()
	container.buildService()
	container.buildHandler()

	return container
}

func (c *ContainerDI) db() {
	dbConfig := database.Config{
		Host:        c.Config.DBHost,
		Port:        c.Config.DBPort,
		User:        c.Config.DBUser,
		Password:    c.Config.DBPassword,
		Database:    c.Config.DBDatabase,
		SSLMode:     c.Config.DBSSLMode,
		Driver:      c.Config.DBDriver,
		Environment: c.Config.Environment,
	}
	c.ConnDB = db_postgresql.ConnDB(&dbConfig, true)
}

func (c *ContainerDI) aws() {
	region := c.Config.AWSRegion
	c.SESClient = aws.NewSESClient(region)
}

func (c *ContainerDI) buildRepository() {
	c.RepositoryUser = repository.NewUser(c.ConnDB)
	c.RepositoryPost = repository.NewPost(c.ConnDB)
	c.RepositoryComment = repository.NewComment(c.ConnDB)

}

func (c *ContainerDI) buildService() {
	c.ServiceUser = service.NewUser(c.RepositoryUser, c.SESClient)
	c.ServicePost = service.NewPost(c.RepositoryPost)
	c.ServiceComment = service.NewComment(c.RepositoryComment)
}

func (c *ContainerDI) buildHandler() {
	c.HandlerUser = handler.NewUser(c.ServiceUser)
	c.HandlerPost = handler.NewPost(c.ServicePost)
	c.HandlerComment = handler.NewComment(c.ServiceComment)
}
