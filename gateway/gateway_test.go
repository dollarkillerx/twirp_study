package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"testing"
)

func TestHand_ServeHTTP(t *testing.T) {
	log.Println("Success")
	log.Fatalln(http.ListenAndServe("0.0.0.0:8085", &router{}))
}

func TestPath(t *testing.T) {
	path1 := "/sadsad/sadsad"
	//path1 := "/"
	path1 = path1[1:]
	split := strings.Split(path1, "/")
	fmt.Println(split)
	fmt.Println(len(split))
}
