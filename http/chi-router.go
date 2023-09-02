package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type chiRouter struct {
	chiRouter *chi.Mux
}

func NewChiRouter(router *chi.Mux) Router {
	return &chiRouter{
		chiRouter: router,
	}
}

func (m *chiRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	m.chiRouter.Get(uri, f)
}

func (m *chiRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	m.chiRouter.Post(uri, f)
}

func (m *chiRouter) PUT(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	m.chiRouter.Put(uri, f)
}

func (m *chiRouter) PATCH(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	m.chiRouter.Patch(uri, f)
}

func (m *chiRouter) DELETE(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	m.chiRouter.Delete(uri, f)
}

func (m *chiRouter) SERVE(port string) {
	fmt.Printf("CHi HTTP server running on port %v", port)
	log.Fatal(http.ListenAndServe(port, m.chiRouter))
}
