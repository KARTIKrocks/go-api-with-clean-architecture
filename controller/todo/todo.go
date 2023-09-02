package todo

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/KARTIKrocks/go-api-with-clean-architecture/errors"
	"github.com/KARTIKrocks/go-api-with-clean-architecture/models"
	"github.com/KARTIKrocks/go-api-with-clean-architecture/response"
	"github.com/KARTIKrocks/go-api-with-clean-architecture/service"
	"github.com/gorilla/mux"
)

type controller struct {
	postService service.PostService
	router      *mux.Router
}

// NewPostController returns a new controller
func NewPostController(service service.PostService, router *mux.Router) PostController {
	return &controller{
		postService: service,
		router:      router,
	}
}

func (c *controller) GetTodos(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Access-Control-Allow-Methods", "GET")

	posts, err := c.postService.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode((errors.ServiceError{Message: "Error getting the posts"}))
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(posts)
}

func (c *controller) AddTodo(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Access-Control-Allow-Methods", "POST")

	var res response.UserResponse
	var post models.Todo
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		res = response.UserResponse{Status: http.StatusBadRequest, Message: "error marshaling the request", Data: map[string]interface{}{"data": err.Error()}}
		json.NewEncoder(resp).Encode(res)
		return
	}

	err1 := c.postService.Validate(&post)
	if err1 != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode((errors.ServiceError{Message: err1.Error()}))
		return
	}
	result, err2 := c.postService.Create(&post)
	if err2 != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode((errors.ServiceError{Message: "Error saving the post"}))
		return
	}
	res = response.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}}
	json.NewEncoder(resp).Encode(res)
}

func (c *controller) GetTodo(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Access-Control-Allow-Methods", "GET")

	params := mux.Vars(req)
	todo, err := c.postService.GetTodo(params["id"])
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode((errors.ServiceError{Message: "Error getting the todo"}))
	}
	json.NewEncoder(resp).Encode(todo)
}

// func (c *controller) UpdateTodo(resp http.ResponseWriter, req *http.Request) {
// 	resp.Header().Set("Content-Type", "application/json")
// 	resp.Header().Set("Access-Control-Allow-Methods", "PUT")

// 	params := mux.Vars(req)

// 	var res response.UserResponse
// 	var post models.Todo

// 	err := json.NewDecoder(req.Body).Decode(&post)
// 	if err != nil {
// 		resp.WriteHeader(http.StatusBadRequest)
// 		res = response.UserResponse{Status: http.StatusBadRequest, Message: "error marshaling the request", Data: map[string]interface{}{"data": err.Error()}}
// 		json.NewEncoder(resp).Encode(res)
// 		return
// 	}
// 	err1 := c.postService.Validate(&post)
// 	if err1 != nil {
// 		resp.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(resp).Encode((errors.ServiceError{Message: err1.Error()}))
// 		return
// 	}
// 	err = c.postService.UpdateTodo(params["id"], &post)
// 	if err != nil {
// 		resp.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(resp).Encode((errors.ServiceError{Message: "Error updating the todo"}))
// 	}
// 	json.NewEncoder(resp).Encode(params["id"])
// }

func (c *controller) CompleteTodo(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Access-Control-Allow-Methods", "POST")

	params := mux.Vars(req)
	err := c.postService.CompleteTodo(params["id"])
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode((errors.ServiceError{Message: "Error completing the todo"}))
	}
	json.NewEncoder(resp).Encode(params["id"])
}

func (c *controller) DeleteTodo(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Access-Control-Allow-Methods", "DELETE")

	params := mux.Vars(req)
	_, err := c.postService.DeleteOne(params["id"])
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode((errors.ServiceError{Message: "Error deleting the todo"}))
	}
	json.NewEncoder(resp).Encode(params["id"])
}

func (c *controller) DeleteTodos(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Access-Control-Allow-Methods", "DELETE")

	numberOfDeletedTodos, err := c.postService.DeleteAll()
	fmt.Println("deleted number", numberOfDeletedTodos)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode((errors.ServiceError{Message: "Error deleting the todos"}))
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(numberOfDeletedTodos)
}
