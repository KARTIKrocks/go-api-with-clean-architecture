package service

import (
	"errors"

	"github.com/KARTIKrocks/go-api-with-clean-architecture/models"
	"github.com/KARTIKrocks/go-api-with-clean-architecture/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo repository.PostRepository
)

type PostService interface {
	Validate(post *models.Post) error
	Create(post *models.Post) (*models.Post, error)
	FindAll() ([]primitive.M, error)
}

type service struct {
}

func NewPostService(repository repository.PostRepository) PostService {
	repo = repository
	return &service{}
}

func (s *service) Validate(post *models.Post) error {
	if post == nil {
		err := errors.New("the post must not be nil")
		return err
	}
	if post.Title == "" {
		err := errors.New("the post title must not be empty")
		return err
	}
	return nil
}

func (s *service) Create(post *models.Post) (*models.Post, error) {
	post.ID = primitive.NewObjectID()
	return repo.Save(post)
}

func (s *service) FindAll() ([]primitive.M, error) {
	return repo.FindAll()
}
