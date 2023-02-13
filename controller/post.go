package controller

import (
	"encoding/json"
	"net/http"

	"github.com/KARTIKrocks/go-api-with-clean-architecture/errors"
	"github.com/KARTIKrocks/go-api-with-clean-architecture/models"
	"github.com/KARTIKrocks/go-api-with-clean-architecture/service"
)

var (
	postService service.PostService
)

type controller struct {
}

type PostController interface {
	GetPosts(resp http.ResponseWriter, req *http.Request)
	AddPost(resp http.ResponseWriter, req *http.Request)
}

func NewPostController(service service.PostService) PostController {
	postService = service
	return &controller{}
}

func (c *controller) GetPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	posts, err := postService.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode((errors.ServiceError{Message: "Error getting the posts"}))
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(posts)
}

func (c *controller) AddPost(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	var post models.Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode((errors.ServiceError{Message: "Error marshalling the request"}))
		return
	}

	err1 := postService.Validate(&post)
	if err1 != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode((errors.ServiceError{Message: err1.Error()}))
		return
	}
	result, err2 := postService.Create(&post)
	if err2 != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode((errors.ServiceError{Message: "Error saving the post"}))
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(result)
}
