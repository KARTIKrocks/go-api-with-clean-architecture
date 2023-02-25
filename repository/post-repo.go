package repository

import (
	"github.com/KARTIKrocks/go-api-with-clean-architecture/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostRepository interface {
	Save(post *models.Todo) (*models.Todo, error)
	FindAll() ([]primitive.M, error)
	DeleteAll() (int64, error)
	DeleteOne(id string) (string, error)
	CompleteTodo(title string) error
	Get(id string) (*models.Todo, error)
	// Update(id string, todo *models.Todo) error
}
