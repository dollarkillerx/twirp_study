package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	helloworld "twirp/demo/demo1/pb"
)

func main() {
	client := helloworld.NewHelloWorldProtobufClient("http://localhost:8080", &http.Client{})
	hello, err := client.Hello(context.TODO(), &helloworld.HelloReq{
		Name: "DollarKiller",
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(hello)
}
