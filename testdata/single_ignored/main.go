package main

import (
	"database/sql"
	"fmt"
)

func main() {
	fmt.Println(query("'test' OR 1=1"))
}

const GetAllQuery = "SELECT COUNT(*) FROM t WHERE arg=%s"

// For this test we expect the second QueryRow to have an SQL injection issue
func query(arg string) error {
	db, err := sql.Open("postgres", "postgresql://test:test@test")
	if err != nil {
		return err
	}

	query := fmt.Sprintf(GetAllQuery, arg)
	//nolint:safesql
	row := db.QueryRow(query)
	var count int
	if err := row.Scan(&count); err != nil {
		return err
	}

	row = db.QueryRow(fmt.Sprintf(GetAllQuery, "Catch me please?"))
	if err := row.Scan(&count); err != nil {
		return err
	}

	return nil
}
