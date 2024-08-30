package main

import (
	"log"
	"os"

	"github.com/idkwhyureadthis/food-delivery/jwt-tokenizer/internal/app"
)

func main() {
	dbPath := os.Getenv("PATH")
	a, err := app.New(dbPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(a.Run(":8081"))
}
