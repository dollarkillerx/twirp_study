package server

import (
	"context"
	"os"
	"twirp/pb"
	"twirp/utils"
)

var (
	dbProxy string
)

func init() {
	gp := os.Getenv("DBPROXY")
	if gp == "" {
		dbProxy = "0.0.0.0:8081"
	}
	dbProxy = gp
}

type ShortUrl struct {

}

func Parsing(ctx context.Context,req *pb.ParsingUrl) (*pb.ParsingKey,error) {
	encode := utils.Md5Encode(req.Url)
	return nil, nil
}

func UnParsing(ctx context.Context,req *pb.ParsingKey) (*pb.ParsingUrl,error) {

	return nil, nil
}