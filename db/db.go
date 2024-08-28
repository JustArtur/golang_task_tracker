package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"golang_task_tracker/config"
	"log"
)

var Db *sql.DB

func ConnectToDB() *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Envs.DbHost, config.Envs.DbPort, config.Envs.DbUser, config.Envs.DbPass, config.Envs.DbName, config.Envs.DbSSLMode)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalf(err.Error())
	}

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("DataBase Successfully connected!")

	return db
}
