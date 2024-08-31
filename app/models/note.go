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

func CreateNote(note types.Note) error {
	query := "INSERT INTO notes (title, body) VALUES ($1, $2)"
	log.Print("pq: ", query, note.Title, note.Body)
	_, err := db.Db.Exec(query, note.Title, note.Body)

	if err != nil {
		return err
	}

	return nil
}

func GetAllUserNotes() {

}
