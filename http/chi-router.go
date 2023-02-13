package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type chiRouter struct {
}

var (
	chiDispatcher = chi.NewRouter()
)

func NewChiRouter() Router {
	return &chiRouter{}
}

func (m *chiRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Get(uri, f)
}

func (m *chiRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Post(uri, f)
}

func (m *chiRouter) PUT(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Put(uri, f)
}

func (m *chiRouter) PATCH(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Patch(uri, f)
}

func (m *chiRouter) DELETE(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Delete(uri, f)
}

func (m *chiRouter) SERVE(port string) {
	fmt.Printf("CHi HTTP server running on port %v", port)
	log.Fatal(http.ListenAndServe(port, chiDispatcher))
}
