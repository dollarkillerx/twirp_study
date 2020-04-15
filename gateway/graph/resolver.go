package graph

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	"fmt"
	"os"
	"twirp/gateway/graph/generated"
	"twirp/gateway/graph/model"
	"twirp/pb"
	"twirp/short_url"
)

var (
	ShortUrl = "http://0.0.0.0:8082"
)

func init() {
	addr := os.Getenv("SHORTURL")
	if addr != "" {
		ShortUrl = addr
	}
}

type Resolver struct {
	Addr string
}

func (r *mutationResolver) AddURL(ctx context.Context, input *model.NewURL) (*pb.ParsingKey, error) {
	client := short_url.NewClient(ShortUrl)
	parsing, err := client.Parsing(ctx, &pb.ParsingUrl{
		Url: input.URL,
	})
	if err != nil {
		return nil,err
	}
	parsing.Key = fmt.Sprintf("%s/%s", r.Addr, parsing.Key)
	return parsing, nil
}

func (r *queryResolver) GetURL(ctx context.Context, key string) (*pb.ParsingUrl, error) {
	client := short_url.NewClient(ShortUrl)
	return client.UnParsing(ctx, &pb.ParsingKey{Key: key})
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
