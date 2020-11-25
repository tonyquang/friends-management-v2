package utils

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	//host = "localhost" // using for running on localhost
	host     = "db" // using for running on docker
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "friend-mgmt"
)

func CreateConnection() *gorm.DB {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{SkipDefaultTransaction: true})

	if err != nil {
		panic(err)
	}
	fmt.Println("Connect to database successfully")
	return db
}
