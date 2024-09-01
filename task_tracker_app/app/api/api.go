package api

import (
	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"task_tracker_app/app/controllers"
	"task_tracker_app/app/services/auth"
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

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/register", controllers.Register).Methods("POST")

	apiRouter := router.PathPrefix("/api/v1").Subrouter()
	apiRouter.Use(auth.JWTMiddleware)
	apiRouter.HandleFunc("/notes", controllers.Create).Methods("POST")
	apiRouter.HandleFunc("/notes", controllers.Index).Methods("GET")

	return router
}
