package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgresql://maxroach@localhost:26257/bank?sslmode=disable")
	if err != nil {
		log.Fatalf("error connection to the database: %s", err)
	}

	insert_sql := "INSERT INTO accounts (id, balance) VALUES (1, 1000), (2, 250)"
	if _, err := db.Exec(insert_sql); err != nil {
		log.Fatalf("failed to execute insert sql: %s", err)
	}

	select_sql := "SELECT id, balance FROM accounts"
	rows, err := db.Query(select_sql)
	if err != nil {
		log.Fatalf("failed to execute select sql: %s", err)
	}
	defer rows.Close()

	fmt.Println("Initial balances:")
	for rows.Next() {
		var id, balance int
		if err := rows.Scan(&id, &balance); err != nil {
			log.Fatalf("failed to scan the rows: %s", err)
		}
		fmt.Printf("%d %d\n", id, balance)
	}
}
