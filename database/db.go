package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver
	"log"
	"os"
)

var DB *sqlx.DB

// ConnectToDB establishes a connection to the database
func ConnectToDB() {
	var err error
	// Replace with your actual connection details
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	DB, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalln("Error connecting to the database:", err)
	}
	fmt.Println("✅ Connected to the PostgreSQL database successfully!")
}
