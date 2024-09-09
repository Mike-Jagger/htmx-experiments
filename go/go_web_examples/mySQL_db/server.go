package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	id int64
	username string
	password string
	createdAt time.Time
}


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

	var lastUserId int64
	password := "password"
	usernames := []string{"johndoe", "marylyn", "mikejagger", "someone"}
	for _, username := range usernames {
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

		lastUserId, err := result.LastInsertId()
		if err != nil {
			fmt.Println("Error grabbing user id:", err)
			return
		} 

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			fmt.Println("Error grabbing user id:", err)
			return
		}

		fmt.Printf("\nAdded user with id %v aftected %v rows\n", lastUserId, rowsAffected)
	
		var user User
	
		query := `SELECT id, username, password, created_at FROM users WHERE id = ?`
	
		err = db.QueryRow(query, lastUserId).Scan(&user.id, &user.username, &user.password, &user.createdAt)
		if err != nil {
			fmt.Println("Error while getting user row:", err)
			return
		}
	
		fmt.Println("Here is info about this user")
		fmt.Println(
			"\t-> id:", user.id,
			"\n\t-> username:", user.username,
			"\n\t-> password:", user.password,
			"\n\t-> Account created at:", user.createdAt,	
		)	
	}


	rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
	if err != nil {
		fmt.Println("Error while executing query:", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
		if err != nil {
			fmt.Println("Error while trying to get user row:", err)
			return
		}
		users = append(users, u)
		lastUserId = u.id
	}

	err = rows.Err()
	if err != nil {
		fmt.Println("Error while reading rows:", err)
		return
	}

	fmt.Println("\nHere are the users gotten from the table")
	for _, user := range users {
		fmt.Println("\t- ", user)
	}

	result, err := db.Exec(`DELETE FROM users WHERE id = ?`, lastUserId)

	if err != nil {
		fmt.Println("Error deleting user from database:", err)
	}

	_, err = result.LastInsertId()
	if err != nil {
		fmt.Println("Error grabbing user id:", err)
		return
	} 

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error grabbing user id:", err)
		return
	}

	fmt.Printf("\nDeleted user with id %v aftected %v rows\n", lastUserId, rowsAffected)

}

/*
create docker instance: docker run --name htmx-experiments -h localhost -p 3307:3306 -v %CD%\data -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=htmx_experiments -d mysql:8.0-debian
open container bash: docker exec -it htmx-experiments bash
map ports and set protocol to tcp:  mysql -h localhost -P 3307 -u root -p
enter password: password

done :)
*/