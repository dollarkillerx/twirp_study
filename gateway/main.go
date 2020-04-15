package main

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"log"
	"net/http"
	"os"
	"strings"
	"twirp/gateway/graph"
	"twirp/gateway/graph/generated"
	"twirp/pb"
	"twirp/short_url"
)

func main() {
	addr := os.Getenv("GATEWAYADDR")
	if addr == "" {
		addr = "0.0.0.0:8085"
	}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		Addr: fmt.Sprintf("http://%s", addr),
	}}))

	//http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	//http.Handle("/query", srv)

	router := &router{html: playground.Handler("GraphQL playground", "/query"), query: srv}

	log.Printf("connect to %s for GraphQL playground", addr)
	//log.Fatal(http.ListenAndServe(addr, nil))
	log.Fatal(http.ListenAndServe(addr, router))
}

type router struct {
	html  http.HandlerFunc
	query *handler.Server
}

func (h *router) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	baseUrl := req.RequestURI[1:]
	switch baseUrl {
	case "":
		h.html.ServeHTTP(resp, req)
	case "query":
		h.query.ServeHTTP(resp, req)
	default:
		split := strings.Split(baseUrl, "/")
		if len(split) != 1 {
			resp.WriteHeader(404)
			return
		}
		key := split[0]
		client := short_url.NewClient(graph.ShortUrl)
		parsing, err := client.UnParsing(context.TODO(), &pb.ParsingKey{
			Key: key,
		})
		if err != nil {
			resp.WriteHeader(404)
			return
		}
		if parsing.Url == "" {
			resp.WriteHeader(404)
			return
		}

		resp.Header().Set("Location", parsing.Url)
		resp.WriteHeader(302)
	}
}
