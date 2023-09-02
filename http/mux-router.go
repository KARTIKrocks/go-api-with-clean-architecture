package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type muxRouter struct {
	muxRouter *mux.Router
}

func NewMuxRouter(router *mux.Router) Router {
	return &muxRouter{
		muxRouter: router,
	}
}

func (m *muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	m.muxRouter.HandleFunc(uri, f).Methods("GET")
}

func (m *muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	m.muxRouter.HandleFunc(uri, f).Methods("POST")
}

func (m *muxRouter) PUT(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	m.muxRouter.HandleFunc(uri, f).Methods("PUT")
}

func (m *muxRouter) PATCH(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	m.muxRouter.HandleFunc(uri, f).Methods("PATCH")
}

func (m *muxRouter) DELETE(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	m.muxRouter.HandleFunc(uri, f).Methods("DELETE")
}

func (m *muxRouter) SERVE(port string) {
	fmt.Printf("Mux HTTP server running on port %v", port)
	log.Fatal(http.ListenAndServe(port, m.muxRouter))
}
