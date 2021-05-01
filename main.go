package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/OscarClemente/backend-tech-challenge-time/db/migration"
	"github.com/OscarClemente/backend-tech-challenge-time/db/postgres"
	"github.com/OscarClemente/backend-tech-challenge-time/graph"
	"github.com/OscarClemente/backend-tech-challenge-time/graph/generated"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

const defaultPort = "8080"
const defaultDbAddress = "postgres:5432"

func main() {
	// Wait for DB to be up, ugly & dirty solution
	// ideal solution is to use scripts
	// for healthcheck between dockers
	time.Sleep(5 * time.Second)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	dbAddress := os.Getenv("DB_ADDRESS")
	if dbAddress == "" {
		dbAddress = defaultDbAddress
	}

	postgres.New(dbAddress)
	migration.CreateSchema()
	migration.PopulateWithTestData()

	// Allow specific connections
	router := chi.NewRouter()
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080", "http://localhost:5000", "http://0.0.0.0:5000"},
		AllowCredentials: true,
		Debug:            false,
	}).Handler)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	// log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Printf("Serving at /query in port: %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
