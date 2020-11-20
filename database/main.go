package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type User struct {
	Name     string
	Surname  string
	Username string
	Password string
}

var schema = `
CREATE DATABASE IF NOT EXISTS authenticator;

CREATE SCHEMA IF NOT EXISTS

CREATE TABLE IF NOT EXISTS user (
	id bigint primary key,
	name varchar(50),
	surname varchar(50),
	username varchar(50),
	password varchar(500)
);
`

func main() {
	db, err = sqlx.Connect("postgres", "user=postgres dbname=postgresgolang sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	db.MustExec(schema)

	users := []User{}
	tx := db.MustBegin()
	tx.Select(&users, "SELECT * FROM user")
	tx.Commit()
	fmt.Println(users)

}

func (tx sqlx.Tx) close() {

}
