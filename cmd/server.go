package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"go-forum-api/app/handlers"
	"go-forum-api/app/repositories"
	repoImpl "go-forum-api/app/repositories/impl"
	"go-forum-api/app/usecases"
	ucImpl "go-forum-api/app/usecases/impl"
	"go-forum-api/utils/validator"
)

type Repositories struct {
	User    repositories.IUserRepository
	Forum   repositories.IForumRepository
	Thread  repositories.IThreadRepository
	Service repositories.IServiceRepository
	Post    repositories.IPostRepository
}

type UseCases struct {
	User    usecases.IUserUseCase
	Forum   usecases.IForumUseCase
	Thread  usecases.IThreadUseCase
	Service usecases.IServiceUseCase
	Post    usecases.IPostUseCase
}

type Handlers struct {
	User    *handlers.UserHandler
	Forum   *handlers.ForumHandler
	Thread  *handlers.ThreadHandler
	Service *handlers.ServiceHandler
	Post    *handlers.PostHandler
}

type Server struct {
	Settings     *Settings
	Repositories Repositories
	UseCases     UseCases
	Handlers     Handlers
}

func CreateServer() *Server {
	server := &Server{
		Settings:     LoadSettings(),
		Repositories: Repositories{},
		UseCases:     UseCases{},
		Handlers:     Handlers{},
	}
	_, err := validator.GetInstance()
	if err != nil {
		fmt.Printf("Can't create validator instance: %v", err)
		return nil
	}
	return server
}

func (server *Server) Run() {
	/* DataBase */
	//TODO: подкрутить конфиг
	config, err := pgxpool.ParseConfig(server.Settings.DSN)
	if err != nil {
		fmt.Printf("Can't parese DSN: %v\n", err)
		return
	}

	db, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		fmt.Printf("Can't create DB connection pool: %v", err)
		return
	}

	/* Repositories & UseCases*/
	server.Repositories.User = repoImpl.CreateUserRepository(db)
	server.UseCases.User = ucImpl.CreateUserUseCase(server.Repositories.User)

	server.Repositories.Thread = repoImpl.CreateThreadRepository(db)
	server.UseCases.Thread = ucImpl.CreateThreadUseCase(server.Repositories.Thread)

	server.Repositories.Forum = repoImpl.CreateForumRepository(db)
	server.UseCases.Forum = ucImpl.CreateForumUseCase(server.Repositories.Forum, server.Repositories.Thread)

	server.Repositories.Service = repoImpl.CreateServiceRepository(db)
	server.UseCases.Service = ucImpl.CreateServiceUseCase(server.Repositories.Service)

	server.Repositories.Post = repoImpl.CreatePostRepository(db)
	server.UseCases.Post = ucImpl.CreatePostUseCase(server.Repositories.Post)

	/* Server */
	gin.SetMode(server.Settings.MODE)
	router := gin.New()
	if server.Settings.MODE == "debug" {
		router.Use(gin.Logger())
	}
	router.Use(gin.Recovery())
	apiGroup := router.Group(server.Settings.Urls.Root)

	/*Handlers*/
	server.Handlers.User = handlers.CreateUserHandler(server.Settings.Urls.User, server.UseCases.User, apiGroup)
	server.Handlers.Forum = handlers.CreateForumHandler(server.Settings.Urls.Forum, server.UseCases.Forum, apiGroup)
	server.Handlers.Thread = handlers.CreateThreadHandler(server.Settings.Urls.Thread, server.UseCases.Thread, apiGroup)
	server.Handlers.Service = handlers.CreateServiceHandler(server.Settings.Urls.Service, server.UseCases.Service, apiGroup)
	server.Handlers.Post = handlers.CreatePostHandler(server.Settings.Urls.Post, server.UseCases.Post, apiGroup)

	err = router.Run(server.Settings.APIAddr)
	if err != nil {
		fmt.Printf("Can't start server: %v\n", err)
		return
	}

}
