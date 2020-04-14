package main

import (
	"log"
	"net/http"
	"os"
	"twirp/db_proxy/server"
	"twirp/pb"
)

func main() {
	addr := os.Getenv("DBPROXY_ADDR")
	if addr == "" {
		addr = "0.0.0.0:8081"
	}
	server := pb.NewDbProxyServer(&server.DbProxy{}, nil)
	mux := http.NewServeMux()
	mux.Handle(server.PathPrefix(), server)
	log.Fatalln(http.ListenAndServe(addr, mux))
}
