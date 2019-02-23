package main

import (
	"github.com/99designs/gqlgen/handler"
	prisma "github.com/tonymj76/DBB/generated/prisma-client"
	"log"
	"net/http"
	"os"
)

const defaultPort = "4000"

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = defaultPort
	}

	client := prisma.New(nil)
	resolver := Resolver{
		Prisma: client,
	}

	http.Handle("/", handler.Playground("GraphQL Playground", "/query"))
	http.Handle("/query", handler.GraphQL(NewExecutableSchema(Config{Resolvers: &resolver})))

	log.Printf("Server is running on http://localhost:%s", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
