package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/idkwhyureadthis/food-service/graph"
	"github.com/idkwhyureadthis/food-service/internal/db"
	"github.com/idkwhyureadthis/food-service/internal/middleware"
	"github.com/idkwhyureadthis/food-service/internal/service"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	dbData := db.DbData{
		User:     os.Getenv("USER"),
		Password: os.Getenv("PASSWORD"),
		Host:     os.Getenv("HOST"),
		DbName:   os.Getenv("DBNAME"),
		Port:     os.Getenv("DBPORT"),
	}
	err = db.Setup(dbData)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	service.Init()
	db.Reset()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", middleware.AuthMiddleWare(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
