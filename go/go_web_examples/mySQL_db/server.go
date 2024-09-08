package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)


func main() {
	db, err := sql.Open("mysql", "root:password@(localhost:3307)/htmx_experiments?parseTime=true")

	if err != nil {
		fmt.Println("Error configuring database connection:", err)
	}

	err = db.Ping()

	if err != nil {
		fmt.Println("Error connecting to database:", err)
	}
}