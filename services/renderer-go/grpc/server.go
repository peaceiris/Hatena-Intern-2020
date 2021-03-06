package grpc

import (
	"context"
	"fmt"

	pb_fetcher "github.com/peaceiris/Hatena-Intern-2020/services/renderer-go/pb/fetcher"
	pb_image_fetcher "github.com/peaceiris/Hatena-Intern-2020/services/renderer-go/pb/image-fetcher"
	pb "github.com/peaceiris/Hatena-Intern-2020/services/renderer-go/pb/renderer"
	"github.com/peaceiris/Hatena-Intern-2020/services/renderer-go/renderer"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

// Server は pb.RendererServer に対する実装
type Server struct {
	pb.UnimplementedRendererServer
	fetcherClient         pb_fetcher.FetcherClient
	ogpImageFetcherClient pb_image_fetcher.FetcherClient
	healthpb.UnimplementedHealthServer
}

// NewServer は gRPC サーバーを作成する
func NewServer(fetcherClient pb_fetcher.FetcherClient, ogpImageFetcherClient pb_image_fetcher.FetcherClient) *Server {
	return &Server{fetcherClient: fetcherClient, ogpImageFetcherClient: ogpImageFetcherClient}
}

// Render は受け取った文書を HTML に変換する
func (s *Server) Render(ctx context.Context, in *pb.RenderRequest) (*pb.RenderReply, error) {
	html, err := renderer.Render(ctx, in.Src, s.fetcherClient, s.ogpImageFetcherClient)
	if err != nil {
		fmt.Errorf("failed to render: %+v", err)
		return nil, err
	}
	return &pb.RenderReply{Html: html}, nil
}
