package router

import (
	// "fmt"
	// "log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ginRouter struct {
	ginRouter *gin.Engine
}

func NewGinRouter(router *gin.Engine) Router {
	return &ginRouter{
		ginRouter: router,
	}
}

func (m *ginRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	// ginDispatcher.GET(uri, f)
}

func (m *ginRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	// ginDispatcher.Post(uri, f)
}

func (m *ginRouter) PUT(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	// ginDispatcher.Put(uri, f)
}

func (m *ginRouter) PATCH(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	// ginDispatcher.Patch(uri, f)
}

func (m *ginRouter) DELETE(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	// ginDispatcher.Delete(uri, f)
}

func (m *ginRouter) SERVE(port string) {
	// fmt.Printf("CHi HTTP server running on port %v", port)
	// log.Fatal(http.ListenAndServe(port, chiDispatcher))
}
