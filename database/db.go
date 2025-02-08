package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver
	"log"
)

var DB *sqlx.DB

// ConnectToDB establishes a connection to the database
func ConnectToDB() {
	var err error
	// Replace with your actual connection details
	DB, err = sqlx.Connect("postgres", "user=postgres dbname=bus_navigator password=your_password_here sslmode=disable")
	if err != nil {
		log.Fatalln("Error connecting to the database:", err)
	}
	fmt.Println("✅ Connected to the PostgreSQL database successfully!")
}
