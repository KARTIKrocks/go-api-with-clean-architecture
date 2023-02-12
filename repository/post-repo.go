package repository

import (
	"github.com/KARTIKrocks/go-api-with-clean-architecture/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostRepository interface {
	Save(post *models.Post) (*models.Post, error)
	FindAll() ([]primitive.M, error)
}
