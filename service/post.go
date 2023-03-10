package service

import (
	"errors"
	"time"

	"github.com/KARTIKrocks/go-api-with-clean-architecture/models"
	"github.com/KARTIKrocks/go-api-with-clean-architecture/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostService interface {
	Validate(todo *models.Todo) error
	Create(todo *models.Todo) (*models.Todo, error)
	FindAll() ([]primitive.M, error)
	DeleteAll() (int64, error)
	DeleteOne(id string) (string, error)
	CompleteTodo(title string) error
	GetTodo(id string) (*models.Todo, error)
	// UpdateTodo(id string, todo *models.Todo) error
}

type service struct {
	repo repository.PostRepository
}

func NewPostService(repository repository.PostRepository) PostService {
	return &service{
		repo: repository,
	}
}

func (s *service) Validate(todo *models.Todo) error {
	if todo == nil {
		err := errors.New("the post must not be nil")
		return err
	}
	if todo.Title == "" {
		err := errors.New("the post title must not be empty")
		return err
	}
	return nil
}

func (s *service) Create(todo *models.Todo) (*models.Todo, error) {
	todo.ID = primitive.NewObjectID()
	currentTime := time.Now()
	todo.CreatedAt = currentTime
	todo.UpdatedAt = currentTime
	todo.Done = false
	return s.repo.Save(todo)
}

func (s *service) FindAll() ([]primitive.M, error) {
	return s.repo.FindAll()
}

func (s *service) DeleteAll() (int64, error) {
	return s.repo.DeleteAll()
}

func (s *service) DeleteOne(id string) (string, error) {
	return s.repo.DeleteOne(id)
}

func (s *service) GetTodo(id string) (*models.Todo, error) {
	return s.repo.Get(id)
}

func (s *service) CompleteTodo(title string) error {
	return s.repo.CompleteTodo(title)
}

// func (s *service) UpdateTodo(id string, todo *models.Todo) error {
// 	return s.repo.Update(id, todo)
// }
