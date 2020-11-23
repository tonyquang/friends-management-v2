package utils

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
<<<<<<< HEAD
	host     = "localhost"
=======
	host = "localhost" // using for running on localhost
	//host     = "db"        // using for running on docker
>>>>>>> DONE-API
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "FriendsManagement"
)

func CreateConnection() *gorm.DB {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{SkipDefaultTransaction: true})

<<<<<<< HEAD
	// dsn := "user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable TimeZone=Asia/ho_chi_minh"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

=======
>>>>>>> DONE-API
	if err != nil {
		panic(err)
	}
	fmt.Println("Connect to database successfully")
	return db
}
