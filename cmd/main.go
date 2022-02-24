package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/RickardA/multiuser/graph"
	"github.com/RickardA/multiuser/graph/generated"
	"github.com/RickardA/multiuser/graph/model"
	"github.com/RickardA/multiuser/internal/app"
	"github.com/RickardA/multiuser/internal/pkg/repository/mongo"
	"github.com/RickardA/multiuser/internal/pkg/sync_handler"
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

	syncHandler, err := sync_handler.New(&db)

	client := app.NewClient(&db, syncHandler)
	// Setup Client

	ticker := time.NewTicker(10 * time.Second)
	quit := make(chan struct{})
	go func() {
		fmt.Println("Go func")
		for {
			select {
			case <-ticker.C:
				fmt.Println("Should send notifications")
				for _, element := range client.Subs {
					fmt.Println("Sending conflict notification")
					element <- &model.GQConflict{
						ID:               "test",
						RunwayID:         "test2",
						ResolutionMethod: "naaaajjs",
					}
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{Client: client}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

	fmt.Println("Hej hopp")
}
