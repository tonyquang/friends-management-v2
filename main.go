package main

import (
	"friends_management_v2/cmd/handlers"
	"friends_management_v2/utils"
	"log"
	"net/http"
)

func main() {
	db := utils.CreateConnection()
	r := handlers.Setup(db)
	log.Println("Server started on: http://localhost:3000")
	http.ListenAndServe(":3000", r)
}
