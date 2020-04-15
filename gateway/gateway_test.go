package main

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"log"
	"net/http"
	"strings"
	"testing"
)

func TestHand_ServeHTTP(t *testing.T) {
	log.Println("Success")
	log.Fatalln(http.ListenAndServe("0.0.0.0:8085", &Router{}))
}

type Router struct {
	html  http.HandlerFunc
	query *handler.Server
}

func (h *Router) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Location", "http://baidu.com")
	resp.WriteHeader(302)
}

func TestPath(t *testing.T) {
	path1 := "/sadsad/sadsad"
	//path1 := "/"
	path1 = path1[1:]
	split := strings.Split(path1, "/")
	fmt.Println(split)
	fmt.Println(len(split))
}
