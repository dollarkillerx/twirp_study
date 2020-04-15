package main

import (
	"log"
	"net/http"
	"os"
	"twirp/discovery/middleware"
	"twirp/discovery/server"
	"twirp/pb"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalln("./discovery 0.0.0.0:8081 token")
	}
	server := pb.NewDiscoveryServer(server.New(), nil)
	mux := http.NewServeMux()
	mux.Handle(server.PathPrefix(), middleware.BaseAuth(server, os.Args[2]))
	http.ListenAndServe(os.Args[1], mux)
}
