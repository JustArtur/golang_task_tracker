package models

import (
	"fmt"
	"golang_task_tracker/app/types"
	"golang_task_tracker/db"
	"log"
)

func GetNoteByID(ID int) (*types.Note, error) {
	note := new(types.Note)

	query := "SELECT * FROM notes WHERE \"id\" = $1"

	log.Print("pq: ", query, ID)
	rows, err := db.Db.Query(query, ID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&note.ID,
			&note.Title,
			&note.Body,
		)
		if err != nil {
			return nil, err
		}
	}

	if note.ID == 0 {
		return nil, fmt.Errorf("note not found")
	}

	return note, nil
}

func CreateNote(note *types.NotePayload) (*types.Note, error) {
	noteRecord := new(types.Note)

	query := "INSERT INTO notes (user_id, title, body) VALUES ($1, $2, $3) RETURNING id, user_id, title, body"
	log.Print("pq: ", query, note.UserID, note.Title, note.Body)
	rows, err := db.Db.Query(query, note.UserID, note.Title, note.Body)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&noteRecord.ID,
			&noteRecord.UserID,
			&noteRecord.Title,
			&noteRecord.Body,
		)
		if err != nil {
			return nil, err
		}
	}

	return noteRecord, nil
}

func GetAllUserNotes(userID int) ([]types.Note, error) {
	var notes []types.Note

	query := "SELECT * FROM notes WHERE \"user_id\" = $1"

	log.Print("pq: ", query, userID)
	rows, err := db.Db.Query(query, userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var note types.Note
		err = rows.Scan(
			&note.ID,
			&note.UserID,
			&note.Title,
			&note.Body,
		)
		if err != nil {
			return nil, err
		}

		notes = append(notes, note)
	}

	return notes, nil
}
