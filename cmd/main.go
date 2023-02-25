package main

import (
	"fmt"
	"net/http"

	"github.com/KARTIKrocks/go-api-with-clean-architecture/controller"
	router "github.com/KARTIKrocks/go-api-with-clean-architecture/http"
	"github.com/KARTIKrocks/go-api-with-clean-architecture/repository"
	"github.com/KARTIKrocks/go-api-with-clean-architecture/service"
	"github.com/gorilla/mux"
)

var (
	muxRouter                                = mux.NewRouter()
	httpRouter     router.Router             = router.NewMuxRouter(muxRouter)
	PostRepository repository.PostRepository = repository.NewMongoRepository(repository.InitMongoDB())
	postService    service.PostService       = service.NewPostService(PostRepository)
	PostController controller.PostController = controller.NewPostController(postService, muxRouter)
)

func main() {
	const port string = ":8080"
	httpRouter.GET("/healthCheck", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(response, "Up and running...")
	})

	httpRouter.POST("/todos", PostController.AddTodo)
	httpRouter.GET("/todos/{id}", PostController.GetTodo)
	httpRouter.GET("/todos", PostController.GetTodos)
	httpRouter.PUT("/todos/{id}", PostController.CompleteTodo)
	// httpRouter.PUT("/todos/update/{id}", PostController.UpdateTodo)
	httpRouter.DELETE("/todos", PostController.DeleteTodos)
	httpRouter.DELETE("/todos/{id}", PostController.DeleteTodo)

	httpRouter.SERVE(port)
}
