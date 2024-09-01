package controllers

import (
	"log"
	"net/http"
	"task_tracker_app/app/helpers"
	"task_tracker_app/app/models"
	"task_tracker_app/app/services/notes"
	"task_tracker_app/app/types"
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
