package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/RickardA/multiuser/graph"
	"github.com/RickardA/multiuser/graph/generated"
	"github.com/RickardA/multiuser/internal/app"
	"github.com/RickardA/multiuser/internal/pkg/repository/mongo"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	db, err := mongo.NewConnection(ctx, "mongodb://localhost")

	if err != nil {
		panic("Could not connect to db")
	}

	client := app.NewClient(&db)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{Client: client}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

	// Setup Client

}
