package main

import (
	"log"
	"net/http"

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
	docs.SwaggerInfo.Description = "Home Test API"
	docs.SwaggerInfo.Version = "2.0"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}

	// var port string = os.Getenv("PORT")
	// fmt.Println(port)
	log.Println("Server started on: http://localhost:3000")
	http.ListenAndServe(":3000", r)
}
