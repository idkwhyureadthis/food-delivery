package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	connUrl := os.Getenv("PATH")
	if connUrl == "" {
		log.Fatal("database path not found")
	}
}
