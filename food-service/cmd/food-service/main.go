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
)

const defaultPort = "8080"

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	dbPath := os.Getenv("PATH")
	err := db.Setup(dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	service.Init()

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
