package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/heaths/booksql/graphql"
	"github.com/heaths/booksql/graphql/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("FUNCTIONS_CUSTOMHANDLER_PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graphql.Resolver{}}))

	http.Handle("/graphql", srv)
	http.Handle("/graphql/play", playground.Handler("GraphQL playground", "/graphql"))

	log.Printf("Connect to http://localhost:%s/graphql/play for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
