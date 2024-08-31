package controllers

import (
	"golang_task_tracker/app/types"
	"log"
	"net/http"
)

var note types.UserPayload

func Create(w http.ResponseWriter, r *http.Request) {
	log.Printf("Started %s %s", r.Method, r.RequestURI)
	defer log.Printf("Completed %s %s", r.Method, r.RequestURI)

	err := ParseRequest(r, &note)
	if err != nil {
		log.Println(err)
		SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	log.Printf("Started %s %s", r.Method, r.RequestURI)
	defer log.Printf("Completed %s %s", r.Method, r.RequestURI)

	err := ParseRequest(r, &note)
	if err != nil {
		log.Println(err)
		SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}
}
