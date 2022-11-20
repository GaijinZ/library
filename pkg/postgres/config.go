package postgres

import "fmt"

const (
	host     = "postgres"
	port     = 5432
	user     = "userapi"
	password = "userapi"
	dbname   = "userapi"
)

func GetPostgresConnectionString() string {
	database := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	return database
}
