package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/supabase-community/supabase-go"
	"log"
	"os"
)

func GetEnv(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

// Hello handler
func Hello(c *fiber.Ctx) error {

	client, err := supabase.NewClient(GetEnv("API_URL"), GetEnv("API_KEY"), &supabase.ClientOptions{})

	if err != nil {
		fmt.Println("cannot initalize client", err)
	}

	data, count, err := client.From("countries").Select("*", "exact", false).Execute()

	log.Println(data, count)

	return c.SendString("Hello, World!")
}
