package main

import (
	"log"
	"net/http"
	"os"
	"twirp/pb"
	"twirp/short_url/server"
)

func main() {
	addr := os.Getenv("SHORTURL_ADDR")
	if addr == "" {
		addr = "0.0.0.0:8082"
	}
	server := pb.NewShortUrlServer(&server.ShortUrl{}, nil)
	mux := http.NewServeMux()
	mux.Handle(server.PathPrefix(),server)
	log.Fatalln(http.ListenAndServe(addr,server))
}
