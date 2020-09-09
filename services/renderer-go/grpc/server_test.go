package grpc

import (
	"context"
	"testing"

	pb_fetcher "github.com/peaceiris/Hatena-Intern-2020/services/renderer-go/pb/fetcher"
	pb "github.com/peaceiris/Hatena-Intern-2020/services/renderer-go/pb/renderer"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

type fakeFecherClient struct {
	FakeFetch func(ctx context.Context, req *pb_fetcher.FetchRequest, opt ...grpc.CallOption) (*pb_fetcher.FetchReply, error)
}

func (c *fakeFecherClient) Fetch(ctx context.Context, req *pb_fetcher.FetchRequest, opt ...grpc.CallOption) (*pb_fetcher.FetchReply, error) {
	return c.FakeFetch(ctx, req)
}

func Test_Server_Render(t *testing.T) {
	fecherCli := &fakeFecherClient{
		FakeFetch: func(ctx context.Context, req *pb_fetcher.FetchRequest, opt ...grpc.CallOption) (*pb_fetcher.FetchReply, error) {
			return &pb_fetcher.FetchReply{Title: "Example Domain"}, nil
		},
	}
	s := NewServer(fecherCli)
	src := "foo https://google.com/ bar"
	reply, err := s.Render(context.Background(), &pb.RenderRequest{Src: src})
	assert.NoError(t, err)
	assert.Equal(t, "<p>foo <a href=\"https://google.com/\">https://google.com/</a> bar</p>\n", reply.Html)
}
