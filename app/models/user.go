package models

import (
	"fmt"
	"golang_task_tracker/app/types"
	"golang_task_tracker/db"
	"log"
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

func CreateUser(user types.User) error {

	query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3)"
	log.Print("pq: ", query, user.Name, user.Email, user.Password)
	_, err := db.Db.Exec(query, user.Name, user.Email, user.Password)

	if err != nil {
		return err
	}

	return nil
}
