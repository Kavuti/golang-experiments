package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// const (
// 	host     = "ec2-54-247-181-232.eu-west-1.compute.amazonaws.com"
// 	port     = 5432
// 	user     = "ypekdohkmvmcew"
// 	password = "8acf4f708538c87049c3ab7b34b1a57583753910c690f23f798db96c1313141c"
// 	dbname   = "db9078uu8gljrn"
// )

const (
	host     = "postgres"
	port     = 5432
	user     = "postres"
	password = "password"
	dbname   = "postgresgolang"
)

func main() {
	// configuration := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
	// 	host, port, user, password, dbname)
	db, err := sql.Open("postgres", "postgresql://postgres:password@172.24.0.2:5432/postgresgolang?sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("mammt")
		log.Fatalln(err)
	}
	defer db.Close()

}
