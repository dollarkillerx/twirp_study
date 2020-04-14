package main

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"log"
	"net/http"
	"os"
	"twirp/gateway/graph"
	"twirp/gateway/graph/generated"
)

func main() {
	addr := os.Getenv("GATEWAYADDR")
	if addr == "" {
		addr = "0.0.0.0:8085"
	}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		Addr: fmt.Sprintf("http://%s", addr),
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to %s for GraphQL playground", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
