package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	helloworld "twirp/demo/demo1/pb"
)

func main() {
	server := helloworld.NewHelloWorldServer(&Server{}, nil)
	mux := http.NewServeMux()
	mux.Handle(server.PathPrefix(),server)
	http.ListenAndServe(":8080",mux)
}

type Server struct {

}

func (s *Server) Hello(ctx context.Context,req *helloworld.HelloReq) (*helloworld.HelloResp,error) {
	if req == nil {
		return nil,errors.New("req is nil")
	}
	return &helloworld.HelloResp{
		Body: fmt.Sprintf("Hello %s \n",req.Name),
	},nil
}