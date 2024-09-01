package main

import (
	"app/app/api"
	"app/config"
	"app/db"
)

func init() {
	config.InitEnvs()
	db.ConnectToDB()
}

func main() {
	api.RunServer()
}
