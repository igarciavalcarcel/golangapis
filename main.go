package main

import (
	"fmt"
	"golangapis/database"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	databaseConnection := database.InitDB()

	//logica Negocio

	defer databaseConnection.Close()
	fmt.Println(databaseConnection)
}
