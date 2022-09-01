package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"goApiFramework/controller"
	"goApiFramework/model"
	"net/http"
)

var routes []model.Route

func init() {
	fmt.Println("Route Init")
	register("GET", "/hello", controller.Hello, nil)
	fmt.Printf("%+v", routes)
}

func NewRouter() http.Handler {
	r := mux.NewRouter()
	for _, route := range routes {
		r.Methods(route.Method).
			Path(route.Pattern).
			Handler(route.Handler)
		if route.Middleware != nil {
			r.Use(route.Middleware)
		}
	}
	handler := cors.Default().Handler(r)
	return handler
}

func register(method, pattern string, handler http.HandlerFunc, middleware mux.MiddlewareFunc) {
	routes = append(routes, model.Route{Method: method, Pattern: pattern, Handler: handler, Middleware: middleware})
}
