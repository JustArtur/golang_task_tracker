package main

import (
	"golang_task_tracker/app/api"
	"golang_task_tracker/config"
	"golang_task_tracker/db"
)

func init() {
	config.InitEnvs()
	db.ConnectToDB()
}

func main() {
	api.RunServer()
}
