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
	User  repositories.IUserRepository
	Forum repositories.IForumRepository
}

type UseCases struct {
	User  usecases.IUserUseCase
	Forum usecases.IForumUseCase
}

type Handlers struct {
	User  *handlers.UserHandler
	Forum *handlers.ForumHandler
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
	server.Repositories.Forum = repoImpl.CreateForumRepository(db)
	server.UseCases.Forum = ucImpl.CreateForumUseCase(server.Repositories.Forum)

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

	err = router.Run(server.Settings.APIAddr)
	if err != nil {
		fmt.Printf("Can't start server: %v\n", err)
		return
	}

}
