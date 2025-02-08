package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var DB *sqlx.DB

func ConnectDB() {
	var err error
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"),
	)
	DB, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	fmt.Println("Connected to PostgreSQL")
}
