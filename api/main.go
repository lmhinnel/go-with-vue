package main

import (
	"time"

	"github.com/kansuke231/go-with-vue/api/database"
)

func main() {

	println("Making the connection to the PostgreSQL instance.....")
	db_connection_string := "host=db port=5432 dbname=postgres user=docker password=docker sslmode=disable"
	time.Sleep(5 * time.Second)
	db, err := database.NewDB(db_connection_string)
	if err != nil {
		println(err.Error())
	}

	println("Connected!")
	db.IsAlive()

}
