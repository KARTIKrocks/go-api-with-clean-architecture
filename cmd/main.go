package main

import (
	"github.com/KARTIKrocks/go-api-with-clean-architecture/controller/todo"
	"github.com/KARTIKrocks/go-api-with-clean-architecture/db"
	router "github.com/KARTIKrocks/go-api-with-clean-architecture/http"
	todo2 "github.com/KARTIKrocks/go-api-with-clean-architecture/repository/todo"
	"github.com/KARTIKrocks/go-api-with-clean-architecture/routes"
	"github.com/KARTIKrocks/go-api-with-clean-architecture/service"
	"github.com/gorilla/mux"
)

func main() {
	const port string = ":8080"

	dbClient, err := db.InitDB()
	if err != nil {
		panic(err)
	}

	muxRouter := mux.NewRouter()
	httpRouter := router.NewMuxRouter(muxRouter)

	PostRepository := todo2.NewMongoRepository(dbClient, "posts")
	postService := service.NewPostService(PostRepository)
	PostController := todo.NewPostController(postService, muxRouter)

	routes.PostRoutes(httpRouter, PostController)

	httpRouter.SERVE(port)
}
