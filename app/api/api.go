package api

import (
	"github.com/gorilla/mux"
	"golang_task_tracker/app/controllers/user"
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

	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	apiRouter.HandleFunc("/login", user.HandleLogin).Methods("POST")
	apiRouter.HandleFunc("/register", user.HandleRegister).Methods("POST")

	return router
}
