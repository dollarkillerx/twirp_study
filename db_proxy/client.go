package db_proxy

import (
	"net/http"
	"twirp/pb"
)

func NewClient(addr string) pb.DbProxy {
	return pb.NewDbProxyProtobufClient(addr, &http.Client{})
}

//func NewClient(addr string,token string) (pb.DbProxy,error) {
//	client := pb.NewDbProxyProtobufClient(addr, &http.Client{})
//
//	header := make(http.Header)
//	header.Set("auth", token)
//
//	// Attach the headers to a context
//	ctx := context.Background()
//	ctx, err := twirp.WithHTTPRequestHeaders(ctx, header)
//	if err != nil {
//		log.Printf("twirp error setting headers: %s", err)
//		return nil,errors.New("dbproxy error")
//	}
//	return client,nil
//}
