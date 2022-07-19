package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tamagossi/bookstore/pkg/routes"
)

func main() {
	route := mux.NewRouter()
	routes.RegisterBookStoreRoutes(route)

	http.Handle("/", route)
	log.Fatal(http.ListenAndServe("localhost:8080", route))
}
