package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"native-restful-api/helper"
	"strconv"
)

func DbConnect() *sql.DB {
	defer func() {
		message := recover()

		if message != nil {
			log.Fatal(message)
		}
	}()

	port, err := strconv.Atoi(helper.Env("DB_PORT"))

	if err != nil {
		panic(err)
	}

	dbCredential := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", helper.Env("DB_HOST"), port, helper.Env("DB_USER"), helper.Env("DB_PASSWORD"), helper.Env("DB_NAME"))

	db, err := sql.Open("postgres", dbCredential)

	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	log.Print("Database successfully connected")

	return db
}
