package api

import (
	"github.com/gorilla/mux"
	"golang_task_tracker/app/controllers"
	"log"
	"net/http"
)

func RunServer() {
	server := http.Server{
		Handler: newRoute(),
		Addr:    ":8000",
	}

	log.Printf("Starting up server on port %s", server.Addr)
	log.Fatal(server.ListenAndServe())
}

func newRoute() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/login", controllers.HandleLogin).Methods("POST")
	router.HandleFunc("/register", controllers.HandleRegister).Methods("POST")

	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	apiRouter.Use()
	return router
}
