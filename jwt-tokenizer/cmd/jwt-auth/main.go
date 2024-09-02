package main

import (
	"log"
	"os"

	"github.com/idkwhyureadthis/food-delivery/jwt-tokenizer/internal/app"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	dbPath := os.Getenv("PATH")
	secretKey := os.Getenv("SECRET")
	a, err := app.New(dbPath, []byte(secretKey))
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(a.Run(":8081"))
}
