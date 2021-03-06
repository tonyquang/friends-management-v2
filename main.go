package main

import (
	"log"
	"net/http"
	"os"

	handlers "friends_management_v2/controller"
	"friends_management_v2/docs"
	"friends_management_v2/utils"
)

// @in header
// @name Authorization
func main() {
	db := utils.CreateConnection()
	r := handlers.Setup(db)
	docs.SwaggerInfo.Title = "Friends Management API"
	docs.SwaggerInfo.Description = "A Restful API for simple Friends Management application with GO, using gin-gonic/gin (A most popular HTTP framework) and gorm (The fantastic ORM library for Golang)"
	docs.SwaggerInfo.Version = "2.0"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}

	var port string = os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Println("Server started on: http://localhost:" + port)
	http.ListenAndServe(":"+port, r)
}
