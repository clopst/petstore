package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "secret"
	dbname   = "petstores"
)

var DBClient *sql.DB

func init() {
	connectionString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	var err error
	DBClient, err = sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database.")
}
