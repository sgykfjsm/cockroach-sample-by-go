package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/cockroachdb/cockroach-go/crdb"
)

func transferFunds(tx *sql.Tx, from int, to int, amount int) error {
	// Read the balance
	var fromBalance int
	select_sql := "SELECT balance FROM accounts WHERE id = $1"
	if err := tx.QueryRow(select_sql, from).Scan(&fromBalance); err != nil {
		return err
	}

	if fromBalance < amount {
		return fmt.Errorf("insufficient funds")
	}

	// Perform the transfer
	update_sql1 := "UPDATE accounts SET balance = balance - $1 WHERE id = $2"
	update_sql2 := "UPDATE accounts SET balance = balance + $1 WHERE id = $2"
	if _, err := tx.Exec(update_sql1, amount, from); err != nil {
		return err
	}

	if _, err := tx.Exec(update_sql2, amount, from); err != nil {
		return err
	}

	return nil
}

func main() {
	db, err := sql.Open("postgres", "postgresql://maxroach@localhost:26257/bank?sslmode=disable")
	if err != nil {
		log.Fatalf("error connection to the database: %s", err)
	}

	crdb.ExecuteTx(db, func(tx *sql.Tx) error {
		return transferFunds(tx, 1 /* from acct# */, 2 /* to acct# */, 100 /* amount */)
	})
	if err != nil {
		log.Fatalf("failed to execute transaction: %s", err)
	}
	fmt.Println("Success")
}
