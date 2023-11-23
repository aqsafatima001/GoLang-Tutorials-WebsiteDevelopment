package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

func connectToSQLServer() (*sql.DB, error) {
	server := "localhost"
	port := 1433
	user := "MSSQLSERVER01"
	database := "customers"

	connString := fmt.Sprintf("server=%s;user id=%s;port=%d;database=%s", server, user, port, database)
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		return nil, err
	}

	// Ping the SQL Server to ensure connectivity
	err = db.PingContext(context.Background())
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to SQL Server!")
	return db, nil
}

func main() {
	// Call your connectToSQLServer function
	_, err := connectToSQLServer()
	if err != nil {
		fmt.Println("Error connecting to SQL Server:", err)
	}
}
