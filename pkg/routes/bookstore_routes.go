package routes

import (
	"github.com/aaryanraj/ebook-store-api/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {

	//login
	router.HandleFunc("/login", controllers.Login).Methods("POST")

	//books
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/{id}", controllers.GetBookByID).Methods("GET")
	router.HandleFunc("/book/", controllers.ListBooks).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("DELETE")

	//user
	router.HandleFunc("/signup/", controllers.CreateUser).Methods("POST") //create_user
	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")
}
