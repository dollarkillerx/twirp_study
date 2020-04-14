package short_url

import (
	"net/http"
	"twirp/pb"
)

func NewClient(addr string) pb.ShortUrl {
	return pb.NewShortUrlProtobufClient(addr, &http.Client{})
}
