package main

import (
	"log"
	"net/http"

	"github.com/aaryanraj/ebook-store-api/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rs/cors"
)

func handelHttpRequests() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "DELETE", "POST", "PUT"},
	})
	handler := c.Handler(r)
	log.Fatal(http.ListenAndServe("localhost:9010", handler))
	http.Handle("/", r)
}

func main() {
	handelHttpRequests()
}
