package database

import (
	"database/sql"
	"fmt"
)

func InitDB() *sql.DB {
	connectionString := "root:admin@tcp(localhost:3606)/northwind"

	databaseConnection, err := sql.Open("mysql", connectionString)
	if err != nil {
		fmt.Println("Errrorrr")
		panic(err.Error()) //error handling
	}
	return databaseConnection
}
