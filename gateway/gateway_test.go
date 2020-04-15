package main

import (
	"context"
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

func TestContext(t *testing.T) {
	ctx := context.Background()

	// 蛤 成了新的ctx
	ctx = context.WithValue(ctx, "key", "val")
	//fmt.Println(ctx)

	i := ctx.Value("key")
	fmt.Println(i)
}

func TestContext2(t *testing.T) {
	a := func (ctx context.Context) {
		ret,ok := ctx.Value("trace_id").(int)
		if !ok {
			ret = 21342423
		}

		fmt.Printf("ret:%d\n", ret)

		s , _ := ctx.Value("session").(string)
		fmt.Printf("session:%s\n", s)

		fmt.Println(ctx.Value("session"))
	}

	ctx := context.WithValue(context.Background(), "trace_id", 13483434)
	ctx = context.WithValue(ctx, "session", "sdlkfjkaslfsalfsafjalskfj")

	a(ctx)
}