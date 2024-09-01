package controllers

import (
	"golang_task_tracker/app/helpers"
	"golang_task_tracker/app/models"
	"golang_task_tracker/app/services/notes"
	"golang_task_tracker/app/types"
	"log"
	"net/http"
)

var note types.NotePayload

func Create(w http.ResponseWriter, r *http.Request) {
	log.Printf("Started %s %s", r.Method, r.RequestURI)
	defer log.Printf("Completed %s %s", r.Method, r.RequestURI)

	note.UserID = helpers.GetUserIDFromContext(r)

	err := helpers.ParseRequest(r, &note)
	if err != nil {
		log.Println(err)
		helpers.SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	notes.Correct(&note)

	noteRecord, err := models.CreateNote(&note)
	if err != nil {
		log.Println(err)
		helpers.SendErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	helpers.SendResponse(w, http.StatusCreated, noteRecord)
}

func Index(w http.ResponseWriter, r *http.Request) {
	log.Printf("Started %s %s", r.Method, r.RequestURI)
	defer log.Printf("Completed %s %s", r.Method, r.RequestURI)

	notes, err := models.GetAllUserNotes(helpers.GetUserIDFromContext(r))
	if err != nil {
		log.Printf("failed to get notes from DB %v", err)
		helpers.SendErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	helpers.SendResponse(w, http.StatusOK, notes)
}
