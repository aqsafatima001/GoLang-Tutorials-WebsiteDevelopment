package main

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	// Define connection string
	// connString := "server=tcp:LAPTOP-G5TDHLRV\\SQL_IAD;port=1433;database=PVFC;user id=LAPTOP-G5TDHLRV\\Aqsa Fatima;password=123;Trusted_Connection=True"
	// connString := "server=LAPTOP-G5TDHLRV\\SQL_IAD;port=1433;database=PVFC;user id=LAPTOP-G5TDHLRV\\Aqsa Fatima;password=123;Trusted_Connection=True"

	connString := "server=LAPTOP-G5TDHLRV\\SQL_IAD;port=1433;database=PVFC;user id=pvfc;password=pvfc;"

	// Open a connection to the database
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		fmt.Println("Error connecting to the database:", err.Error())
		return
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging database:", err.Error())
		return
	}

	fmt.Println("Connected to the database!")

	// Insert data into the Customer table
	// _, err = db.Exec(`
	//     USE PVFC;
	//     INSERT INTO Customer (CustomerID, FirstName, LastName, Email)
	//     VALUES (3, 'Alice', 'Johnson', 'alice.johnson@example.com');
	// `)
	// if err != nil {
	// 	fmt.Println("Error inserting data:", err.Error())
	// 	return
	// }

	// fmt.Println("Data inserted successfully!")

	rows, err := db.Query("SELECT CustomerID, FirstName, LastName, Email FROM Customer")
	if err != nil {
		fmt.Println("Error querying data:", err.Error())
		return
	}
	defer rows.Close()

	fmt.Println("Retrieved data from the Customer table:")
	for rows.Next() {
		var customerID int
		var firstName, lastName, email string

		err := rows.Scan(&customerID, &firstName, &lastName, &email)
		if err != nil {
			fmt.Println("Error scanning row:", err.Error())
			return
		}

		fmt.Printf("%d: %s %s (%s)\n", customerID, firstName, lastName, email)
	}

}
