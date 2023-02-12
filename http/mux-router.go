package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type muxRouter struct {
}

var (
	muxDispatcher = mux.NewRouter()
)

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (m *muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}

func (m *muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}

func (m *muxRouter) SERVE(port string) {
	fmt.Printf("Mux HTTP server running on port %v", port)
	log.Fatal(http.ListenAndServe(port, muxDispatcher))
}
