package main

import (
	"goApiFramework/router"
	"net/http"
)

func main() {
	routers := router.NewRouter()
	http.ListenAndServe(":80", routers)
}
