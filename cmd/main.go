package main

import (
	"fmt"
	"github.com/KARTIKrocks/go-api-with-clean-architecture/controller/todo"
	"github.com/KARTIKrocks/go-api-with-clean-architecture/db"
	router "github.com/KARTIKrocks/go-api-with-clean-architecture/http"
	todo2 "github.com/KARTIKrocks/go-api-with-clean-architecture/repository/todo"
	"github.com/KARTIKrocks/go-api-with-clean-architecture/routes"
	"github.com/KARTIKrocks/go-api-with-clean-architecture/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	const port string = ":8080"

	dbClient, err := db.InitDB()
	if err != nil {
		panic(err)
	}

	muxRouter := mux.NewRouter()
	muxRouter.Use(LoggingMiddleware)

	httpRouter := router.NewMuxRouter(muxRouter)

	PostRepository := todo2.NewMongoRepository(dbClient, "posts")
	postService := service.NewPostService(PostRepository)
	PostController := todo.NewPostController(postService, muxRouter)

	routes.PostRoutes(httpRouter, PostController)

	httpRouter.SERVE(port)
}

// LoggingMiddleware is a middleware function that logs incoming requests.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Log the request details
		log.Printf(
			"Method: %s | URL: %s | RemoteAddr: %s",
			r.Method,
			r.URL.String(),
			r.RemoteAddr,
		)

		// Call the next handler in the chain
		next.ServeHTTP(w, r)

		// Calculate and log the request duration
		elapsed := time.Since(start)
		log.Printf("Request processed in %s", elapsed)
		fmt.Println()
	})
}
