package todo

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/KARTIKrocks/go-api-with-clean-architecture/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type repo struct {
	collection *mongo.Collection
}

const (
	dbName = "demoDB"
)

var (
	ctx = context.Background()
)

// NewMongoRepository creates a new repo
func NewMongoRepository(client *mongo.Client, collectionName string) PostRepository {
	collection := client.Database(dbName).Collection(collectionName)
	return &repo{collection: collection}
}

func (r *repo) CreateNew(todo *models.Todo) *models.Todo {
	todo.ID = primitive.NewObjectID()
	currentTime := time.Now()
	todo.CreatedAt = currentTime
	todo.UpdatedAt = currentTime
	todo.Done = false

	return todo
}

func (r *repo) Save(todo *models.Todo) (*models.Todo, error) {
	inserted, err := r.collection.InsertOne(ctx, todo)

	if err != nil {
		log.Fatalf("failed adding a new post: %v", err)
		return nil, err
	}
	fmt.Printf("Inserted 1 todo in db\nid: %s\ntitle: %s ", inserted.InsertedID, todo.Title)
	return todo, nil
}

func (r *repo) Get(id string) (*models.Todo, error) {
	fmt.Println("Get called")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var todo *models.Todo
	filter := bson.M{"_id": objId}
	err = r.collection.FindOne(ctx, filter).Decode(&todo)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("no document with that Id exists")
	}
	fmt.Println(todo)
	return todo, nil
}

func (r *repo) FindAll() ([]primitive.M, error) {
	cursor, err := r.collection.Find(ctx, bson.D{{}})
	if err != nil {
		log.Fatalf("problem occured: %v", err)
	}

	var todos []primitive.M
	for cursor.Next(ctx) {
		var todo bson.M
		err := cursor.Decode(&todo)
		if err != nil {
			log.Fatal(err)
		}
		todos = append(todos, todo)
	}
	defer cursor.Close(ctx)

	return todos, nil
}

// delete all records from mongodb
func (r *repo) DeleteOne(id string) (string, error) {
	fmt.Println("deleteOne called")
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objId}
	deleteCount, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return "", err
	}
	fmt.Println("todo got delete with delete count: ", deleteCount)
	return id, nil
}

// delete all records from mongodb
func (r *repo) DeleteAll() (int64, error) {
	fmt.Println("deleteAll entered")
	filter := bson.D{{}}
	deleteResult, err := r.collection.DeleteMany(ctx, filter, nil)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	fmt.Printf("Number of todos delete: %v", deleteResult.DeletedCount)
	return deleteResult.DeletedCount, nil
}

// CompleteTodo sets 'Done' feild as true
func (r *repo) CompleteTodo(id string) error {
	fmt.Println("CompletedTodo entered")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objId}
	update := bson.M{"done": true, "updated_at": time.Now()}

	result, err := r.collection.UpdateOne(ctx, filter, bson.M{"$set": update})
	if err != nil {
		return err
	}
	fmt.Println("modified count: ", result.ModifiedCount)
	return nil
}

// func (r *repo) Update(id string, todo *models.Todo)error  {
// 	fmt.Println("Update called")
// 	objId, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return err
// 	}

// 	filter := bson.M{"_id": objId}
// 	update := bson.M{"title": todo.Title, "description": todo.Description, "updated_at":time.Now()}

// 	result, err := r.collection.UpdateOne(ctx, filter, bson.M{"$set": update})
// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}
// 	fmt.Println("modified count: ", result.ModifiedCount)
// 	return nil

// }
