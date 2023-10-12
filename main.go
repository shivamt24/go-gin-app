package main

import (
	api "shivamthabe.me/go-gin-app/api"
	"shivamthabe.me/go-gin-app/database"
)

func init() {
	database.NewPostgreSQLClient()
}
func main() {
	r := api.SetupRouter()
	r.Run(":5050")
}
