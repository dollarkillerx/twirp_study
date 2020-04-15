package server

import (
	"context"
	"os"
	"twirp/db_proxy"
	"twirp/pb"
	"twirp/utils"
)

var (
	dbProxy = "http://0.0.0.0:8081"
)

func init() {
	gp := os.Getenv("DBPROXY")
	if gp != "" {
		dbProxy = gp
	}
}

type ShortUrl struct {
}

func (s *ShortUrl) Parsing(ctx context.Context, req *pb.ParsingUrl) (*pb.ParsingKey, error) {
	client := db_proxy.NewClient(dbProxy)

	ex, err := client.CheckUrlEx(ctx, &pb.CheckUrlExReq{Url: req.Url})
	if err != nil {
		return nil, err
	}

	encode := utils.Md5Encode(req.Url)

	if ex.Key != "" {
		return &pb.ParsingKey{
			Key: ex.Key,
		}, nil
	}

	// First appearance
	k := ""
	idx := 5
	for {
		k = encode[:idx]
		url, err := client.GetUrl(ctx, &pb.GetUrlReq{
			Key: k,
		})
		if err != nil {
			return nil, err
		}
		if url.Addr != "" {
			idx++
			continue
		}
		break
	}

	_, err = client.AddUrl(ctx, &pb.UrlReq{Key: k, Url: req.Url})
	if err != nil {
		return nil, err
	}
	return &pb.ParsingKey{Key: k}, nil
}

func (s *ShortUrl) UnParsing(ctx context.Context, req *pb.ParsingKey) (*pb.ParsingUrl, error) {
	client := db_proxy.NewClient(dbProxy)
	url, err := client.GetUrl(ctx, &pb.GetUrlReq{Key: req.Key})
	if err != nil {
		return nil, err
	}

	return &pb.ParsingUrl{Url: url.Addr}, nil
}
