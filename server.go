package main

import (
	"fmt"
	"net/http"

	"github.com/KARTIKrocks/go-api-with-clean-architecture/controller"
	router "github.com/KARTIKrocks/go-api-with-clean-architecture/http"
)

var (
	httpRouter     router.Router             = router.NewMuxRouter()
	PostController controller.PostController = controller.NewPostController()
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
