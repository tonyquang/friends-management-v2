package main

import (
	handlers "friends_management_v2/controller"
	"friends_management_v2/utils"
	"log"
	"net/http"
	"os"
)
	
func main() {
	db := utils.CreateConnection()
	r := handlers.Setup(db)

	log.Println("Server started on: http://localhost:3000")
	http.ListenAndServe(os.Getenv("PORT"), r)
}
