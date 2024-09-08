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

/*
create docker instance: docker run --name htmx-experiments -h localhost -p 3307:3306 -v %CD%\data -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=htmx_experiments -d mysql:8.0-debian
open container bash: docker exec -it htmx-experiments bash
map ports and set protocol to tcp:  mysql -h localhost -P 3307 -u root -p
enter password: password

done :)
*/