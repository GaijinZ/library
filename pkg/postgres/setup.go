package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var PostgreSQL *sql.DB

func SetUpPostgres() {
	var err error
	// connection string
	psqlconn := GetPostgresConnectionString()

	// open database
	PostgreSQL, err = sql.Open("postgres", psqlconn)
	if err != nil {
		log.Panic(err)
	}

	// close database
	defer PostgreSQL.Close()

	// check db
	err = PostgreSQL.Ping()
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Connected!")
}
