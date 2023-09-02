package todo

import "net/http"

type PostController interface {
	GetTodos(resp http.ResponseWriter, req *http.Request)
	AddTodo(resp http.ResponseWriter, req *http.Request)
	GetTodo(resp http.ResponseWriter, req *http.Request)
	// UpdateTodo(resp http.ResponseWriter, req *http.Request)
	CompleteTodo(resp http.ResponseWriter, req *http.Request)
	DeleteTodo(resp http.ResponseWriter, req *http.Request)
	DeleteTodos(resp http.ResponseWriter, req *http.Request)
}
