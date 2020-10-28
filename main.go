package main

import (
	"friends_management_v2/utils"
)

func main() {
	db := utils.CreateConnection()

	// log.Println("Successfully connected!")

	// r := handlers.API(db)
	// log.Println("Server started on: http://localhost:3000")
	// http.ListenAndServe(":3000", r)
}
