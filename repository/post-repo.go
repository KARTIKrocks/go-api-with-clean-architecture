package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/KARTIKrocks/go-api-with-clean-architecture/config"
	"github.com/KARTIKrocks/go-api-with-clean-architecture/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PostRepository interface {
	Save(post *models.Post) (*models.Post, error)
	FindAll() ([]primitive.M, error)
}

type repo struct{}

func NewPostRepository() PostRepository {
	return &repo{}
}

// MOST IMPORTANT
var (
	collection       *mongo.Collection
	conf             = config.NewConfig()
	connectionString = conf.MongoURI()
	ctx              = context.Background()
)

const (
	dbName  = "demoDB"
	colName = "posts"
)

func init() {
	// client option
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongodb
	client, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		log.Fatalf("failed connecting client: %v", err)
	}

	collection = client.Database(dbName).Collection(colName)
	fmt.Println("mongo db connection success")
}

func (r *repo) Save(post *models.Post) (*models.Post, error) {

	post.ID = primitive.NewObjectID()
	inserted, err := collection.InsertOne(context.Background(), post)

	if err != nil {
		log.Fatalf("failed adding a new post: %v", err)
		return nil, err
	}
	fmt.Println("Inserted 1 post in db with id: ", inserted.InsertedID)
	return post, nil
}

func (r *repo) FindAll() ([]primitive.M, error) {

	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatalf("problem occured: %v", err)
	}

	var posts []primitive.M
	for cursor.Next(ctx) {
		var post bson.M
		err := cursor.Decode(&post)
		if err != nil {
			log.Fatal(err)
		}
		posts = append(posts, post)
	}
	defer cursor.Close(context.Background())

	return posts, nil
}
