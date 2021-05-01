package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/OscarClemente/backend-tech-challenge-time/db/migration"
	"github.com/OscarClemente/backend-tech-challenge-time/db/postgres"
	"github.com/OscarClemente/backend-tech-challenge-time/graph"
	"github.com/OscarClemente/backend-tech-challenge-time/graph/generated"
)

const defaultPort = "8080"

func main() {
	postgres.New()
	migration.CreateSchema()
	migration.PopulateWithTestData()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
