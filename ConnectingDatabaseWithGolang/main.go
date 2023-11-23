package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	// Replace these values with your actual SQL Server connection details
	connString := "Server=localhost\\MSSQLSERVER01;Database=customers;Trusted_Connection=True;"

	// Server=localhost\MSSQLSERVER03;Database=master;Trusted_Connection=True;

	// Create a new context
	ctx := context.Background()

	// Open a connection to the SQL Server
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Check if the connection to the database is successful
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database")

	// Update the "customers" table
	customerID := 1          // Replace with the actual customer ID you want to update
	newPostalCode := "54321" // Replace with the new postal code

	// Prepare the SQL statement
	stmt, err := db.PrepareContext(ctx, "UPDATE customer_t SET postal_code = ? WHERE customer_id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Execute the SQL statement
	_, err = stmt.ExecContext(ctx, newPostalCode, customerID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Updated postal code for customer ID %d to %s\n", customerID, newPostalCode)
}
