package api

import (
	"app/app/controllers"
	"app/app/services/auth"
	"github.com/gorilla/mux"
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

	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/register", controllers.Register).Methods("POST")

	apiRouter := router.PathPrefix("/api/v1").Subrouter()
	apiRouter.Use(auth.JWTMiddleware)
	apiRouter.HandleFunc("/notes", controllers.Create).Methods("POST")
	apiRouter.HandleFunc("/notes", controllers.Index).Methods("GET")

	return router
}
