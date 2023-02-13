package main

import (
	"fmt"
	"net/http"

	"github.com/KARTIKrocks/go-api-with-clean-architecture/controller"
	router "github.com/KARTIKrocks/go-api-with-clean-architecture/http"
	"github.com/KARTIKrocks/go-api-with-clean-architecture/repository"
	"github.com/KARTIKrocks/go-api-with-clean-architecture/service"
)

var (
	PostRepository repository.PostRepository = repository.NewMongoRepository()
	postService    service.PostService       = service.NewPostService(PostRepository)
	PostController controller.PostController = controller.NewPostController(postService)
	httpRouter     router.Router             = router.NewMuxRouter()
)

func main() {
	const port string = ":8080"
	httpRouter.GET("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(response, "Up and running...")
	})

	httpRouter.GET("/posts", PostController.GetPosts)
	httpRouter.POST("/posts", PostController.AddPost)
	httpRouter.SERVE(port)
}
