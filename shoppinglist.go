package main

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

func main() {
	var (
		id int
		type_id int
		name string
		description string
		json string
	)

	db, err := sql.Open("postgres", "user=xxx password=xxx dbname=xxx sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("SELECT * FROM product")
	if err != nil {
			log.Fatal(err)
		}

	defer rows.Close()
	
	for rows.Next() {
		err := rows.Scan(&id, &type_id, &description, &name, &json)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name, description, type_id, json)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}