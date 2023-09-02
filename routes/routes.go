package routes

import (
	"fmt"
	"github.com/KARTIKrocks/go-api-with-clean-architecture/controller/todo"
	router "github.com/KARTIKrocks/go-api-with-clean-architecture/http"
	"net/http"
)

func PostRoutes(router router.Router, postController todo.PostController) {
	router.GET("/healthCheck", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(response, "Up and running...")
	})

	router.POST("/todos", postController.AddTodo)
	router.GET("/todos/{id}", postController.GetTodo)
	router.GET("/todos", postController.GetTodos)
	router.PUT("/todos/{id}", postController.CompleteTodo)
	// router.PUT("/todos/update/{id}", postController.UpdateTodo)
	router.DELETE("/todos", postController.DeleteTodos)
	router.DELETE("/todos/{id}", postController.DeleteTodo)
}
