package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)


func main() {
	db, err := sql.Open("mysql", "root:password@(localhost:3307)/htmx_experiments?parseTime=true")
	if err != nil {
		fmt.Println("Error configuring database connection:", err)
		return
	}
	
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}

	dropTableQuery := "DROP TABLE IF EXISTS users"
	createTabeleQuery := `
		CREATE TABLE users(
			id INT AUTO_INCREMENT,
			username TEXT NOT NULL,
			password TEXT NOT NULL,
			created_at DATETIME,
			PRIMARY KEY (id)
		);
	`

	_, err = db.Exec(dropTableQuery)

	if err != nil {
		fmt.Println("Error creating USERS table:", err)
	}

	_, err = db.Exec(createTabeleQuery)

	if err != nil {
		fmt.Println("Error creating USERS table:", err)
	}

	username := "johndoe"
	password := "password"
	createdAt := time.Now()

	result, err := db.Exec(
		`INSERT INTO users (username, password, created_at)
		 VALUES (?, ?, ?)`, 
		username, 
		password, 
		createdAt)

	if err != nil {
		fmt.Println("Error adding new user to table:", err)
		return
	}

	userId, err := result.LastInsertId()
	if err != nil {
		fmt.Println("Error grabbing user id:", err)
		return
	} 

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error grabbing user id:", err)
		return
	}

	fmt.Printf("UserId with id %v aftected %v\n", userId, rowsAffected)

	type User struct {
		id int
		username string
		password string
		createdAt time.Time
	}

	var user User

	query := `SELECT id, username, password, created_at FROM users WHERE id = ?`

	err = db.QueryRow(query, userId).Scan(&user.id, &user.username, &user.password, &user.createdAt)

	if err != nil {
		fmt.Println("Error while getting user row:", err)
	}

	fmt.Println("\nHere is info about this user")
	fmt.Println(
		"\t-> id:", user.id,
		"\n\t-> username:", user.username,
		"\n\t-> password:", user.password,
		"\n\t-> Account created at:", user.createdAt,	
	)	
}

/*
create docker instance: docker run --name htmx-experiments -h localhost -p 3307:3306 -v %CD%\data -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=htmx_experiments -d mysql:8.0-debian
open container bash: docker exec -it htmx-experiments bash
map ports and set protocol to tcp:  mysql -h localhost -P 3307 -u root -p
enter password: password

done :)
*/