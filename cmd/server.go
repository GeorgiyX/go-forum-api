package main

import (
	"github.com/gin-gonic/gin"
	"go-forum-api/app/handlers"
	"go-forum-api/app/repositories"
	"go-forum-api/app/usecases"
)

type Repositories struct {
	User repositories.IUserRepository
}

type UseCases struct {
	User usecases.IUserUseCase
}

type Handlers struct {
	User handlers.UserHandler
}

type Server struct {
	Settings     Settings
	Repositories Repositories
	UseCases     UseCases
	Handlers     Handlers
}

func CreateServer() *Server {
	server := &Server{
		Settings:     Settings{},
		Repositories: Repositories{},
		UseCases:     UseCases{},
		Handlers:     Handlers{},
	}
}

func (server *Server) Run() {

	// Creates a router without any middleware by default
	router := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())
	apiGroup := router.Group(server.Settings.Urls.Root)

}
