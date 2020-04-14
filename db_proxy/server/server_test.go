package server

import (
	"context"
	"fmt"
	"log"
	"testing"
	"twirp/pb"
)

func TestDbProxy_CheckUrlEx(t *testing.T) {
	proxy := DbProxy{}
	url := "jjj"
	req, err := proxy.CheckUrlEx(context.TODO(), &pb.CheckUrlExReq{Url: url})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(req)
}

func TestDbProxy_AddUrl(t *testing.T) {
	proxy := DbProxy{}
	url := "jjj"
	addUrl, err := proxy.AddUrl(context.TODO(), &pb.UrlReq{Key: "zz", Url: url})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(addUrl)
}

func TestDbProxy_GetUrl(t *testing.T) {
	proxy := DbProxy{}
	url, err := proxy.GetUrl(context.TODO(), &pb.GetUrlReq{Key: "47f5f"})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(url)
}
