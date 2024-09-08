package main

import (
	"database/sql"
	"fmt"
)


func main() {
	db, err := sql.Open("mysql", "username:password@(127.0.0.1:3306)/dbname?parseTime=true")

	if err != nil {
		fmt.Println("Error configuring database connection:", err)
	}

	err = db.Ping()

	if err != nil {
		fmt.Println("Error connecting to database:", err)
	}
}