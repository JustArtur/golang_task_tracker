package models

import (
	"fmt"
	"log"
	"task_tracker_app/app/types"
	"task_tracker_app/db"
)

func GetUserByEmail(email string) (*types.User, error) {
	user := new(types.User)

	query := "SELECT * FROM users WHERE \"email\" = $1"

	log.Print("pq: ", query, email)
	rows, err := db.Db.Query(query, email)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
		)
		if err != nil {
			return nil, err
		}
	}

	if user.ID == 0 {
		return nil, fmt.Errorf("user not found, incorrect email or password")
	}

	return user, nil
}

func GetUserByID(ID int) (*types.User, error) {
	user := new(types.User)

	query := "SELECT * FROM users WHERE \"id\" = $1"

	log.Print("pq: ", query, ID)
	rows, err := db.Db.Query(query, ID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
		)
		if err != nil {
			return nil, err
		}
	}

	if user.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func CreateUser(user types.User) error {

	query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3)"
	log.Print("pq: ", query, user.Name, user.Email, user.Password)
	_, err := db.Db.Exec(query, user.Name, user.Email, user.Password)

	if err != nil {
		return err
	}

	return nil
}
