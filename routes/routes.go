package routes

import (
	"log"
	"net/http"

	"github.com/Samaraferreira/go-api-rest/controllers"
	middleware "github.com/Samaraferreira/go-api-rest/middlewares"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	router := mux.NewRouter()
	router.Use(middleware.ContentTypeMiddleware)
	router.HandleFunc("/api/users", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/api/users", controllers.AddNewUser).Methods("POST")
	router.HandleFunc("/api/users/{id}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/api/users/{id}", controllers.DeleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8081", router))
}
