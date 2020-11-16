package main

import (
	handlers "friends_management_v2/controller"
	"friends_management_v2/utils"
	"log"
	"net/http"
)

func main() {
	db := utils.CreateConnection()
	r := handlers.Setup(db)
	// var port string = os.Getenv("PORT")
	// fmt.Println(port)
	log.Println("Server started on: http://localhost:3000")
	http.ListenAndServe(":3000", r)
}