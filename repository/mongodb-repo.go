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

type repo struct {
	collection *mongo.Collection
}

var (
	collection *mongo.Collection
)

const (
	dbName  = "demoDB"
	colName = "posts"
)

func InitDB() *mongo.Collection {
	conf, err := config.LoadConfig("./")
	if err != nil {
		log.Fatalf("Error while reading config file: %s", err)
	}
	connectionString := conf.MongoURI()
	ctx := context.Background()

	// client option
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongodb
	client, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		log.Fatalf("failed connecting client: %v", err)
	}

	collection = client.Database(dbName).Collection(colName)
	fmt.Println("mongo db connection success")
	return collection
}

// NewMongoRepository creates a new repo
func NewMongoRepository(coll *mongo.Collection) PostRepository {
	return &repo{collection: coll}
}

// MOST IMPORTANT
var (
	ctx = context.Background()
)

func (r *repo) Save(post *models.Post) (*models.Post, error) {
	inserted, err := r.collection.InsertOne(ctx, post)

	if err != nil {
		log.Fatalf("failed adding a new post: %v", err)
		return nil, err
	}
	fmt.Println("Inserted 1 post in db with id: ", inserted.InsertedID)
	return post, nil
}

func (r *repo) FindAll() ([]primitive.M, error) {
	cursor, err := r.collection.Find(ctx, bson.D{{}})
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
	defer cursor.Close(ctx)

	return posts, nil
}
