package main

import (
	"task_tracker_app/app/api"
	"task_tracker_app/config"
	"task_tracker_app/db"
)

func init() {
	config.InitEnvs()
	db.ConnectToDB()
}

func main() {
	api.RunServer()
}
